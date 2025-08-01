package post

import (
	"github.com/gin-gonic/gin"
	"github.com/thoseJanes/tinyblog/internal/pkg/core"
)


func (ctrl *PostController) Get(c *gin.Context) {
	resp, err := ctrl.b.Post().Get(c, c.GetString(core.XUsernameKey), c.Param("postId"))
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, resp)
}