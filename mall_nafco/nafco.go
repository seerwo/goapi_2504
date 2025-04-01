package mall_nafco

import (
	"github.com/seerwo/goapi_2504/credential"
	"github.com/seerwo/goapi_2504/mall_nafco/config"
	"github.com/seerwo/goapi_2504/mall_nafco/context"
)

type Nafco struct {
	ctx *context.Context
}

func NewNafco(cfg *config.Config) *Nafco {
	defaultAkHandle := credential.NewWorkAccessToken(cfg.CorpID, cfg.CorpSecret, cfg.AgentID, credential.CacheKeyWorkPrefix, cfg.Cache)
	ctx := &context.Context{
		Config:            cfg,
		AccessTokenHandle: defaultAkHandle,
	}
	return &Nafco{ctx: ctx}
}

func (nf *Nafco) GetContext() *context.Context {
	return nf.ctx
}