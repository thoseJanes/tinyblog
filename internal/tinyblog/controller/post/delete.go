package post

import (
	"github.com/gin-gonic/gin"
	"github.com/thoseJanes/tinyblog/internal/pkg/core"
)


func (ctrl *PostController) Delete(c *gin.Context) {
	err := ctrl.b.Post().Delete(c, c.GetString(core.XUsernameKey), c.Param("postId"))
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, nil)
}