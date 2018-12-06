package handlers

import (
	"fmt"

	restful "github.com/emicklei/go-restful"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	protobuf "github.com/ilovelili/dongfeng-protobuf"
)

// Login login
func Login(req *restful.Request, rsp *restful.Response) {
	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newclient().Login(ctx(req), &protobuf.LoginRequest{Token: idtoken})
	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
	} else {
		rsp.WriteAsJson(response)
	}
}

// Logout logout
func Logout(req *restful.Request, rsp *restful.Response) {
	idtoken, _ := utils.ResolveIDToken(req)
	// save token to blacklist
	err := redisclient.Set(fmt.Sprintf("%s_%s", sessionkeyprefix, idtoken), "_" /*value does not matter*/, 0 /* no ttl */).Err()
	if err != nil {
		writeError(rsp, errorcode.CoreProxyFailedToSaveSession)
		return
	}
	rsp.Write([]byte("OK"))
}
