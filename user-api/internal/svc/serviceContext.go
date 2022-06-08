package svc

import (
	"go-zore/user-api/internal/config"
	"go-zore/user-api/internal/middleware"
	"go-zore/user-api/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config config.Config
	// kafka kafka.Server
	UserModel      model.UserModel
	TestMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		TestMiddleware: middleware.NewTestMiddleware().Handle,
		// kafka:kafka.New(url,pass),
		// redis:redis.New(url,pass),
		UserModel: model.NewUserModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
