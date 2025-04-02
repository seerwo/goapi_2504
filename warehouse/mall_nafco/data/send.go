package data

import "github.com/seerwo/goapi_2504/mall_nafco/context"

const(
	SEND_ULR = "https://www.vinas.jp/web-edi/transfer/Receive.do"
	SEND_URL_WATCH = "https://www.vinas.jp/web-edi/transfer/ProgressReceive.do"
)

type Send struct {
	*context.Context
}

func NewSend(context *context.Context) *Send {
	send := new(Send)
	send.Context = context
	return send
}