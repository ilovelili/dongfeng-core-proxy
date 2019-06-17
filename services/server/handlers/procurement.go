package handlers

import (
	restful "github.com/emicklei/go-restful"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// GetProcurement get ingredient procurement based on unit amount and attandance
func GetProcurement(req *restful.Request, rsp *restful.Response) {
	from, to := req.QueryParameter("from"), req.QueryParameter("to")

	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newcoreclient().GetProcurements(ctx(req), &proto.GetProcurementRequest{
		Token: idtoken,
		From:  from,
		To:    to,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}
