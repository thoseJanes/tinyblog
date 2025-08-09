package ai

import (
	"io"

	"github.com/gin-gonic/gin"
	"github.com/thoseJanes/tinyblog/internal/pkg/core"
	pb "github.com/thoseJanes/tinyblog/pkg/proto/aiservice/v1"
)


func (ctrl *AiController) PolishContent(c *gin.Context) {
	var in pb.PromptContentRequest
	in.Prompt = c.Query("prompt")
	in.Content = c.Query("content")

	stream, err := ctrl.c.PolishContent(c, &in)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	c.Header("Content-Type", "text/event-stream")
    c.Header("Cache-Control", "no-cache")
    c.Header("Connection", "keep-alive")

	for{
		chunk, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				c.SSEvent("end", "stream completed")
				break;
			}
			core.WriteResponse(c, err, nil)
			return
		}

		c.SSEvent("message", chunk.ContentChunk)
		c.Writer.Flush()
	}
}