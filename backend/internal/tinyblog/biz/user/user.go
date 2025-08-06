package user

import (
	"context"
	"errors"
	"regexp"
	"sync"

	"github.com/jinzhu/copier"
	"github.com/thoseJanes/tinyblog/internal/pkg/errno"
	"github.com/thoseJanes/tinyblog/internal/pkg/log"
	"github.com/thoseJanes/tinyblog/internal/pkg/model"
	"github.com/thoseJanes/tinyblog/internal/tinyblog/store"
	"github.com/thoseJanes/tinyblog/pkg/api/tinyblog/v1"
	"github.com/thoseJanes/tinyblog/pkg/auth"
	"github.com/thoseJanes/tinyblog/pkg/token"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

//implement
//Create(ctx context.Context, r *v1.CreateUserRequest) error
//Delete(ctx context.Context, username string) error
//Update(ctx context.Context, username string, r *v1.UpdateUserRequest) error
//ChangePassword(ctx context.Context, username string, r *v1.ChangePasswordRequest) error
//Login(ctx context.Context, r *v1.LoginRequest) (v1.LoginResponse, error)
//Get(ctx context.Context, username string) (*v1.GetUserResponse, error)
//List(ctx context.Context, offset int, limit int) (*v1.ListUserResponse, error)
type userBiz struct {
	ds store.IStore
}

var _ UserBiz = (*userBiz)(nil)


func New(ds store.IStore) UserBiz {
	return &userBiz{ds}
}


func (u *userBiz) Create(ctx context.Context, r *v1.CreateUserRequest) error {
	var userM model.User
	copier.Copy(&userM, r)

	if err := u.ds.UserStore().Create(ctx, &userM); err != nil {
		if bl, _ := regexp.MatchString("Duplicate .*", err.Error()); bl {
			return errno.ErrUserAlreadyExist
		}
		return err
	}

	return nil
}


func (u *userBiz) Delete(ctx context.Context, username string) error {
	if err := u.ds.UserStore().Delete(ctx, username); err != nil {
		return err
	}
	return nil
}


func (u *userBiz) Update(ctx context.Context, username string, r *v1.UpdateUserRequest) error {
	var userM model.User

	// if r.Email != nil {
	// 	userM.Email = *r.Email
	// }
	// if r.Nickname != nil {
	// 	userM.Nickname = *r.Nickname
	// }
	// if r.Phone != nil {
	// 	userM.Phone = *r.Phone
	// }

	copier.Copy(&userM, r)
	userM.Username = username
	if err := u.ds.UserStore().Update(ctx, &userM); err != nil {
		return err
	}

	return nil
}


func (u *userBiz) ChangePassword(ctx context.Context, username string, r *v1.ChangePasswordRequest) error {
	user, err := u.ds.UserStore().Get(ctx, username)
	if err != nil {
		return err
	}

	if err := auth.Compare(r.OldPassword, user.Password); err != nil {
		return errno.ErrPasswordIncorrect
	}

	password, err := auth.Encrypt(r.NewPassword)
	if err != nil {
		return err
	}
	user.Password = password
	return u.ds.UserStore().Update(ctx, user)
}


func (u *userBiz) Login(ctx context.Context, r *v1.LoginRequest) (*v1.LoginResponse, error) {
	user, err := u.ds.UserStore().Get(ctx, r.Username)
	if err != nil {
		return nil, err
	}

	if err := auth.Compare(r.Password, user.Password); err != nil {
		return nil, errno.ErrPasswordIncorrect
	}

	token, err := token.Sign(r.Username)
	if err != nil{
		return nil, errno.ErrSignToken
	}

	return &v1.LoginResponse{Token: token}, nil
}


func (u *userBiz) Get(ctx context.Context, username string) (*v1.GetUserResponse, error) {
	user, err := u.ds.UserStore().Get(ctx, username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return nil, errno.ErrUserNotFound
		}
		return nil, err
	}

	var resp v1.GetUserResponse
	_ = copier.Copy(&resp, user)
	resp.CreatedAt = user.CreatedAt.Format("2006-01-02 15:04:05")
	resp.UpdatedAt = user.UpdatedAt.Format("2006-01-02 15:04:05")

	return &resp, nil
}


func (u *userBiz) List(ctx context.Context, offset int, limit int) (*v1.ListUserResponse, error) {
	tolPost, users, err := u.ds.UserStore().List(ctx, offset, limit)
	if err != nil {
		log.C(ctx).Errorw("Failed to list user from storage.", "err", err)
		return nil, err
	}

	var m sync.Map
	eg, c := errgroup.WithContext(ctx)
	for _, userM := range users {
		user := userM
		eg.Go(func() error {
			select{
			case <-c.Done():
				return nil
			default:
				count, _, err := u.ds.PostStore().List(c, user.Username, 0, 0)
				if err != nil {
					log.C(ctx).Errorw("Failed to get user posts count.", "err", err)
					return err
				}
				m.Store(user.Id, &v1.UserInfo{
					Username: user.Username,
					Nickname: user.Nickname,
					Email: user.Email,
					Phone: user.Phone,
					PostCount: count,
					CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
					UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
				})
				return nil
			}
		})
	}

	if err := eg.Wait(); err != nil {
		log.C(ctx).Errorw("Failed to handle all goroutines", "err", err)
		return nil, err
	}

	usersInfo := make([]*v1.UserInfo, len(users))
	for _, user := range users {
		userInfo, _ := m.Load(user.Id)
		usersInfo = append(usersInfo, userInfo.(*v1.UserInfo))
	}
	
	var resp v1.ListUserResponse
	resp.TotalCount = tolPost
	resp.Users = usersInfo
	return &resp, nil
}
