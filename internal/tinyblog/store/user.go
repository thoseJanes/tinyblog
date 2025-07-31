package store

import (
	"context"
	"github.com/thoseJanes/tinyblog/internal/pkg/model"
	"gorm.io/gorm"
)

type userStore struct {
	db *gorm.DB
}

var _ UserStore = (*userStore)(nil)

func newUserStore(db *gorm.DB) *userStore {
	return &userStore{db}
}

func (s *userStore) Create(c context.Context, user *model.User) error {
	err := s.db.Create(user).Error
	return err
}

func (s *userStore) Get(c context.Context, username string) (*model.User, error) {
	user := model.User{Username: username}
	err := s.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

// func (s *userStore) ChangePassword(c context.Context, username string, newPassword string) error {
// 	err := s.db.Where("username=?", username).Update("password", newPassword).Error
// 	return err
// }

func (s *userStore) Update(c context.Context, user *model.User) error {
	err := s.db.Where("username=?", user.Username).Updates(*user).Error// 只更新，不获取更新结果，因此不用传入*model.User指针。
	return err
}

func (s *userStore) List(c context.Context, offset,limit int) (int64, []model.User, error) {
	var users []model.User
	var count int64
	err := s.db.Offset(offset).Limit(limit).Find(&users).Offset(-1).Limit(-1).Count(&count).Error
	return count, users, err
}

func (s *userStore) Delete(c context.Context, username string) error {
	err := s.db.Where("username = ?", username).Delete(&model.User{}).Error
	if(err == gorm.ErrRecordNotFound){
		return nil
	}
	return err
}