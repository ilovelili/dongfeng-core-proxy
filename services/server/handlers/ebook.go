package handlers

import (
	"encoding/json"

	restful "github.com/emicklei/go-restful"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// EbookRequestItem ebook request
type EbookRequestItem struct {
	Year   string   `json:"year"`
	Class  string   `json:"class"`
	Name   string   `json:"name"`
	Date   string   `json:"date"`
	Images []string `json:"images"`
	HTML   string   `json:"html"`
	CSS    string   `json:"css"`
}

// GetEbooks get ebooks
func GetEbooks(req *restful.Request, rsp *restful.Response) {
	year, class, name, from, to := req.QueryParameter("year"), req.QueryParameter("class"), req.QueryParameter("name"), req.QueryParameter("from"), req.QueryParameter("to")

	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newcoreclient().GetEbooks(ctx(req), &proto.GetEbooksRequest{
		Token: idtoken,
		Year:  year,
		Class: class,
		Name:  name,
		From:  from,
		To:    to,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

// UpdateEbook update ebook
func UpdateEbook(req *restful.Request, rsp *restful.Response) {
	decoder := json.NewDecoder(req.Request.Body)
	var updatereq *EbookRequestItem
	err := decoder.Decode(&updatereq)
	if err != nil {
		writeError(rsp, errorcode.CoreProxyInvalidEbookUpdateRequest)
		return
	}

	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newcoreclient().UpdateEbook(ctx(req), &proto.UpdateEbookRequest{
		Token:  idtoken,
		Year:   updatereq.Year,
		Class:  updatereq.Class,
		Name:   updatereq.Name,
		Date:   updatereq.Date,
		Images: updatereq.Images,
		Html:   updatereq.HTML,
		Css:    updatereq.CSS,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}
