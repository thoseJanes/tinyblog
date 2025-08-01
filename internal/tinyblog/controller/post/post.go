package post

import (
	"github.com/thoseJanes/tinyblog/internal/tinyblog/biz"
	"github.com/thoseJanes/tinyblog/internal/tinyblog/store"
)



type PostController struct {
	b biz.IBiz
}


func New(ds store.IStore) *PostController {
	return &PostController{biz.New(ds)}
}