package middlewares

import (
	"fmt"

	restful "github.com/emicklei/go-restful"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	"github.com/ilovelili/dongfeng-shared-lib"
)

// JwtAuthenticate JWT auth middleware.
func JwtAuthenticate(req *restful.Request, rsp *restful.Response, chain *restful.FilterChain) {
	idtoken, valid := utils.ResolveIDToken(req)
	if !valid {
		writeError(rsp, errorcode.GenericNotAuthorized)
		return
	}

	// check if idtoken in blacklist
	exist, err := redisclient.Exists(fmt.Sprintf("%s_%s", sessionkeyprefix, idtoken)).Result()
	if err != nil {
		writeError(rsp, errorcode.CoreProxyFailedToGetSession)
		return
	}

	// token in blacklist
	if exist > 0 {
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
