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
	class, year, name := req.QueryParameter("class"), req.QueryParameter("year"), req.QueryParameter("name")
	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newcoreclient().GetProfile(ctx(req), &proto.GetProfileRequest{
		Token: idtoken,
		Year:  year,
		Class: class,
		Name:  name,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	var result interface{}
	err = json.Unmarshal([]byte(response.GetProfile()), &result)
	if err != nil {
		writeError(rsp, errorcode.CoreProxyFailedToUnmarshalGrowthProfileData)
	}

	rsp.WriteAsJson(result)
}

// UpdateProfile update profile
func UpdateProfile(req *restful.Request, rsp *restful.Response) {
	class, year, name := req.QueryParameter("class"), req.QueryParameter("year"), req.QueryParameter("name")
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
		Profile: string(body),
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}
