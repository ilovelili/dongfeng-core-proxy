package handlers

import (
	restful "github.com/emicklei/go-restful"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// GetMenu get menu
func GetMenu(req *restful.Request, rsp *restful.Response) {
	from, to := req.PathParameter("from"), req.PathParameter("to")
	if from > to {
		writeError(rsp, errorcode.CoreProxyInvalidGetMenuRequest)
		return
	}

	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newnutritionclient().GetMenu(ctx(req), &proto.GetMenuRequest{
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
