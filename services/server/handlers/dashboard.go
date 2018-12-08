package handlers

import (
	restful "github.com/emicklei/go-restful"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	protobuf "github.com/ilovelili/dongfeng-protobuf"
)

// Dashboard dashboard
func Dashboard(req *restful.Request, rsp *restful.Response) {
	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newcoreclient().Dashboard(ctx(req), &protobuf.DashboardRequest{Token: idtoken})
	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
	} else {
		rsp.WriteAsJson(response)
	}
}
