package svc

import (
	"gozero/application/like/mq/internal/config"
	"gozero/application/like/mq/internal/model"
	"gozero/pkg/orm"
)

type ServiceContext struct {
	Config         config.Config
	DB             *orm.DB
	LikeModel      *model.LikeModel
	LikeCountModel *model.LikeCountModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := orm.MustNewMysql(&orm.Config{
		DSN:          c.DB.DataSource,
		MaxOpenCoons: c.DB.MaxOpenConns,
		MaxIdleCoons: c.DB.MaxIdleConns,
		MaxLifetime:  c.DB.MaxLifetime,
	})
	return &ServiceContext{
		Config:         c,
		DB:             db,
		LikeModel:      model.NewLikeModel(db.DB),
		LikeCountModel: model.NewLikeCountModel(db.DB),
	}
}
