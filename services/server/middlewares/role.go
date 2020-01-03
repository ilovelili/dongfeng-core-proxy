package middlewares

import (
	"context"

	restful "github.com/emicklei/go-restful"
	api "github.com/ilovelili/dongfeng-core-proxy/services/proto"
	"github.com/ilovelili/dongfeng-core-proxy/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
	sharedlib "github.com/ilovelili/dongfeng-shared-lib"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
)

func newcoreclient() api.ApiService {
	config := utils.GetConfig()
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

// RoleAuthenticate role authenticate
func RoleAuthenticate(req *restful.Request, rsp *restful.Response, chain *restful.FilterChain) {
	_, pid, _ := utils.ResolveHeaderInfo(req)	
	paths, err := newcoreclient().GetAccessiblePaths(ctx(req), &proto.GetAccessiblePathsRequest{
		Pid: pid,
	})

	if err != nil {
		writeError(rsp, errorcode.Pipe, err.Error())
		return
	}

	for _, path := range paths.GetPaths() {
		if req.Request.URL.Path == path || path == "*" {
			chain.ProcessFilter(req, rsp)
			return
		}
	}

	writeError(rsp, errorcode.GenericInsufficientPrivileges)
}
