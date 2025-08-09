package post

import(
	"context"
	"github.com/thoseJanes/tinyblog/pkg/api/tinyblog/v1"
)

//go:generate gencode ./interface.go -o post.go -t i -r PostBiz:postBiz
type PostBiz interface {
	Create(ctx context.Context, username string, r *v1.CreatePostRequest) (*v1.CreatePostResponse, error)
	Update(ctx context.Context, username, postId string, r *v1.UpdatePostRequest) error
	Get(ctx context.Context, username, postId string) (*v1.GetPostResponse, error)
	Search(ctx context.Context, r *v1.SearchPostRequest) (*v1.SearchPostResponse, error)
	AiSearch(ctx context.Context, r *v1.AiSearchPostRequest) (*v1.AiSearchPostResponse, error)
	Delete(ctx context.Context, username, postId string) error
	List(ctx context.Context, username string, r *v1.ListPostRequest) (*v1.ListPostResponse, error)
}


