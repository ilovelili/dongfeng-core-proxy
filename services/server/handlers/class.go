package handlers

import (
	restful "github.com/emicklei/go-restful"
	"github.com/gocarina/gocsv"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// ClassRequestItem class request
type ClassRequestItem struct {
	Name string `csv:"班级"`
}

// GetClasses get class list
func GetClasses(req *restful.Request, rsp *restful.Response) {
	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newcoreclient().GetClasses(ctx(req), &proto.GetClassRequest{
		Token: idtoken,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

// UpdateClasses update classes
func UpdateClasses(req *restful.Request, rsp *restful.Response) {
	file, _, err := req.Request.FormFile("file")
	if err != nil {
		writeError(rsp, errorcode.CoreProxyInvalidClassUploadFile)
		return
	}
	defer file.Close()

	classes := []*ClassRequestItem{}
	if err := gocsv.Unmarshal(file, &classes); err != nil {
		writeError(rsp, errorcode.CoreProxyInvalidClassUploadFile)
		return
	}

	_classes := []*proto.Class{}
	for _, class := range classes {
		_classes = append(_classes, &proto.Class{
			Name: class.Name,
		})
	}

	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newcoreclient().UpdateClasses(ctx(req), &proto.UpdateClassRequest{
		Token:   idtoken,
		Classes: _classes,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}
