package handlers

import (
	restful "github.com/emicklei/go-restful"
	"github.com/gocarina/gocsv"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// TeacherRequestItem teacher request
type TeacherRequestItem struct {
	Year  string `csv:"学年" json:"-"`
	Name  string `csv:"姓名" json:"name"`
	Class string `csv:"指导班级" json:"class"`
	Email string `csv:"邮箱" json:"email"`
	Role  string `csv:"权限" json:"role"`
}

// GetTeachers get teachers
func GetTeachers(req *restful.Request, rsp *restful.Response) {
	class, year := req.QueryParameter("class"), req.QueryParameter("year")

	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newcoreclient().GetTeachers(ctx(req), &proto.GetTeacherRequest{
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

// UpdateTeacher update teacher
func UpdateTeacher(req *restful.Request, rsp *restful.Response) {
	// tbd
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
			Role:  resolveRole(teacher.Role),
		})
	}

	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newcoreclient().UpdateTeachers(ctx(req), &proto.UpdateTeacherRequest{
		Token:    idtoken,
		Teachers: _teachers,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

func resolveRole(rawrole string) string {
	if rawrole == "管理员" {
		return "admin"
	}
	return "teacher"
}
