package handlers

import (
	"encoding/json"

	restful "github.com/emicklei/go-restful"
	"github.com/gocarina/gocsv"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// PhysiqueReqItem physique request
type PhysiqueReqItem struct {
	ID        int64   `csv:"-" json:"id"`
	Name      string  `csv:"姓名" json:"name"`
	Class     string  `csv:"班级" json:"class"`
	Year      string  `csv:"学年" json:"year"`
	Gender    string  `csv:"性别" json:"gender"`
	BirthDate string  `csv:"出生日期" json:"birth_date"`
	ExamDate  string  `csv:"体检日期" json:"exam_date"`
	Height    float64 `csv:"身高" json:"height,string"`
	Weight    float64 `csv:"体重" json:"weight,string"`
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

// UpdatePhysique update physique
func UpdatePhysique(req *restful.Request, rsp *restful.Response) {
	decoder := json.NewDecoder(req.Request.Body)
	var updatereq *PhysiqueReqItem
	err := decoder.Decode(&updatereq)
	if err != nil {
		writeError(rsp, errorcode.CoreProxyInvalidPhysiqueUpdateRequestBody)
		return
	}

	if gendervalid := validateGender(updatereq.Gender); !gendervalid {
		writeError(rsp, errorcode.CoreProxyInvalidPhysiqueUpdateRequestBody)
		return
	}

	birthDate, ok := resolveDate(updatereq.BirthDate)
	if !ok {
		// date format wrong
		writeError(rsp, errorcode.CoreProxyInvalidPhysiqueUploadFile)
	}

	examDate, ok := resolveDate(updatereq.ExamDate)
	if !ok {
		writeError(rsp, errorcode.CoreProxyInvalidPhysiqueUploadFile)
	}

	physique := &proto.Physique{
		Id:        updatereq.ID,
		Year:      updatereq.Year,
		Class:     updatereq.Class,
		Name:      updatereq.Name,
		Gender:    updatereq.Gender,
		BirthDate: birthDate,
		ExamDate:  examDate,
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

// UpdatePhysiques update physiques
func UpdatePhysiques(req *restful.Request, rsp *restful.Response) {
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
		if gendervalid := validateGender(physique.Gender); !gendervalid {
			writeError(rsp, errorcode.CoreProxyInvalidPhysiqueUploadFile)
			return
		}

		birthDate, ok := resolveDate(physique.BirthDate)
		if !ok {
			// date format wrong
			writeError(rsp, errorcode.CoreProxyInvalidPhysiqueUploadFile)
		}

		examDate, ok := resolveDate(physique.ExamDate)
		if !ok {
			writeError(rsp, errorcode.CoreProxyInvalidPhysiqueUploadFile)
		}

		_physiques = append(_physiques, &proto.Physique{
			Name:      physique.Name,
			Year:      physique.Year,
			Class:     physique.Class,
			Gender:    physique.Gender,
			BirthDate: birthDate,
			ExamDate:  examDate,
			Height:    physique.Height,
			Weight:    physique.Weight,
		})
	}

	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newcoreclient().UpdatePhysiques(ctx(req), &proto.UpdatePhysiqueRequest{
		Token:     idtoken,
		Physiques: _physiques,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

// GetMasters get masters
func GetMasters(req *restful.Request, rsp *restful.Response) {
	id := req.QueryParameter("id")

	idtoken, _ := utils.ResolveIDToken(req)
	var response interface{}
	var err error

	switch id {
	case "1":
		response, err = newcoreclient().GetAgeHeightWeightPMasters(ctx(req), &proto.GetAgeHeightWeightPMasterRequest{Token: idtoken})
		break
	case "2":
		response, err = newcoreclient().GetAgeHeightWeightSDMasters(ctx(req), &proto.GetAgeHeightWeightSDMasterRequest{Token: idtoken})
		break
	case "3":
		response, err = newcoreclient().GetBMIMasters(ctx(req), &proto.GetBMIMasterRequest{Token: idtoken})
		break
	case "4":
		response, err = newcoreclient().GetHeightToWeightPMasters(ctx(req), &proto.GetHeightToWeightPMasterRequest{Token: idtoken})
		break
	case "5":
		response, err = newcoreclient().GetHeightToWeightSDMasters(ctx(req), &proto.GetHeightToWeightSDMasterRequest{Token: idtoken})
		break
	default:
		writeError(rsp, errorcode.CoreProxyInvalidPhysiqueMasterRequest)
		return
	}

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

func validateGender(gender string) bool {
	return gender == "女" || gender == "男"
}
