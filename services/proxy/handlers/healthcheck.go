package handlers

import (
	"net/http"

	restful "github.com/emicklei/go-restful"
	api "github.com/ilovelili/dongfeng/core-proxy/services/proto"
)

// HealthCheck do healthcheck
func HealthCheck(req *restful.Request, rsp *restful.Response) {
	response, err := newclient().Healthcheck(req.Request.Context(), &api.HealthcheckRequest{})

	if err != nil {
		rsp.WriteError(http.StatusInternalServerError, err)
	}

	rsp.WriteEntity(map[string]string{
		"healthy": response.Message,
	})
}
