package handlers

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	excelize "github.com/360EntSecGroup-Skylar/excelize"
	restful "github.com/emicklei/go-restful"
	errorcode "github.com/ilovelili/dongfeng-error-code"
)

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

	// attendancereq := &protobuf.AttendanceRequest{}
	for _, sheet := range excel.WorkBook.Sheets.Sheet {
		rows := excel.GetRows(sheet.Name)
		for rindex, row := range rows {
			if rindex == 1 {
				year, month, class, err := parseYearMonthClass(row)
				if err != nil {
					writeError(rsp, errorcode.CoreProxyBadFormatAttendanceFile)
				}

				fmt.Println(year)
				fmt.Println(month)
				fmt.Println(class)

				// attendance := &protobuf.ClassAttendance{
				// 	Year:  year,
				// 	Month: month,
				// 	Class: class,
				// }

			}
		}
	}
}

func parseYearMonthClass(row []string) (year, month int32, class string, err error) {
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
		year = int32(_year)

		_month, parserr := strconv.ParseInt(matches[2], 10, 32)
		if err != nil {
			err = parserr
			return
		}
		month = int32(_month)
	} else {
		err = fmt.Errorf("invalid row")
	}

	return
}
