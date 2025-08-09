package ai

import (
	"github.com/gin-gonic/gin"
	"github.com/thoseJanes/tinyblog/internal/pkg/core"
	v1 "github.com/thoseJanes/tinyblog/pkg/api/tinyblog/v1"
	pb "github.com/thoseJanes/tinyblog/pkg/proto/aiservice/v1"
)


func (ctrl *AiController) GenerateTitle(c *gin.Context) {
	var in pb.PromptContentRequest
	in.Prompt = c.Query("prompt")
	in.Content = c.Query("content")

	out, err := ctrl.c.GenerateTitleAndTag(c, &in)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	var resp v1.GenerateTitleResponse
	resp.Title = out.Title
	core.WriteResponse(c, nil, resp)
}