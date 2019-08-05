package middlewares

import (
	"encoding/json"
	"fmt"

	restful "github.com/emicklei/go-restful"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	sharedlib "github.com/ilovelili/dongfeng-shared-lib"
)

// UserVerification user verification is used to check token mismatch
type UserVerification struct {
	Token string
}

// JwtAuthenticate JWT auth middleware.
func JwtAuthenticate(req *restful.Request, rsp *restful.Response, chain *restful.FilterChain) {
	idtoken, pid, valid := utils.ResolveHeaderInfo(req)
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

	// verify login
	authClient := sharedlib.NewAuthClient(config.Auth.ClientID, config.Auth.ClientSecret)
	status, err := authClient.VerifyLogin(idtoken)
	if err != nil || !status {
		writeError(rsp, errorcode.GenericNotAuthorized)
		return
	}

	// parse user info
	userinfo, err := authClient.ParseUserInfo(pid)
	if err != nil {
		writeError(rsp, errorcode.GenericNotAuthorized)
		return
	}

	var userverification *UserVerification
	err = json.Unmarshal(userinfo, &userverification)
	if err != nil {
		writeError(rsp, errorcode.GenericNotAuthorized)
		return
	}

	// token mismatch, which is a hacking attemp
	if userverification.Token != idtoken {
		writeError(rsp, errorcode.GenericNotAuthorized)
		return
	}

	chain.ProcessFilter(req, rsp)
}
