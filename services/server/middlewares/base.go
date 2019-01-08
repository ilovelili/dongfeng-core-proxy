package middlewares

import (
	restful "github.com/emicklei/go-restful"
	"github.com/go-redis/redis"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
)

var (
	config      = utils.GetConfig()
	jwks        = config.Auth.JWKS
	redisclient = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Host,
		Password: config.Redis.Password,
	})
)

// sessionkey string used as Redis session store key
const sessionkeyprefix = "session"

func writeError(rsp *restful.Response, errorcode *errorcode.Error, detail ...string) {
	e := utils.NewError(errorcode, detail...)
	rsp.WriteError(int(errorcode.Code), e)
}
