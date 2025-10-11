package grpc

import (
	"context"
	"protos"
	"usrsvc/dao"

	"github.com/spf13/cast"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct {
	protos.UnimplementedUserServiceServer
}

func (u *UserServer) GetUserInfo(ctx context.Context, user *protos.User) (*protos.User, error) {
	//获取user信息
	info := dao.UserDao.Get(ctx, user.GetId())
	if len(info) == 0 {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}
	return &protos.User{
		Name: cast.ToString(info["name"]),
		Id:   cast.ToInt64(info["id"]),
	}, nil
}
