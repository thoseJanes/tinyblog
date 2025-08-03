/*
可以基于RESTful来考虑。
获取信息（GetInfo）使用GET，因此不需要Request，直接从Url的Path中提取。
创建信息（Create）、修改信息（Update、ChangePassword）使用PUT（幂等创建资源）、PATCH，直接回复成功与否。
登陆信息（Login）使用POST，获取令牌
*/
package v1



type LoginRequest struct {
	Username string `json:"username" valid:"alphanum,required,stringlength(1|255)"`
	Password string `json:"password" valid:"required,stringlength(6|18)"`
}

type LoginResponse struct {
	Token string `json:"token"`
}


type CreateUserRequest struct {
	Username string `json:"username" valid:"alphanum,required,stringlength(1|255)"`
	Password string `json:"password" valid:"required,stringlength(6|18)"`
	Nickname string `json:"nickname" valid:"required,stringlength(1|255)"`
	Phone *string `json:"phone" valid:"stringlength(11|11)"`
	Email *string `json:"email" valid:"email"`
}


type UserInfo struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	PostCount int64 `json:"postCount"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type GetUserResponse UserInfo


type UpdateUserRequest struct {
	Nickname *string `json:"nickname" valid:"stringlength(1|255)"`
	Phone *string `json:"phone" valid:"stringlength(11|11)"`
	Email *string `json:"email" valid:"email"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"oldPassword" valid:"required,stringlength(6|18)"`
	NewPassword string `json:"newPassword" valid:"required,stringlength(6|18)"`
}



type ListUserRequest struct {
	Offset int `form:"offset"`
	Limit int `form:"limit"`
}

type ListUserResponse struct {
	Users []*UserInfo `json:"users"`
	TotalCount int64 `json:"count"`
}