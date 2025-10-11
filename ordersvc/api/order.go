package api

import (
	"context"
	"dogapm"
	"net/http"
	"ordersvc/grpcclient"
	"protos"
	"strconv"

	"github.com/google/uuid"
)

type order struct {
}

var Order = &order{}

func (o *order) Add(w http.ResponseWriter, request *http.Request) {
	//从URL中获取参数
	values := request.URL.Query()
	var (
		uid, _   = strconv.Atoi(values.Get("uid"))
		skuid, _ = strconv.Atoi(values.Get("sku_id"))
		num, _   = strconv.Atoi(values.Get("num"))
	)
	//检查用户信息
	_, err := grpcclient.UserClient.GetUserInfo(context.TODO(), &protos.User{
		Id: int64(uid),
	})
	if err != nil {
		dogapm.Logger.Error(context.TODO(), "checkUser", map[string]interface{}{
			"uid":    uid,
			"sku_id": skuid,
		}, err)
		dogapm.HttpStatus.Error(w, err.Error(), nil)
		return
	}

	//库存调整
	skuMsg, err := grpcclient.SkuClient.DecreaseStock(context.TODO(), &protos.Sku{
		Id:  int64(skuid),
		Num: int32(num),
	})
	if err != nil {
		dogapm.Logger.Error(context.TODO(), "createOrder", map[string]interface{}{
			"uid":    uid,
			"sku_id": skuid,
		}, err)
		dogapm.HttpStatus.Error(w, err.Error(), nil)
		return
	}
	//创建订单
	_, err = dogapm.Infra.Db.ExecContext(context.TODO(), "insert into t_order(order_id, sku_id, num, price, uid) values ( ?, ?, ?, ?, ?)",
		uuid.New().String(), skuid, num, int(skuMsg.Price)*num, uid)
	if err != nil {
		dogapm.Logger.Error(context.TODO(), "createOrder", map[string]interface{}{
			"uid":    uid,
			"sku_id": skuid,
		}, err)
		dogapm.HttpStatus.Error(w, err.Error(), nil)
	}
	dogapm.HttpStatus.OK(w)

}
