package data

import (
	"fmt"
	"github.com/seerwo/goapi_2504/mall_nafco/context"
	"github.com/seerwo/goapi_2504/util"
)

const(
	url = "https://www.vinas.jp/web-edi/transfer/Send.do"
	urlUpd = "https://www.vinas.jp/web-edi/transfer/SendUpdate.do"
	urlDL = "https://www.vinas.jp/web-edi/transfer/SendDownLoad.do"
	urlWatch = "https://www.vinas.jp/web-edi/transfer/ProgressSend.do"
	urlCancel = "https://www.vinas.jp/web-edi/transfer/CancelSend.do"
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

	response, err := util.PostJSON(url, nil)
	if err != nil {
		return err
	}

	return util.DecodeWithCommonError(response, "DeleteConditional")
}