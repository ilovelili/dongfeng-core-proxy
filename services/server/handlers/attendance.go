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

// GetAttendance get attenance list
func GetAttendance(req *restful.Request, rsp *restful.Response) {
	class, from, to, name := req.QueryParameter("class"), req.QueryParameter("from"), req.QueryParameter("to"), req.QueryParameter("name")
	if from != "" && to != "" && from > to {
		writeError(rsp, errorcode.CoreProxyInvalidAttendanceRequest)
		return
	}

	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newattendanceclient().GetAttendance(ctx(req), &protobuf.GetAttendanceRequest{
		Token: idtoken,
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

// UploadAttendance upload attenance list
func UploadAttendance(req *restful.Request, rsp *restful.Response) {
	if err := req.Request.ParseMultipartForm(32 << 20); err != nil {
		writeError(rsp, errorcode.CoreProxyFailedToReadAttendanceFile)
		return
	}

	file, _, err := req.Request.FormFile("attendance")
	defer file.Close()
	if err != nil {
		writeError(rsp, errorcode.CoreProxyFailedToReadAttendanceFile)
		return
	}

	excel, err := excelize.OpenReader(file)
	if err != nil {
		writeError(rsp, errorcode.CoreProxyUnsupportedMimeType)
		return
	}

	classattendances := make([]*protobuf.ClassAttendance, 0)
	for _, sheet := range excel.WorkBook.Sheets.Sheet {
		rows := excel.GetRows(sheet.Name)

		var year, month int64
		var class string
		var dates []int64

		for rindex, row := range rows {
			if rindex == 1 {
				year, month, class, err = parseYearMonthClass(row)
				if err != nil {
					writeError(rsp, errorcode.CoreProxyBadFormatAttendanceFile)
					return
				}

			} else if rindex == 3 {
				dates, err = parseDates(row)
				if err != nil {
					writeError(rsp, errorcode.CoreProxyBadFormatAttendanceFile)
					return
				}

			} else if rindex > 3 {
				// starts with number
				if _, err := strconv.Atoi(row[0]); err == nil {
					name, attendances, err := parseNameAttendences(row, dates)
					if err != nil {
						writeError(rsp, errorcode.CoreProxyBadFormatAttendanceFile)
						return
					}

					if name != "" && year > 0 && month > 0 && class != "" {
						classattendances = append(classattendances, &protobuf.ClassAttendance{
							Year:        year,
							Month:       month,
							Name:        name,
							Class:       class,
							Attendances: attendances,
						})
					}
				}
			}
		}
	}

	idtoken, _ := utils.ResolveIDToken(req)
	response, err := newattendanceclient().UpdateAttendance(ctx(req), &protobuf.UpdateAttendanceRequest{
		Token:       idtoken,
		Attendances: classattendances,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	rsp.WriteAsJson(response)
}

func parseYearMonthClass(row []string) (year, month int64, class string, err error) {
	// filter out empty
	var sb strings.Builder
	for _, col := range row {
		if col != "" {
			sb.WriteString(strings.Replace(col, " ", "", -1))
		}
	}
	rowstr := sb.String()

	r := regexp.MustCompile(`(19|20\d\d)年(0?[1-9]|1[012])月班级(.+)`)
	matches := r.FindStringSubmatch(rowstr)
	if len(matches) == 4 {
		_year, parserr := strconv.ParseInt(matches[1], 10, 32)
		if parserr != nil {
			err = parserr
			return
		}
		year = _year

		_month, parserr := strconv.ParseInt(matches[2], 10, 32)
		if err != nil {
			err = parserr
			return
		}
		month = _month

		class = matches[3]

	} else {
		err = fmt.Errorf("invalid row")
	}

	return
}

func parseDates(row []string) (dates []int64, err error) {
	dates = make([]int64, 0)
	for _, col := range row {
		if col == "0" {
			return
		}

		date, err := strconv.ParseInt(col, 10, 32)
		if err == nil {
			dates = append(dates, date)
		}
	}

	return
}

func parseNameAttendences(row []string, dates []int64) (name string, attendances []int64, err error) {
	if len(dates) == 0 {
		err = fmt.Errorf("invalid parsed dates")
		return
	}

	attendances = make([]int64, 0)
	for i, col := range row {
		if i == 1 {
			if col != "" && col != "0" {
				name = col
			} else {
				continue
			}
		} else if i > 1 {
			if col == "O" && i < len(dates)+2 {
				attendances = append(attendances, dates[i-2] /*since i starts from 2*/)
			}
		}
	}

	return
}
