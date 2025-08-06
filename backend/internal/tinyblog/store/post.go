package store

import (
	"context"
	"github.com/thoseJanes/tinyblog/internal/pkg/model"
	"gorm.io/gorm"
)

// type PostStore interface {
// 	Create(c context.Context, post *model.Post) error
// 	Get(c context.Context, username, postId string) (*model.Post, error)
// 	Update(c context.Context, post *model.Post) error
// 	List(c context.Context, offset,limit int64) (int64, []*model.Post, error)
// }


type postStore struct{
	db *gorm.DB
}

var _ PostStore = (*postStore)(nil)

func newPostStore(db *gorm.DB) *postStore {
	return &postStore{db}
}

func (p *postStore)	Create(c context.Context, post *model.Post) error {
	return p.db.Create(post).Error
}
func (p *postStore)	Get(c context.Context, username, postId string) (*model.Post, error) {
	var post model.Post
	err := p.db.Where("username = ? and postId = ?", username, postId).First(&post).Error
	return &post, err
}
func (p *postStore)	Update(c context.Context, post *model.Post) error {
	return p.db.Where("username = ? and postId = ?", post.Username, post.PostId).Updates(*post).Error
}
func (p *postStore)	List(c context.Context, username string, offset,limit int) (int64, []model.Post, error) {
	var posts []model.Post
	var count int64
	err := p.db.Where("username = ?", username).Offset(offset).Limit(limit).Find(&posts).Offset(-1).Limit(-1).Count(&count).Error
	return count, posts, err
}
func (p *postStore) Delete(c context.Context, username string, postIds []string) error {
	return p.db.Where("username = ? and postId in (?)", username, postIds).Delete(&model.Post{}).Error
}