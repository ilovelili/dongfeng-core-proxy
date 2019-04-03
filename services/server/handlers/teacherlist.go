package handlers

import (
	restful "github.com/emicklei/go-restful"
	"github.com/gocarina/gocsv"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// TeacherlistRequestItem teacherlist request
type TeacherlistRequestItem struct {
	Year  string `csv:"学年"`
	Name  string `csv:"姓名"`
	Class string `csv:"指导班级"`
	Email string `csv:"邮箱"`
	Role  string `csv:"权限"`
}

// GetTeacherlist get teacher list
func GetTeacherlist(req *restful.Request, rsp *restful.Response) {
	class, year := req.QueryParameter("class"), req.QueryParameter("year")

	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newcoreclient().GetTeacherlist(ctx(req), &proto.GetTeacherlistRequest{
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

// UpdateTeacherlist update teacher list
func UpdateTeacherlist(req *restful.Request, rsp *restful.Response) {
	file, _, err := req.Request.FormFile("file")
	if err != nil {
		writeError(rsp, errorcode.CoreProxyInvalidTeacherListFile)
		return
	}
	defer file.Close()

	teachers := []*TeacherlistRequestItem{}
	if err := gocsv.Unmarshal(file, &teachers); err != nil {
		writeError(rsp, errorcode.CoreProxyInvalidTeacherListFile)
		return
	}

	teacherlistmap := make(map[string] /*year*/ []*proto.TeacherItem /*teacherlist*/)
	for _, teacher := range teachers {
		key := teacher.Year
		if v, ok := teacherlistmap[key]; ok {
			teacherlistmap[key] = append(v, &proto.TeacherItem{
				Name:  teacher.Name,
				Class: teacher.Class,
				Email: teacher.Email,
				Role:  resolveRole(teacher.Role),
			})
		} else {
			teacherlistmap[key] = []*proto.TeacherItem{&proto.TeacherItem{
				Name:  teacher.Name,
				Class: teacher.Class,
				Email: teacher.Email,
				Role:  resolveRole(teacher.Role),
			}}
		}
	}

	teacherlistitems := make([]*proto.TeacherlistItem, 0)
	for k, v := range teacherlistmap {
		teacherlistitems = append(teacherlistitems, &proto.TeacherlistItem{
			Year:  k,
			Items: v,
		})
	}

	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newcoreclient().UpdateTeacherlist(ctx(req), &proto.UpdateTeacherlistRequest{
		Token: idtoken,
		Items: teacherlistitems,
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
