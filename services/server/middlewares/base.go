package middlewares

import (
	restful "github.com/emicklei/go-restful"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
)

var (
	config = utils.GetConfig()
)

func writeError(rsp *restful.Response, errorcode *errorcode.Error, detail ...string) {
	e := utils.NewError(errorcode, detail...)
	rsp.WriteError(int(errorcode.Code), e)
}
