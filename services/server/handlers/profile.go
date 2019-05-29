package handlers

import (
	"encoding/json"
	"io/ioutil"

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

// GetProfileNames get profile names
func GetProfileNames(req *restful.Request, rsp *restful.Response) {	year, class, name, date := req.QueryParameter("year"), req.QueryParameter("class"), req.QueryParameter("q"), req.QueryParameter("date")
	response, err := newcoreclient().GetProfileNames(ctx(req), &proto.GetProfileNamesRequest{
		Year:  year,
		Class: class,
		Name:  name,
		Date:  date,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	if len(response.Names) == 0 {
		rsp.WriteAsJson([]string{})
	} else {
		rsp.WriteAsJson(response.Names)
	}
}

// GetProfileDates get profile dates
func GetProfileDates(req *restful.Request, rsp *restful.Response) {
	year, class, name, date := req.QueryParameter("year"), req.QueryParameter("class"), req.QueryParameter("name"), req.QueryParameter("q")
	response, err := newcoreclient().GetProfileDates(ctx(req), &proto.GetProfileDatesRequest{
		Year:  year,
		Class: class,
		Name:  name,
		Date:  date,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	if len(response.Dates) == 0 {
		rsp.WriteAsJson("[]")
	} else {
		rsp.WriteAsJson(response.Dates)
	}
}

// UpdateProfile update profile
func UpdateProfile(req *restful.Request, rsp *restful.Response) {
	class, year, name, date := req.QueryParameter("class"), req.QueryParameter("year"), req.QueryParameter("name"), req.QueryParameter("date")
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
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}
