package handlers

import (
	restful "github.com/emicklei/go-restful"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	protobuf "github.com/ilovelili/dongfeng-protobuf"
)

// GetNamelist get name list
func GetNamelist(req *restful.Request, rsp *restful.Response) {
	class, year := req.QueryParameter("class"), req.QueryParameter("year")

	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newcoreclient().GetNamelist(ctx(req), &protobuf.GetNamelistRequest{
		Token: idtoken,
		Year:  year,
		Class: class,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}
