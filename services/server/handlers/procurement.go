package handlers

import (
	"encoding/json"

	restful "github.com/emicklei/go-restful"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// ProcurementRequestItem procurement request
type ProcurementRequestItem struct {
	ID     int64   `json:"id"`
	Amount float64 `json:"amount"`
}

// GetProcurements get ingredient procurement based on unit amount and attandance
func GetProcurements(req *restful.Request, rsp *restful.Response) {
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

// UpdateProcurement update procurement
func UpdateProcurement(req *restful.Request, rsp *restful.Response) {
	decoder := json.NewDecoder(req.Request.Body)
	var updatereq *ProcurementRequestItem
	err := decoder.Decode(&updatereq)
	if err != nil {
		writeError(rsp, errorcode.CoreProxyInvalidProcurementUpdateRequest)
		return
	}

	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newcoreclient().UpdateProcurement(ctx(req), &proto.UpdateProcurementRequest{
		Token:  idtoken,
		Id:     updatereq.ID,
		Amount: updatereq.Amount,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}
