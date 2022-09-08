package svc

import (
	"wiki-knogen-zero/internal/config"
	"wiki-knogen-zero/service/freshstorage"
	"wiki-knogen-zero/service/translate"
)

type ServiceContext struct {
	Config         config.Config
	FreshS         *freshstorage.FreshStorage
	BaiduTranslate *translate.Translate
}

func NewServiceContext(c config.Config) *ServiceContext {
	baiduTranslate := translate.NewTranslate(
		c.BaiduVipTranslateAppid,
		c.BaiduVipTranslateSecret,
		c.PgConnUri,
	)

	freshS := freshstorage.NewFreshStorage(c.PgConnUri)
	return &ServiceContext{
		Config:         c,
		FreshS:         freshS,
		BaiduTranslate: baiduTranslate,
	}
}
