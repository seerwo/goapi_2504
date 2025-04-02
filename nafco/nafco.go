package nafco

import (
	"github.com/seerwo/goapi_2504/credential"
	"github.com/seerwo/goapi_2504/nafco/config"
	"github.com/seerwo/goapi_2504/nafco/context"
	"github.com/seerwo/goapi_2504/nafco/data"
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

func (nf *Nafco) GetVersion() *data.Version {
	return data.NewVersion(nf.ctx)
}

func (nf *Nafco) GetSend() *data.Send {
	return data.NewSend(nf.ctx)
}

func (nf *Nafco) GetReceive() *data.Receive {
	return data.NewReceive(nf.ctx)
}