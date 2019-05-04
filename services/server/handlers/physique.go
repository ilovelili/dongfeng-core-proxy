package handlers

import (
	"encoding/json"
	"fmt"

	restful "github.com/emicklei/go-restful"
	"github.com/gocarina/gocsv"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// PhysiqueReqItem physique request
type PhysiqueReqItem struct {
	ID        int64   `csv:"-" json:"id"`
	Name      string  `csv:"姓名" json:"-"`
	Class     string  `csv:"班级" json:"-"`
	Year      string  `csv:"学年" json:"-"`
	Gender    string  `csv:"性别" json:"gender"`
	BirthDate string  `csv:"出生日期" json:"birth_date"`
	ExamDate  string  `csv:"体检日期" json:"exam_date"`
	Height    float64 `csv:"身高" json:"height"`
	Weight    float64 `csv:"体重" json:"weight"`
}

// GetPhysiques get physiques
func GetPhysiques(req *restful.Request, rsp *restful.Response) {
	class, year, name := req.QueryParameter("class"), req.QueryParameter("year"), req.QueryParameter("name")

	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newcoreclient().GetPhysiques(ctx(req), &proto.GetPhysiqueRequest{
		Token: idtoken,
		Year:  year,
		Class: class,
		Name:  name,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

// UploadPhysique update physique
func UploadPhysique(req *restful.Request, rsp *restful.Response) {
	decoder := json.NewDecoder(req.Request.Body)
	var updatereq *PhysiqueReqItem
	err := decoder.Decode(&updatereq)
	if err != nil {
		writeError(rsp, errorcode.CoreProxyInvalidPhysiqueUpdateRequestBody)
		return
	}

	gender, err := resolveGender(updatereq.Gender)
	if err != nil {
		writeError(rsp, errorcode.CoreProxyInvalidPhysiqueUpdateRequestBody)
		return
	}

	physique := &proto.Physique{
		Id:        updatereq.ID,
		Gender:    gender,
		BirthDate: updatereq.BirthDate,
		ExamDate:  updatereq.ExamDate,
		Height:    updatereq.Height,
		Weight:    updatereq.Weight,
	}

	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newcoreclient().UpdatePhysique(ctx(req), &proto.UpdatePhysiqueRequest{
		Token:     idtoken,
		Physiques: []*proto.Physique{physique},
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

// UploadPhysiques upload physiques
func UploadPhysiques(req *restful.Request, rsp *restful.Response) {
	file, _, err := req.Request.FormFile("file")
	if err != nil {
		writeError(rsp, errorcode.CoreProxyFailedToReadPhysiqueFile)
		return
	}
	defer file.Close()

	physiques := []*PhysiqueReqItem{}
	if err := gocsv.Unmarshal(file, &physiques); err != nil {
		writeError(rsp, errorcode.CoreProxyInvalidPhysiqueUploadFile)
		return
	}

	_physiques := []*proto.Physique{}
	for _, physique := range physiques {
		gender, err := resolveGender(physique.Gender)
		if err != nil {
			writeError(rsp, errorcode.CoreProxyInvalidPhysiqueUploadFile)
			return
		}

		_physiques = append(_physiques, &proto.Physique{
			Name:      physique.Name,
			Year:      physique.Year,
			Class:     physique.Class,
			Gender:    gender,
			BirthDate: physique.BirthDate,
			ExamDate:  physique.ExamDate,
			Height:    physique.Height,
			Weight:    physique.Weight,
		})
	}

	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newcoreclient().UpdatePhysique(ctx(req), &proto.UpdatePhysiqueRequest{
		Token:     idtoken,
		Physiques: _physiques,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

func resolveGender(gender string) (g proto.Physique_Gender, err error) {
	if gender == "女" {
		g = proto.Physique_F
		return
	}

	if gender == "男" {
		g = proto.Physique_M
		return
	}

	err = fmt.Errorf("failed to parse gender")
	return
}
