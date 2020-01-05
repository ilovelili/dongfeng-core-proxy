package handlers

import (
	restful "github.com/emicklei/go-restful"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// GetRole get role
func GetRole(req *restful.Request, rsp *restful.Response) {
	_, pid, _ := utils.ResolveHeaderInfo(req)
	response, err := newcoreclient().GetRole(ctx(req), &proto.GetRoleRequest{
		Pid: pid,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}
