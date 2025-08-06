package post

import (
	"github.com/gin-gonic/gin"
	"github.com/thoseJanes/tinyblog/internal/pkg/core"
	v1 "github.com/thoseJanes/tinyblog/pkg/api/tinyblog/v1"
)

func (ctrl *PostController) List(c *gin.Context) {
	var req v1.ListPostRequest
	if err := c.ShouldBind(&req); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	resp, err := ctrl.b.Post().List(c, c.GetString(core.XUsernameKey), &req)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, resp)
	
}