package handlers

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	restful "github.com/emicklei/go-restful"
	"github.com/gocarina/gocsv"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// AttendanceRequestItem attendance request
type AttendanceRequestItem struct {
	Year       string `csv:"学年",json:"year"`
	Class      string `csv:"班级",json:"class"`
	Date       string `csv:"日期",json:"date"`
	Name       string `csv:"姓名",json:"name"`
	Attendance bool   `csv:"-",json:"attendance"`
}

// GetAttendances get pupils
func GetAttendances(req *restful.Request, rsp *restful.Response) {
	year, class, from, to, name := req.QueryParameter("year"), req.QueryParameter("class"), req.QueryParameter("from"), req.QueryParameter("to"), req.QueryParameter("name")
	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newcoreclient().GetAttendances(ctx(req), &proto.GetAttendanceRequest{
		Token: idtoken,
		Year:  year,
		From:  from,
		To:    to,
		Class: class,
		Name:  name,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

// UpdateAttendance update single attendance
func UpdateAttendance(req *restful.Request, rsp *restful.Response) {
	decoder := json.NewDecoder(req.Request.Body)
	var updatereq *AttendanceRequestItem
	err := decoder.Decode(&updatereq)
	if err != nil {
		writeError(rsp, errorcode.CoreProxyInvalidAttendanceUpdateRequestBody)
		return
	}

	attendance := &proto.Attendance{
		Year:        updatereq.Year,
		Date:        updatereq.Date,
		Class:       updatereq.Class,
		Attendances: []string{},
		Absences:    []string{},
	}

	if updatereq.Attendance {
		attendance.Attendances = append(attendance.Attendances, updatereq.Name)
	} else {
		attendance.Absences = append(attendance.Absences, updatereq.Name)
	}

	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newcoreclient().UpdateAttendance(ctx(req), &proto.UpdateAttendanceRequest{
		Token:       idtoken,
		Attendances: []*proto.Attendance{attendance},
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

// UpdateAttendances update attendances
func UpdateAttendances(req *restful.Request, rsp *restful.Response) {
	file, _, err := req.Request.FormFile("file")
	if err != nil {
		writeError(rsp, errorcode.CoreProxyInvalidAttendanceUploadFile)
		return
	}
	defer file.Close()

	absences := []*AttendanceRequestItem{}
	if err := gocsv.Unmarshal(file, &absences); err != nil {
		writeError(rsp, errorcode.CoreProxyInvalidAttendanceUploadFile)
		return
	}

	absencemap := make(map[string] /*year_class_date*/ []string)
	for _, absence := range absences {
		key := fmt.Sprintf("%s_%s_%s", absence.Year, absence.Class, absence.Date)
		if v, ok := absencemap[key]; ok {
			absencemap[key] = append(v, absence.Name)
		} else {
			absencemap[key] = []string{absence.Name}
		}
	}

	_attendances := []*proto.Attendance{}
	for k, v := range absencemap {
		segments := strings.Split(k, "_")
		if len(segments) != 3 {
			writeError(rsp, errorcode.CoreProxyInvalidAttendanceUploadFile)
		}

		year, class, rawdate := segments[0], segments[1], segments[2]
		date, ok := resolveDate(rawdate)
		if !ok {
			// date format wrong
			writeError(rsp, errorcode.CoreProxyInvalidAttendanceUploadFile)
		}

		_attendances = append(_attendances, &proto.Attendance{
			Year:     year,
			Date:     date,
			Class:    class,
			Absences: v,
		})
	}

	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newcoreclient().UpdateAttendances(ctx(req), &proto.UpdateAttendanceRequest{
		Token:       idtoken,
		Attendances: _attendances,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

func resolveDate(date string) (match string, ok bool) {
	re := regexp.MustCompile(`\d{4}[-,/]\d{2}[-,/]\d{2}`)
	if ok = re.MatchString(date); !ok {
		return
	}
	match = strings.Replace(re.FindString(date), "/", "-", -1)
	return
}
