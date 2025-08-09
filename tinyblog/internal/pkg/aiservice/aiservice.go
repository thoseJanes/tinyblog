package aiservice

import (
	pb "github.com/thoseJanes/tinyblog/pkg/proto/aiservice/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var Client pb.AIServiceClient

func InitClient(address string) error {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	Client = pb.NewAIServiceClient(conn)
	return nil
}