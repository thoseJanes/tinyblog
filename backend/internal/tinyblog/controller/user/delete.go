package user

import (
	"github.com/gin-gonic/gin"
	"github.com/thoseJanes/tinyblog/internal/pkg/core"
)


func (ctrl *UserController) Delete(c *gin.Context) {

	
	if err := ctrl.b.User().Delete(c, c.Param("name")); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	if _, err := ctrl.a.Enforcer.RemoveNamedPolicy("p", c.Param("name"), "", ""); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, nil)
}