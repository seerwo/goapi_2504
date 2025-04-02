package goapi_2504

import (
	"github.com/seerwo/goapi_2504/cache"
	"github.com/seerwo/goapi_2504/mall_nafco"
	mallConfig "github.com/seerwo/goapi_2504/mall_nafco/config"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

type GoApi2504 struct {
	cache cache.Cache
}

func NewGoApi2504() *GoApi2504 {
	return &GoApi2504{}
}

func (ga *GoApi2504) SetCache(cache cache.Cache) {
	ga.cache = cache
}

func (ga *GoApi2504) GetMallNafco(cfg *mallConfig.Config) *mall_nafco.Nafco {
	return mall_nafco.NewNafco(cfg)
}