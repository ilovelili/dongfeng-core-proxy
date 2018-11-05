package middlewares

import (
	"github.com/boj/redistore"
	restful "github.com/emicklei/go-restful"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	"github.com/ilovelili/dongfeng-shared-lib"
)

func init() {
	sessionstore, _ = redistore.NewRediStore(config.Redis.GetMaxConnectionCount(), "tcp", config.Redis.Host, config.Redis.Password, []byte("session-store"))
}

// JwtAuthenticate JWT auth middleware. go-restful has poor support for middleware injection
func JwtAuthenticate(req *restful.Request, rsp *restful.Response, chain *restful.FilterChain) {
	idtoken, valid := utils.ResolveIDToken(req)
	if !valid {
		writeError(rsp, errorcode.GenericNotAuthorized)
		return
	}

	// check if idtoken in blacklist
	session, err := sessionstore.Get(req.Request, "login-session")
	if err != nil {
		writeError(rsp, errorcode.CoreProxyFailedToGetSession)
		return
	}
	flashes := session.Flashes(sessionkey)

	// token in blacklist
	if containString(flashes, idtoken) {
		writeError(rsp, errorcode.GenericNotAuthorized)
		return
	}

	// parse and validate the token
	_, token, err := sharedlib.ParseJWT(idtoken, jwks)
	if err != nil || !token.Valid {
		writeError(rsp, errorcode.GenericNotAuthorized)
		return
	}

	chain.ProcessFilter(req, rsp)
}
