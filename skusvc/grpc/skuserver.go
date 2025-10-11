package grpc

import (
	"context"
	"protos"
	"skusvc/dao"

	"github.com/spf13/cast"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SkuServer struct {
	protos.UnimplementedSkuServiceServer
}

func (s *SkuServer) DecreaseStock(ctx context.Context, sku *protos.Sku) (*protos.Sku, error) {
	//获取商品信息
	info := dao.SkuDao.Get(ctx, sku.Id)
	if len(info) == 0 {
		return nil, status.Errorf(codes.NotFound, "sku not found")
	}
	//扣减库存
	decrRes, err := dao.SkuDao.Decr(ctx, sku.Id, sku.Num)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, err.Error())
	}
	if affected, _ := decrRes.RowsAffected(); affected == 0 {
		return nil, status.Errorf(codes.PermissionDenied, "no enough sku")
	}
	return &protos.Sku{
		Name:  cast.ToString(info["name"]),
		Id:    cast.ToInt64(info["id"]),
		Price: cast.ToInt32(info["price"]),
		Num:   cast.ToInt32(info["num"]) - sku.Num,
	}, nil
}
