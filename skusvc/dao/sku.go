package dao

import (
	"context"
	"database/sql"
	"dogapm"
)

type skuDao struct {
}

var SkuDao = &skuDao{}

func (s *skuDao) Get(ctx context.Context, id int64) map[string]interface{} {
	info := dogapm.DBUtil.QueryFirst(dogapm.Infra.Db.QueryContext(ctx, "select * from t_sku where id = ?;", id))
	if len(info) == 0 {
		return nil
	}
	return info
}

func (s *skuDao) Decr(ctx context.Context, id int64, num int32) (sql.Result, error) {
	res, err := dogapm.Infra.Db.ExecContext(ctx, "update t_sku set num = num - ? where id = ? and (num-?) >= 0", num, id, num)
	if err != nil {
		return nil, err
	}
	return res, err
}
