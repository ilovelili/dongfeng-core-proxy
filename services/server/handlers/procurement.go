package handlers

import (
	"strconv"

	restful "github.com/emicklei/go-restful"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// GetProcurement get ingredient procurement based on unit amount and attandance
func GetProcurement(req *restful.Request, rsp *restful.Response) {
	from, to, id := req.QueryParameter("from"), req.QueryParameter("to"), req.QueryParameter("id")
	if id != "" && id != "0" && id != "1" && id != "2" {
		writeError(rsp, errorcode.CoreProxyInvalidGetProcurementRequest)
		return
	}

	var _id int64
	if id == "" {
		_id = -1
	} else {
		_id, _ = strconv.ParseInt(id, 10, 64)
	}

	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newcoreclient().GetProcurements(ctx(req), &proto.GetProcurementRequest{
		Token:  idtoken,
		From:   from,
		To:     to,
		MealId: _id,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}
