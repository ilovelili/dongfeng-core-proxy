package handlers

import (
	"encoding/json"
	"io/ioutil"

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

// GetProfileTemplates get templates
func GetProfileTemplates(req *restful.Request, rsp *restful.Response) {
	_, pid, email, _ := utils.ResolveHeaderInfo(req)
	response, err := newcoreclient().GetProfileTemplates(ctx(req), &proto.GetProfileTemplatesRequest{
		Pid:   pid,
		Email: email,
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

	_, pid, email, _ := utils.ResolveHeaderInfo(req)
	response, err := newcoreclient().UpdateProfileTemplate(ctx(req), &proto.UpdateProfileTemplateRequest{
		Pid:     pid,
		Email:   email,
		Name:    profiletemplatereq.Name,
		Enabled: profiletemplatereq.Enabled,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

// GetProfileTemplateDetail get profile template detail by grapejs
func GetProfileTemplateDetail(req *restful.Request, rsp *restful.Response) {
	name := req.QueryParameter("name")
	_, pid, email, _ := utils.ResolveHeaderInfo(req)
	response, err := newcoreclient().GetProfileTemplate(ctx(req), &proto.GetProfileTemplateRequest{
		Pid:   pid,
		Email: email,
		Name:  name,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	// not found
	if response.GetProfile() == "" {
		rsp.WriteAsJson(response.GetProfile())
		return
	}

	var result interface{}
	err = json.Unmarshal([]byte(response.GetProfile()), &result)
	if err != nil {
		writeError(rsp, errorcode.CoreProxyFailedToUnmarshalGrowthProfileData)
	}

	rsp.WriteAsJson(result)
}

// UpdateProfileTemplateDetail update profile template detail by grapejs
func UpdateProfileTemplateDetail(req *restful.Request, rsp *restful.Response) {
	name := req.QueryParameter("name")
	body, err := ioutil.ReadAll(req.Request.Body)
	if err != nil {
		writeError(rsp, errorcode.CoreProxyInvalidProfileTemplateUpdateRequestBody)
		return
	}

	_, pid, email, _ := utils.ResolveHeaderInfo(req)
	response, err := newcoreclient().UpdateProfileTemplate(ctx(req), &proto.UpdateProfileTemplateRequest{
		Pid:     pid,
		Email:   email,
		Name:    name,
		Profile: string(body),
		Enabled: true,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}
