package handlers

import (
	"strings"

	restful "github.com/emicklei/go-restful"
	"github.com/gocarina/gocsv"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// AttendanceRequestItem attendance request
type AttendanceRequestItem struct {
	Year       string `csv:"学年"`
	Class      string `csv:"班级"`
	Date       string `csv:"日期"`
	Name       string `csv:"姓名"`
	Attendance string `csv:"出勤"`
}

// GetAttendances get pupils
func GetAttendances(req *restful.Request, rsp *restful.Response) {
	class, from, to, name := req.QueryParameter("class"), req.QueryParameter("from"), req.QueryParameter("to"), req.QueryParameter("name")

	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newcoreclient().GetAttendances(ctx(req), &proto.GetAttendanceRequest{
		Token: idtoken,
		From:  from,
		To:    to,
		Class: class,
		Name:  name,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

// UpdateAttendances update attendances
func UpdateAttendances(req *restful.Request, rsp *restful.Response) {
	file, _, err := req.Request.FormFile("file")
	if err != nil {
		writeError(rsp, errorcode.CoreProxyInvalidAttendanceUploadFile)
		return
	}
	defer file.Close()

	attendances := []*AttendanceRequestItem{}
	if err := gocsv.Unmarshal(file, &attendances); err != nil {
		writeError(rsp, errorcode.CoreProxyInvalidAttendanceUploadFile)
		return
	}

	_attendances := []*proto.Attendance{}
	for _, attendance := range attendances {
		if resolveAttendance(attendance.Attendance) {
			_attendances = append(_attendances, &proto.Attendance{
				Year:  attendance.Year,
				Date:  attendance.Date,
				Class: attendance.Class,
				Name:  attendance.Name,
			})
		}
	}

	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newcoreclient().UpdateAttendances(ctx(req), &proto.UpdateAttendanceRequest{
		Token:       idtoken,
		Attendances: _attendances,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

func resolveAttendance(attendance string) bool {
	return strings.ToLower(attendance) != "x"
}
