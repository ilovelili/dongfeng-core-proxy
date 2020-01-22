package handlers

import (
	restful "github.com/emicklei/go-restful"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// Dashboard dashboard
func Dashboard(req *restful.Request, rsp *restful.Response) {
	_, pid, email, _ := utils.ResolveHeaderInfo(req)
	response, err := newcoreclient().Dashboard(ctx(req), &proto.DashboardRequest{
		Pid:   pid,
		Email: email,
	})
	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
	} else {
		rsp.WriteAsJson(response)
	}
}
