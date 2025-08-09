package store

import (
	"context"
	"github.com/thoseJanes/tinyblog/internal/pkg/model"
	"gorm.io/gorm"
)

type IStore interface {
	UserStore() UserStore
	PostStore() PostStore
	DB() *gorm.DB
}

type UserStore interface {
	Create(c context.Context, user *model.User) error
	Get(c context.Context, username string) (*model.User, error)
	// ChangePassword(c context.Context, username string, newPassword string) error
	Update(c context.Context, user *model.User) error
	List(c context.Context, offset,limit int) (int64, []model.User, error)
	Delete(c context.Context, username string) error
}

type PostStore interface {
	Create(c context.Context, post *model.Post) error
	Get(c context.Context, username, postId string) (*model.Post, error)
	Update(c context.Context, post *model.Post) error
	List(c context.Context, username string, offset,limit int) (int64, []model.Post, error)
	Delete(c context.Context, username string, postIds []string) error
}




