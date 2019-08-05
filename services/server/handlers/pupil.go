package handlers

import (
	"encoding/json"

	restful "github.com/emicklei/go-restful"
	"github.com/gocarina/gocsv"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// PupilRequestItem pupil request
type PupilRequestItem struct {
	ID    int64  `csv:"-" json:"id"`
	Year  string `csv:"学年" json:"year"`
	Class string `csv:"班级" json:"class"`
	Name  string `csv:"姓名" json:"name"`
}

// GetPupils get pupils
func GetPupils(req *restful.Request, rsp *restful.Response) {
	class, year := req.QueryParameter("class"), req.QueryParameter("year")

	_,  pid, _ := utils.ResolveHeaderInfo(req)
	response, err := newcoreclient().GetPupils(ctx(req), &proto.GetPupilRequest{
		Pid: pid,
		Year:  year,
		Class: class,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

// UpdatePupil update pupil
func UpdatePupil(req *restful.Request, rsp *restful.Response) {
	decoder := json.NewDecoder(req.Request.Body)
	var updatereq *PupilRequestItem
	err := decoder.Decode(&updatereq)
	if err != nil {
		writeError(rsp, errorcode.CoreProxyInvalidPupilUpdateRequestBody)
		return
	}

	pupil := &proto.Pupil{
		Id:    updatereq.ID,
		Name:  updatereq.Name,
		Class: updatereq.Class,
		Year:  updatereq.Year,
	}

	_,  pid, _ := utils.ResolveHeaderInfo(req)
	response, err := newcoreclient().UpdatePupil(ctx(req), &proto.UpdatePupilRequest{
		Pid: pid,
		Pupils: []*proto.Pupil{pupil},
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

	_,  pid, _ := utils.ResolveHeaderInfo(req)
	response, err := newcoreclient().UpdatePupils(ctx(req), &proto.UpdatePupilRequest{
		Pid: pid,
		Pupils: _pupils,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}
