package handlers

import (
	"encoding/json"

	restful "github.com/emicklei/go-restful"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// ProfileTemplateReqItem profile template request
type ProfileTemplateReqItem struct {
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
}

// GetProfileTemplate get template
func GetProfileTemplate(req *restful.Request, rsp *restful.Response) {
	name := req.QueryParameter("name")
	_, pid, _ := utils.ResolveHeaderInfo(req)
	response, err := newcoreclient().GetProfileTemplate(ctx(req), &proto.GetProfileTemplateRequest{
		Pid:  pid,
		Name: name,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

// GetProfileTemplates get templates
func GetProfileTemplates(req *restful.Request, rsp *restful.Response) {
	_, pid, _ := utils.ResolveHeaderInfo(req)
	response, err := newcoreclient().GetProfileTemplates(ctx(req), &proto.GetProfileTemplatesRequest{
		Pid: pid,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

// UpdateProfileTemplate update profile template
func UpdateProfileTemplate(req *restful.Request, rsp *restful.Response) {
	decoder := json.NewDecoder(req.Request.Body)
	var profiletemplatereq *ProfileTemplateReqItem
	err := decoder.Decode(&profiletemplatereq)
	if err != nil {
		writeError(rsp, errorcode.CoreProxyInvalidProfileTemplateUpdateRequestBody)
		return
	}

	_, pid, _ := utils.ResolveHeaderInfo(req)
	response, err := newcoreclient().UpdateProfileTemplate(ctx(req), &proto.UpdateProfileTemplateRequest{
		Pid:     pid,
		Name:    profiletemplatereq.Name,
		Enabled: profiletemplatereq.Enabled,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}
