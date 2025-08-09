package post

import (
	"context"
	"errors"

	"github.com/jinzhu/copier"
	"github.com/thoseJanes/tinyblog/internal/pkg/errno"
	"github.com/thoseJanes/tinyblog/internal/pkg/model"
	"github.com/thoseJanes/tinyblog/internal/tinyblog/store"
	"github.com/thoseJanes/tinyblog/pkg/api/tinyblog/v1"
	"gorm.io/gorm"
)

//implement
//Create(ctx context.Context, username string, r *v1.CreatePostRequest) (*v1.CreatePostResponse, error)
//Update(ctx context.Context, username string, r *v1.UpdatePostRequest) error
//Get(ctx context.Context, username string, postId string) (*v1.GetPostResponse, error)
//Delete(ctx context.Context, username string, postId string) error
//List(ctx context.Context, username string, offset int, limit int) (*v1.ListPostResponse, error)
type postBiz struct {
	ds store.IStore
}

var _ PostBiz = (*postBiz)(nil)


func New(ds store.IStore) PostBiz {
	return &postBiz{ds}
}


func (p *postBiz) Create(ctx context.Context, username string, r *v1.CreatePostRequest) (*v1.CreatePostResponse, error) {
	var postM model.Post
	copier.Copy(&postM, r)
	postM.Username = username
	if err := p.ds.PostStore().Create(ctx, &postM); err != nil {
		return nil, err
	}

	return &v1.CreatePostResponse{PostId: postM.PostId}, nil
}


func (p *postBiz) Update(ctx context.Context, username, postId string, r *v1.UpdatePostRequest) error {
	var postM model.Post
	copier.Copy(&postM, r)
	postM.PostId = postId
	postM.Username = username
	if err := p.ds.PostStore().Update(ctx, &postM); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errno.ErrPostNotFound
		}
		return err
	}

	return nil
}


func (p *postBiz) Get(ctx context.Context, username string, postId string) (*v1.GetPostResponse, error) {
	post, err := p.ds.PostStore().Get(ctx, username, postId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.ErrPostNotFound
		}
		return nil, err
	}

	resp := v1.GetPostResponse{
		Title: post.Title,
		Content: post.Content,
		PostId: post.PostId,
		Username: post.Username,
		CreatedAt: post.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: post.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	return &resp, nil
}


func (p *postBiz) Delete(ctx context.Context, username string, postId string) error {
	err := p.ds.PostStore().Delete(ctx, username, []string{postId})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}

	return nil
}


func (p *postBiz) List(ctx context.Context, username string, r *v1.ListPostRequest) (*v1.ListPostResponse, error) {
	tolPost, posts, err := p.ds.PostStore().List(ctx, username, r.Offset, r.Limit)
	if err != nil {
		return nil, err
	}

	postsInfo := make([]*v1.PostInfo, 0, len(posts))
	for _, post := range posts {
		postsInfo = append(postsInfo, &v1.PostInfo{
			Title: post.Title,
			Content: post.Content,
			PostId: post.PostId,
			Username: post.Username,
			CreatedAt: post.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: post.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.ListPostResponse{TotalCount: tolPost, Posts: postsInfo}, nil
}



