package handlers

import (
	"encoding/json"

	restful "github.com/emicklei/go-restful"
	"github.com/gocarina/gocsv"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// TeacherRequestItem teacher request
type TeacherRequestItem struct {
	ID    int64  `csv:"-" json:"id"`
	Year  string `csv:"学年" json:"-"`
	Name  string `csv:"姓名" json:"name"`
	Class string `csv:"指导班级" json:"class"`
	Email string `csv:"邮箱" json:"email"`
	Role  string `csv:"权限" json:"-"`
}

// GetTeachers get teachers
func GetTeachers(req *restful.Request, rsp *restful.Response) {
	class, year := req.QueryParameter("class"), req.QueryParameter("year")

	_,  pid, _ := utils.ResolveHeaderInfo(req)
	response, err := newcoreclient().GetTeachers(ctx(req), &proto.GetTeacherRequest{
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

// UpdateTeacher update teacher
func UpdateTeacher(req *restful.Request, rsp *restful.Response) {
	decoder := json.NewDecoder(req.Request.Body)
	var updatereq *TeacherRequestItem
	err := decoder.Decode(&updatereq)
	if err != nil {
		writeError(rsp, errorcode.CoreProxyInvalidTeacherUpdateRequestBody)
		return
	}

	teacher := &proto.Teacher{
		Id:    updatereq.ID,
		Name:  updatereq.Name,
		Class: updatereq.Class,
		Email: updatereq.Email,
	}

	_,  pid, _ := utils.ResolveHeaderInfo(req)
	response, err := newcoreclient().UpdateTeacher(ctx(req), &proto.UpdateTeacherRequest{
		Pid: pid,
		Teachers: []*proto.Teacher{teacher},
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

// UpdateTeachers update teachers
func UpdateTeachers(req *restful.Request, rsp *restful.Response) {
	file, _, err := req.Request.FormFile("file")
	if err != nil {
		writeError(rsp, errorcode.CoreProxyInvalidTeacherUploadFile)
		return
	}
	defer file.Close()

	teachers := []*TeacherRequestItem{}
	if err := gocsv.Unmarshal(file, &teachers); err != nil {
		writeError(rsp, errorcode.CoreProxyInvalidTeacherUploadFile)
		return
	}

	_teachers := []*proto.Teacher{}
	for _, teacher := range teachers {
		_teachers = append(_teachers, &proto.Teacher{
			Year:  teacher.Year,
			Name:  teacher.Name,
			Class: teacher.Class,
			Email: teacher.Email,
			Role:  teacher.Role,
		})
	}

	_,  pid, _ := utils.ResolveHeaderInfo(req)
	response, err := newcoreclient().UpdateTeachers(ctx(req), &proto.UpdateTeacherRequest{
		Pid: pid,
		Teachers: _teachers,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}
