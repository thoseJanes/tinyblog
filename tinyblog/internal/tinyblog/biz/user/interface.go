package user

import (
	"context"
	"github.com/thoseJanes/tinyblog/pkg/api/tinyblog/v1"
)

//go:generate gencode ./interface.go -o user.go -t i -r UserBiz:userBiz
type UserBiz interface {
	Create(ctx context.Context, r *v1.CreateUserRequest) error
	Delete(ctx context.Context, username string) error
	Update(ctx context.Context, username string, r *v1.UpdateUserRequest) error
	ChangePassword(ctx context.Context, username string, r *v1.ChangePasswordRequest) error
	Login(ctx context.Context, r *v1.LoginRequest) (*v1.LoginResponse, error)
	Get(ctx context.Context, username string) (*v1.GetUserResponse, error)
	List(ctx context.Context, offset, limit int) (*v1.ListUserResponse, error)
}


