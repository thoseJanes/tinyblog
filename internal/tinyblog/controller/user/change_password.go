package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/thoseJanes/tinyblog/internal/pkg/core"
	"github.com/thoseJanes/tinyblog/internal/pkg/errno"
	v1 "github.com/thoseJanes/tinyblog/pkg/api/tinyblog/v1"
)


func (ctrl *UserController) ChangePassword(c *gin.Context) {
	var req v1.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	if _, err := govalidator.ValidateStruct(&req); err != nil {
		core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage("%s", err.Error()), nil)
		return
	}

	if err := ctrl.b.User().ChangePassword(c, c.GetString(core.XUsernameKey), &req); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, nil)
}