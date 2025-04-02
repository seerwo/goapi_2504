package data

import (
	"fmt"
	"github.com/seerwo/goapi_2504/nafco/context"
	"github.com/seerwo/goapi_2504/util"
)

const(
	RECEIVE_URL = "https://www.vinas.jp/web-edi/transfer/Send.do"
	RECEIVE_URL_UPD = "https://www.vinas.jp/web-edi/transfer/SendUpdate.do"
	RECEIVE_URL_DL = "https://www.vinas.jp/web-edi/transfer/SendDownLoad.do"
	RECEIVE_URL_WATCH = "https://www.vinas.jp/web-edi/transfer/ProgressSend.do"
	RECEIVE_URL_CANCEL = "https://www.vinas.jp/web-edi/transfer/CancelSend.do"
)

type Receive struct {
	*context.Context
}

func NewReceive(context *context.Context) *Receive {
	receive := new(Receive)
	receive.Context = context
	return receive
}

func (receive *Receive) Post() error {
	accessToken, err := receive.GetAccessToken()
	if err != nil {
		return err
	}
	fmt.Printf(accessToken)
	//uri := fmt.Sprintf("%", menuDeleteConditionalURL, accessToken)
	//reqDeleteConditional := &reqDeleteConditional{
	//	MenuID: menuID,
	//}

	response, err := util.PostJSON(RECEIVE_URL, nil)
	if err != nil {
		return err
	}

	return util.DecodeWithCommonError(response, "DeleteConditional")
}