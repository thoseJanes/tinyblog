package biz

import(
	"github.com/thoseJanes/tinyblog/internal/tinyblog/biz/user"
	"github.com/thoseJanes/tinyblog/internal/tinyblog/biz/post"
)

//go:generate gencode ./interface.go -t i -o biz.go -r IBiz:biz
type IBiz interface {
	User() user.UserBiz
	Post() post.PostBiz

}