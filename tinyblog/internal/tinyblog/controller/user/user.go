package user

import (
	"github.com/thoseJanes/tinyblog/internal/tinyblog/biz"
	"github.com/thoseJanes/tinyblog/internal/tinyblog/store"
	"github.com/thoseJanes/tinyblog/pkg/auth"
	pb "github.com/thoseJanes/tinyblog/pkg/proto/tinyblog/v1"
)


type UserController struct {
	a *auth.Authz
	b biz.IBiz
	pb.UnimplementedTinyBlogServer
}

func New(ds store.IStore, a *auth.Authz) *UserController {
	return &UserController{a: a, b: biz.New(ds)}
}