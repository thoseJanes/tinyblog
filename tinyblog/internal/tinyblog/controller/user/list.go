package user

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/thoseJanes/tinyblog/internal/pkg/core"
	v1 "github.com/thoseJanes/tinyblog/pkg/api/tinyblog/v1"
	pb "github.com/thoseJanes/tinyblog/pkg/proto/tinyblog/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)



func (ctrl *UserController) List(c *gin.Context) {
	var req v1.ListUserRequest
	if err := c.ShouldBind(&req); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	fmt.Printf("%#v", req)

	resp, err := ctrl.b.User().List(c, req.Offset, req.Limit)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return 
	}

	core.WriteResponse(c, nil, resp)
}


func (ctrl *UserController) ListUser(c context.Context, r *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	resp, err := ctrl.b.User().List(c, int(r.Offset), int(r.Limit))
	if err != nil {
		return nil, err
	}

	var users = make([]*pb.UserInfo, len(resp.Users))
	for i, u := range resp.Users {
		createdAt, _ := time.Parse("2006-01-02 15:04:05", u.CreatedAt)
		updatedAt, _ := time.Parse("2006-01-02 15:04:05", u.UpdatedAt)
		copier.Copy(users[i], u)
		// users[i] = &pb.UserInfo{
		// 	Username: u.Username,
		// 	Nickname: u.Nickname,
		// 	Email: u.Email,
		// 	Phone: u.Phone,
		// 	PostCount: u.PostCount,
		// 	CreatedAt: timestamppb.New(createdAt),
		// 	UpdatedAt: timestamppb.New(updatedAt),
		// }
		users[i].CreatedAt = timestamppb.New(createdAt)
		users[i].UpdatedAt = timestamppb.New(updatedAt)
	}

	return &pb.ListUserResponse{TotalCount: resp.TotalCount, Users: users}, nil
}