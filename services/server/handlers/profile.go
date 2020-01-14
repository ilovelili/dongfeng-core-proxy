package handlers

import (
	"encoding/json"
	"io/ioutil"

	restful "github.com/emicklei/go-restful"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// ProfileRequestItem profile request
type ProfileRequestItem struct {
	Year     string `json:"year"`
	Class    string `json:"class"`
	Name     string `json:"name"`
	Date     string `json:"date"`
	Template string `json:"template"`
	Enabled  bool   `json:"enabled,string"`
}

// GetProfile get profile
func GetProfile(req *restful.Request, rsp *restful.Response) {
	class, year, name, date := req.QueryParameter("class"), req.QueryParameter("year"), req.QueryParameter("name"), req.QueryParameter("date")
	if class == "" || year == "" || name == "" || date == "" {
		writeError(rsp, errorcode.CoreProxyInvalidProfileGetRequest)
		return
	}

	_, pid, _ := utils.ResolveHeaderInfo(req)
	response, err := newcoreclient().GetProfile(ctx(req), &proto.GetProfileRequest{
		Pid:   pid,
		Year:  year,
		Class: class,
		Name:  name,
		Date:  date,
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

// GetPrevProfile get prev profile
func GetPrevProfile(req *restful.Request, rsp *restful.Response) {
	class, year, name, date := req.QueryParameter("class"), req.QueryParameter("year"), req.QueryParameter("name"), req.QueryParameter("date")
	if class == "" || year == "" || name == "" || date == "" {
		writeError(rsp, errorcode.CoreProxyInvalidProfileGetRequest)
		return
	}

	_, pid, _ := utils.ResolveHeaderInfo(req)
	response, err := newcoreclient().GetPrevProfile(ctx(req), &proto.GetPrevOrNextProfileRequest{
		Pid:   pid,
		Year:  year,
		Class: class,
		Name:  name,
		Date:  date,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

// GetNextProfile get next profile
func GetNextProfile(req *restful.Request, rsp *restful.Response) {
	class, year, name, date := req.QueryParameter("class"), req.QueryParameter("year"), req.QueryParameter("name"), req.QueryParameter("date")
	if class == "" || year == "" || name == "" || date == "" {
		writeError(rsp, errorcode.CoreProxyInvalidProfileGetRequest)
		return
	}

	_, pid, _ := utils.ResolveHeaderInfo(req)
	response, err := newcoreclient().GetNextProfile(ctx(req), &proto.GetPrevOrNextProfileRequest{
		Pid:   pid,
		Year:  year,
		Class: class,
		Name:  name,
		Date:  date,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

// GetProfiles get profiles
func GetProfiles(req *restful.Request, rsp *restful.Response) {
	class, year, name := req.QueryParameter("class"), req.QueryParameter("year"), req.QueryParameter("name")
	_, pid, _ := utils.ResolveHeaderInfo(req)
	response, err := newcoreclient().GetProfiles(ctx(req), &proto.GetProfilesRequest{
		Pid:   pid,
		Year:  year,
		Class: class,
		Name:  name,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

// UpdateProfile update profile endpoint used by grapejs
func UpdateProfile(req *restful.Request, rsp *restful.Response) {
	class, year, name, date := req.QueryParameter("class"), req.QueryParameter("year"), req.QueryParameter("name"), req.QueryParameter("date")
	body, err := ioutil.ReadAll(req.Request.Body)
	if err != nil {
		writeError(rsp, errorcode.CoreProxyInvalidProfileUpdateRequest)
		return
	}

	_, pid, _ := utils.ResolveHeaderInfo(req)
	response, err := newcoreclient().UpdateProfile(ctx(req), &proto.UpdateProfileRequest{
		Pid:     pid,
		Year:    year,
		Class:   class,
		Name:    name,
		Date:    date,
		Profile: string(body),
		Enabled: true,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

// DeleteProfile delete profile
func DeleteProfile(req *restful.Request, rsp *restful.Response) {
	decoder := json.NewDecoder(req.Request.Body)
	var updatereq *ProfileRequestItem
	err := decoder.Decode(&updatereq)
	if err != nil {
		writeError(rsp, errorcode.CoreProxyInvalidProfileUpdateRequest)
		return
	}

	_, pid, _ := utils.ResolveHeaderInfo(req)
	response, err := newcoreclient().DeleteProfile(ctx(req), &proto.UpdateProfileRequest{
		Pid:   pid,
		Year:  updatereq.Year,
		Class: updatereq.Class,
		Name:  updatereq.Name,
		Date:  updatereq.Date,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

// CreateProfile create profile
func CreateProfile(req *restful.Request, rsp *restful.Response) {
	decoder := json.NewDecoder(req.Request.Body)
	var updatereq *ProfileRequestItem
	err := decoder.Decode(&updatereq)
	if err != nil {
		writeError(rsp, errorcode.CoreProxyInvalidProfileUpdateRequest)
		return
	}

	_, pid, _ := utils.ResolveHeaderInfo(req)
	response, err := newcoreclient().CreateProfile(ctx(req), &proto.UpdateProfileRequest{
		Pid:          pid,
		Year:         updatereq.Year,
		Class:        updatereq.Class,
		Name:         updatereq.Name,
		Date:         updatereq.Date,
		TemplateName: updatereq.Template,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}
