package svc

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhanghongliang12/lebron/apps/user/rpc/internal/config"
	"github.com/zhanghongliang12/lebron/apps/user/rpc/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	fmt.Println(c.Mysql.Datasource)
	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	fmt.Println("sqlConn------->", sqlConn)
	//fmt.Println("CacheRedis---->", c.CacheRedis)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlConn, c.CacheRedis),
	}
}
