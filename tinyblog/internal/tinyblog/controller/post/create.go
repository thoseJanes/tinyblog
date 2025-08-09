package post

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/thoseJanes/tinyblog/internal/pkg/core"
	"github.com/thoseJanes/tinyblog/internal/pkg/errno"
	v1 "github.com/thoseJanes/tinyblog/pkg/api/tinyblog/v1"
)


func (ctrl *PostController) Create(c *gin.Context) {
	var req v1.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	if _, err := govalidator.ValidateStruct(&req); err != nil {
		core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage("%s", err.Error()), nil)
		return
	}

	resp, err := ctrl.b.Post().Create(c, c.GetString(core.XUsernameKey), &req)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, resp)
}