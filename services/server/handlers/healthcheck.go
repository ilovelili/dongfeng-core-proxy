package handlers

import (
	restful "github.com/emicklei/go-restful"
)

// HealthCheck do healthcheck
func HealthCheck(req *restful.Request, rsp *restful.Response) {
	rsp.WriteAsJson(struct {
		Healthy bool `json:"healthy"`
	}{true})
}
