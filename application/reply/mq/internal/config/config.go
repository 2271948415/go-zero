package config

import (
	"github.com/zeromicro/go-queue/kq"
)

type Config struct {
	KqConsumerConf kq.KqConf
	DB             struct {
		DataSource   string
		MaxOpenConns int `json:",default=10"`
		MaxIdleConns int `json:",default=100"`
		MaxLifetime  int `json:",default=3600"`
	}
}
