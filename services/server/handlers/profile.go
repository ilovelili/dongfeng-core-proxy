package handlers

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	restful "github.com/emicklei/go-restful"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// GetProfile get profile
func GetProfile(req *restful.Request, rsp *restful.Response) {
	class, year, name, date := req.QueryParameter("class"), req.QueryParameter("year"), req.QueryParameter("name"), req.QueryParameter("date")
	if class == "" || year == "" || name == "" || date == "" {
		writeError(rsp, errorcode.CoreProxyInvalidProfileUpdateRequest)
		return
	}

	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newcoreclient().GetProfile(ctx(req), &proto.GetProfileRequest{
		Token: idtoken,
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

// GetProfiles get profiles
func GetProfiles(req *restful.Request, rsp *restful.Response) {
	class, year, name := req.QueryParameter("class"), req.QueryParameter("year"), req.QueryParameter("name")
	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newcoreclient().GetProfiles(ctx(req), &proto.GetProfilesRequest{
		Token: idtoken,
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

// UpdateProfile update profile
func UpdateProfile(req *restful.Request, rsp *restful.Response) {
	class, year, name, date, enabled := req.QueryParameter("class"), req.QueryParameter("year"), req.QueryParameter("name"), req.QueryParameter("date"), req.QueryParameter("enabled")
	body, err := ioutil.ReadAll(req.Request.Body)
	if err != nil {
		writeError(rsp, errorcode.CoreProxyInvalidPupilUpdateRequestBody)
		return
	}

	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newcoreclient().UpdateProfile(ctx(req), &proto.UpdateProfileRequest{
		Token:   idtoken,
		Year:    year,
		Class:   class,
		Name:    name,
		Date:    date,
		Profile: string(body),
		Enabled: resolveEnabled(enabled),
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

func resolveEnabled(enabled string) bool {
	str := strings.ToLower(enabled)
	return str == "true" || str == "1"
}
