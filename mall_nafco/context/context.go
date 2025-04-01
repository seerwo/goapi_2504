package context

import (
	"github.com/seerwo/goapi_2504/credential"
	"github.com/seerwo/goapi_2504/mall_nafco/config"
)

type Context struct {
	*config.Config
	credential.AccessTokenHandle
}

