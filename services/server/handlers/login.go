package handlers

import (
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
