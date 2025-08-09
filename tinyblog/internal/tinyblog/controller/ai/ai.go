package ai

import pb "github.com/thoseJanes/tinyblog/pkg/proto/aiservice/v1"


type AiController struct {
	c pb.AIServiceClient
}


func New(c pb.AIServiceClient) *AiController {
	return &AiController{c: c}
}