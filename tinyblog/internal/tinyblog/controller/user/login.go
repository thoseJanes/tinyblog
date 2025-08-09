package user

import (
	"log"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/thoseJanes/tinyblog/internal/pkg/core"
	"github.com/thoseJanes/tinyblog/internal/pkg/errno"

	v1 "github.com/thoseJanes/tinyblog/pkg/api/tinyblog/v1"
)

func (ctrl *UserController) Login(c *gin.Context) {
	var req v1.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("bind failed %#v", c.Request.Body)
		core.WriteResponse(c, err, nil)
		return
	}
	
	if _, err := govalidator.ValidateStruct(&req); err != nil {
		log.Print("valid failed")
		core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage("%s", err.Error()), nil)
		return
	}

	resp, err := ctrl.b.User().Login(c, &req)
	if err != nil {
		log.Print("login failed")
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, resp)
}