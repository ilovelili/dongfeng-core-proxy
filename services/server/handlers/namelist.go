package handlers

import (
	"strings"

	restful "github.com/emicklei/go-restful"
	"github.com/gocarina/gocsv"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// NamelistRequestItem namelist request
type NamelistRequestItem struct {
	Year  string `csv:"年度"`
	Class string `csv:"班级"`
	Name  string `csv:"姓名"`
}

// GetNamelist get name list
func GetNamelist(req *restful.Request, rsp *restful.Response) {
	class, year := req.QueryParameter("class"), req.QueryParameter("year")

	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newcoreclient().GetNamelist(ctx(req), &proto.GetNamelistRequest{
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

// UpdateNamelist update name list
func UpdateNamelist(req *restful.Request, rsp *restful.Response) {
	file, _, err := req.Request.FormFile("file")
	if err != nil {
		writeError(rsp, errorcode.CoreProxyInvalidClassNameFile)
		return
	}
	defer file.Close()

	names := []*NamelistRequestItem{}
	if err := gocsv.Unmarshal(file, &names); err != nil {
		writeError(rsp, errorcode.CoreProxyInvalidClassNameFile)
		return
	}

	namelistmap := make(map[string] /*year_class*/ []*proto.NameItem /*namelist*/)
	for _, name := range names {
		key := name.Year + "_" + name.Class
		if v, ok := namelistmap[key]; ok {
			namelistmap[key] = append(v, &proto.NameItem{
				Name: name.Name,
			})
		} else {
			namelistmap[key] = []*proto.NameItem{&proto.NameItem{
				Name: name.Name,
			}}
		}
	}

	namelistitems := make([]*proto.NamelistItem, 0)
	for k, v := range namelistmap {
		segment := strings.Split(k, "_")
		if len(segment) != 2 {
			writeError(rsp, errorcode.CoreProxyInvalidClassNameFile)
			return
		}

		year, class := segment[0], segment[1]
		namelistitems = append(namelistitems, &proto.NamelistItem{
			Year:  year,
			Class: class,
			Names: v,
		})
	}

	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newcoreclient().UpdateNamelist(ctx(req), &proto.UpdateNamelistRequest{
		Token: idtoken,
		Items: namelistitems,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}
