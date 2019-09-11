package handlers

import (
	"fmt"

	restful "github.com/emicklei/go-restful"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// Login login
func Login(req *restful.Request, rsp *restful.Response) {
	_, pid, _ := utils.ResolveHeaderInfo(req)
	response, err := newcoreclient().Login(ctx(req), &proto.LoginRequest{
		Pid: pid,
	})
	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

// Logout logout
func Logout(req *restful.Request, rsp *restful.Response) {
	idtoken, _, _ := utils.ResolveHeaderInfo(req)
	// save token to blacklist
	err := redisclient.Set(fmt.Sprintf("%s_%s", sessionkeyprefix, idtoken), "_" /*value does not matter*/, 0 /* no ttl */).Err()
	if err != nil {
		writeError(rsp, errorcode.CoreProxyFailedToSaveSession)
		return
	}

	rsp.WriteAsJson(struct {
		Succeed bool `json:"succeed"`
	}{true})
}
