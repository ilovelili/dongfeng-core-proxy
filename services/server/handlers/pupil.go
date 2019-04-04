package handlers

import (
	restful "github.com/emicklei/go-restful"
	"github.com/gocarina/gocsv"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// PupilRequestItem pupil request
type PupilRequestItem struct {
	Year  string `csv:"学年"`
	Class string `csv:"班级"`
	Name  string `csv:"姓名"`
}

// GetPupils get pupils
func GetPupils(req *restful.Request, rsp *restful.Response) {
	class, year := req.QueryParameter("class"), req.QueryParameter("year")

	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newcoreclient().GetPupils(ctx(req), &proto.GetPupilRequest{
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

// UpdatePupils update pupils
func UpdatePupils(req *restful.Request, rsp *restful.Response) {
	file, _, err := req.Request.FormFile("file")
	if err != nil {
		writeError(rsp, errorcode.CoreProxyInvalidPupilUploadFile)
		return
	}
	defer file.Close()

	pupils := []*PupilRequestItem{}
	if err := gocsv.Unmarshal(file, &pupils); err != nil {
		writeError(rsp, errorcode.CoreProxyInvalidPupilUploadFile)
		return
	}

	_pupils := []*proto.Pupil{}
	for _, pupil := range pupils {
		_pupils = append(_pupils, &proto.Pupil{
			Year:  pupil.Year,
			Class: pupil.Class,
			Name:  pupil.Name,
		})
	}

	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newcoreclient().UpdatePupils(ctx(req), &proto.UpdatePupilRequest{
		Token:  idtoken,
		Pupils: _pupils,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}
