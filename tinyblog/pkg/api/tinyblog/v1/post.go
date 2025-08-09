package v1

type PostInfo struct{
	Title string `json:"title"`
	Content string `json:"content"`
	PostId string `json:"postId"`
	Username string `json:"username"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type ListPostRequest struct{
	Offset int `form:"offset"`
	Limit int `form:"limit"`
}

type ListPostResponse struct{
	Posts []*PostInfo `json:"posts"`
	TotalCount int64 `json:"count"`
}

type GetPostResponse PostInfo

type UpdatePostRequest struct {
	Title *string `json:"title" valid:"stringlength(1|255)"`
	Content *string `json:"content" valid:"stringlength(0|16384)"`
	// PostId string `json:"postId" valid:"required"`,postId参数是从url中获取的。
}

type CreatePostRequest struct {
	Title *string `json:"title" valid:"stringlength(1|255)"`
	Content *string `json:"content" valid:"stringlength(0|16384)"`
}

type CreatePostResponse struct {
	PostId string `json:"postId"`
}

type DeletePostRequest struct {
	PostId string `json:"postId" valid:"required"`
}

