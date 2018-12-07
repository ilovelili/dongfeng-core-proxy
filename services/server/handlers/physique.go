package handlers

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	excelize "github.com/360EntSecGroup-Skylar/excelize"
	restful "github.com/emicklei/go-restful"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	protobuf "github.com/ilovelili/dongfeng-protobuf"
)

// physiqueItem for excel parsing
type physiqueItem struct {
	name      string
	gender    string
	birthdate string
	examdate  string
	height    float32
	weight    float32
}

// UploadPhysique upload physique list
func UploadPhysique(req *restful.Request, rsp *restful.Response) {
	if err := req.Request.ParseMultipartForm(32 << 20); err != nil {
		writeError(rsp, errorcode.CoreProxyFailedToReadPhysiqueFile)
		return
	}

	file, _, err := req.Request.FormFile("physique")
	defer file.Close()
	if err != nil {
		writeError(rsp, errorcode.CoreProxyFailedToReadPhysiqueFile)
		return
	}

	excel, err := excelize.OpenReader(file)
	if err != nil {
		writeError(rsp, errorcode.CoreProxyUnsupportedMimeType)
		return
	}

	var class string
	physiques := make([]*protobuf.Physique, 0)

	for _, sheet := range excel.WorkBook.Sheets.Sheet {
		physiqueitems := make([]*protobuf.PhysiqueItem, 0)
		rows := excel.GetRows(sheet.Name)

		for rindex, row := range rows {
			if rindex == 1 {
				_class, err := parseClass(row)
				if err != nil {
					writeError(rsp, errorcode.CoreProxyBadFormatPhysiqueFile)
					return
				}

				class = _class

			} else if rindex > 3 {
				// starts with number
				if _, err := strconv.Atoi(row[0]); err == nil {
					physiqueitem, err := parsePhysiqueItem(row)
					if err != nil {
						writeError(rsp, errorcode.CoreProxyBadFormatPhysiqueFile)
						return
					}

					if physiqueitem != nil {
						gender, err := resolveGender(physiqueitem.gender)
						if err != nil {
							writeError(rsp, errorcode.CoreProxyBadFormatPhysiqueFile)
							return
						}

						physiqueitems = append(physiqueitems, &protobuf.PhysiqueItem{
							Name:      physiqueitem.name,
							Gender:    gender,
							BirthDate: physiqueitem.birthdate,
							ExamDate:  physiqueitem.examdate,
							Height:    physiqueitem.height,
							Weight:    physiqueitem.weight,
						})
					}
				}
			}
		}

		physiques = append(physiques, &protobuf.Physique{
			Class: class,
			Items: physiqueitems,
		})
	}

	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newclient().UpdatePhysique(ctx(req), &protobuf.UpdatePhysiqueRequest{
		Token:     idtoken,
		Physiques: physiques,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

func parseClass(row []string) (class string, err error) {
	// filter out empty
	var sb strings.Builder
	for _, col := range row {
		if col != "" {
			sb.WriteString(strings.Replace(col, " ", "", -1))
		}
	}
	rowstr := sb.String()

	r := regexp.MustCompile(`班级:(.+)`)
	matches := r.FindStringSubmatch(rowstr)
	if len(matches) == 2 {
		class = matches[1]
	} else {
		err = fmt.Errorf("invalid row")
	}

	return
}

func parsePhysiqueItem(row []string) (item *physiqueItem, err error) {
	if row[0] == "" || row[0] == "0" {
		return
	}
	if row[1] == "" || row[1] == "0" {
		return
	}

	item = new(physiqueItem)
	item.name = row[1]
	item.gender = row[2]
	item.birthdate = row[3]
	item.examdate = row[4]

	_height, err := strconv.ParseFloat(row[5], 32)
	if err != nil {
		return
	}
	item.height = float32(_height)

	_weight, err := strconv.ParseFloat(row[6], 32)
	if err != nil {
		return
	}
	item.weight = float32(_weight)
	return
}

func resolveGender(gender string) (g protobuf.PhysiqueItem_Gender, err error) {
	if gender == "女" {
		g = protobuf.PhysiqueItem_F
		return
	}

	if gender == "男" {
		g = protobuf.PhysiqueItem_M
		return
	}

	err = fmt.Errorf("failed to parse gender")
	return
}
