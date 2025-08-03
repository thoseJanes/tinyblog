package user

import (
	"github.com/gin-gonic/gin"
	"github.com/thoseJanes/tinyblog/internal/pkg/core"
)

func (ctrl *UserController) Get(c *gin.Context) {
	resp, err := ctrl.b.User().Get(c, c.GetString(core.XUsernameKey))
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, resp)
}