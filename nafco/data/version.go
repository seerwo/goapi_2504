package data

import (
	"fmt"
	"github.com/seerwo/goapi_2504/nafco/context"
	"github.com/seerwo/goapi_2504/util"
)

const(
	VERSION_URL = "https://www.vinas.jp/web-edi/transfer/Version.do"
	VERION_URL_WATCH = "https://www.vinas.jp/web-edi/transfer/ProgressVersion.do"
	VERION_URL_DL = "https://www.vinas.jp/web-edi/transfer/VersionDownLoad.do"
)

type Version struct {
	*context.Context
}

func NewVersion(context *context.Context) *Version {
	version := new(Version)
	version.Context = context
	return version
}

func (version *Version) GetVersion() error {
	accessToken, err := version.GetAccessToken()
	if err != nil {
		return err
	}
	fmt.Printf(accessToken)
	//uri := fmt.Sprintf("%", menuDeleteConditionalURL, accessToken)
	//reqDeleteConditional := &reqDeleteConditional{
	//	MenuID: menuID,
	//}

	response, err := util.PostJSON(VERSION_URL, nil)
	if err != nil {
		return err
	}

	return util.DecodeWithCommonError(response, "DeleteConditional")
}

func (version *Version) WatchState() error {
	accessToken, err := version.GetAccessToken()
	if err != nil {
		return err
	}
	fmt.Printf(accessToken)
	//uri := fmt.Sprintf("%", menuDeleteConditionalURL, accessToken)
	//reqDeleteConditional := &reqDeleteConditional{
	//	MenuID: menuID,
	//}

	response, err := util.PostJSON(VERION_URL_WATCH, nil)
	if err != nil {
		return err
	}

	return util.DecodeWithCommonError(response, "DeleteConditional")
}

func (version *Version) Download() error {
	accessToken, err := version.GetAccessToken()
	if err != nil {
		return err
	}
	fmt.Printf(accessToken)
	//uri := fmt.Sprintf("%", menuDeleteConditionalURL, accessToken)
	//reqDeleteConditional := &reqDeleteConditional{
	//	MenuID: menuID,
	//}

	response, err := util.PostJSON(VERION_URL_DL, nil)
	if err != nil {
		return err
	}

	return util.DecodeWithCommonError(response, "DeleteConditional")
}