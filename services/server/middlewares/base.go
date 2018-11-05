package middlewares

import (
	"github.com/boj/redistore"
	restful "github.com/emicklei/go-restful"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
)

var (
	config       = utils.GetConfig()
	jwks         = config.Auth.JWKS
	sessionstore *redistore.RediStore
)

// sessionkey string used as Redis session store key
const sessionkey = "blacklisttoken"

func writeError(rsp *restful.Response, errorcode *errorcode.Error, detail ...string) {
	e := utils.NewError(errorcode, detail...)
	rsp.WriteError(int(errorcode.Code), e)
}

func containString(s []interface{}, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
