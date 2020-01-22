package middlewares

import (
	restful "github.com/emicklei/go-restful"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	sharedlib "github.com/ilovelili/dongfeng-shared-lib"
)

// JwtAuthenticate JWT auth middleware.
func JwtAuthenticate(req *restful.Request, rsp *restful.Response, chain *restful.FilterChain) {
	idtoken, _, _, valid := utils.ResolveHeaderInfo(req)
	if !valid {
		writeError(rsp, errorcode.GenericNotAuthorized)
		return
	}

	// verify login
	authClient := sharedlib.NewAuthClient(config.Auth.ClientID, config.Auth.ClientSecret)
	status, err := authClient.VerifyLogin(idtoken)
	if err != nil || !status {
		writeError(rsp, errorcode.GenericNotAuthorized)
		return
	}

	chain.ProcessFilter(req, rsp)
}
