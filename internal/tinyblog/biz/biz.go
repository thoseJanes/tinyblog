package biz

import (
	"github.com/thoseJanes/tinyblog/internal/tinyblog/biz/post"
	"github.com/thoseJanes/tinyblog/internal/tinyblog/biz/user"
	"github.com/thoseJanes/tinyblog/internal/tinyblog/store"
)

//implement
//User() user.UserBiz
//Post() post.PostBiz
type biz struct {
	ds store.IStore
}

var _ IBiz = (*biz)(nil)


func New(ds store.IStore) IBiz {
	return &biz{ds}
}


func (b *biz) User() user.UserBiz {
	return user.New(b.ds)
}


func (b *biz) Post() post.PostBiz {
	return post.New(b.ds)
}



