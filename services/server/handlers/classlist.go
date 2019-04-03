package handlers

import (
	restful "github.com/emicklei/go-restful"
	"github.com/gocarina/gocsv"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// ClasslistRequestItem classlist request
type ClasslistRequestItem struct {
	Name string `csv:"班级"`
}

// GetClasslist get class list
func GetClasslist(req *restful.Request, rsp *restful.Response) {
	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newcoreclient().GetClasslist(ctx(req), &proto.GetClasslistRequest{
		Token: idtoken,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

// UpdateClasslist update class list
func UpdateClasslist(req *restful.Request, rsp *restful.Response) {
	file, _, err := req.Request.FormFile("file")
	if err != nil {
		writeError(rsp, errorcode.CoreProxyInvalidClassListFile)
		return
	}
	defer file.Close()

	classes := []*ClasslistRequestItem{}
	if err := gocsv.Unmarshal(file, &classes); err != nil {
		writeError(rsp, errorcode.CoreProxyInvalidClassListFile)
		return
	}

	classlistitems := make([]*proto.ClassItem, 0)
	for _, class := range classes {
		classlistitems = append(classlistitems, &proto.ClassItem{
			Name: class.Name,
		})
	}

	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newcoreclient().UpdateClasslist(ctx(req), &proto.UpdateClasslistRequest{
		Token: idtoken,
		Items: classlistitems,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}
