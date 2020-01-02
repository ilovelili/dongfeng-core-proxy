package handlers

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"regexp"

	restful "github.com/emicklei/go-restful"
	"github.com/gocarina/gocsv"
	api "github.com/ilovelili/dongfeng-core-proxy/services/proto"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	sharedlib "github.com/ilovelili/dongfeng-shared-lib"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
)

// Response api response
type Response struct {
	Message string
}

// NewResponse constructor
func NewResponse(message string) *Response {
	return &Response{Message: message}
}

var (
	config = utils.GetConfig()
)

func newcoreclient() api.ApiService {
	cli := client.NewClient(client.RequestTimeout(config.ServiceMeta.GetDefaultRequestTimeout()))
	return api.NewApiService(config.ServiceNames.CoreServer, cli) // rpc client
}

func ctx(req *restful.Request) context.Context {
	ip := sharedlib.ResolveRemoteIP(req.Request)
	ua := req.HeaderParameter("user-agent")

	// Set arbitrary headers in context
	return metadata.NewContext(req.Request.Context(), map[string]string{
		sharedlib.MetaDataIP: ip,
		sharedlib.UserAgent:  ua,
	})
}

func writeError(rsp *restful.Response, errorcode *errorcode.Error, detail ...string) {
	e := utils.NewError(errorcode, detail...)
	rsp.WriteError(int(errorcode.Code), e)
}

func init() {
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		// If GBK encoding needed
		// r := csv.NewReader(transform.NewReader(in, simplifiedchinese.GBK.NewEncoder()))
		r := csv.NewReader(in)
		r.LazyQuotes = true
		return r // Allows use dot as delimiter and use quotes in CSV
	})
}

func resolveDate(date string) (match string, ok bool) {
	re := regexp.MustCompile(`(\d{4})[-,/](\d{1,2})[-,/](\d{1,2})`)
	if ok = re.MatchString(date); !ok {
		return
	}

	matches := re.FindStringSubmatch(date)
	if len(matches) != 4 {
		return date, false
	}

	year, month, day := matches[1], fmt.Sprintf("%02s", matches[2]), fmt.Sprintf("%02s", matches[3])
	return fmt.Sprintf("%s-%s-%s", year, month, day), true
}
