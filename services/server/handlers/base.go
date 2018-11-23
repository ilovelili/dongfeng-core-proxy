package handlers

import (
	"context"
	"sync"

	restful "github.com/emicklei/go-restful"
	"github.com/go-redis/redis"
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
	instance    *Client
	once        sync.Once
	config      = utils.GetConfig()
	redisclient = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Host,
		Password: config.Redis.Password,
	})
)

// sessionkey string used as Redis session store key
const sessionkeyprefix = "session"

// Client struct represents InvastBroker API client
type Client struct {
	ServiceName string
	client      api.ApiService
}

// creates a new rpc client
func new() *Client {
	once.Do(func() {
		config := utils.GetConfig()
		servicename := config.ServiceNames.CoreServer
		instance = &Client{
			ServiceName: servicename,
			client:      api.NewApiService(servicename, client.DefaultClient), // rpc client
		}
	})

	return instance
}

func newclient() api.ApiService {
	return new().client
}

func ctx(req *restful.Request) context.Context {
	ip := sharedlib.ResolveRemoteIP(req.Request)
	jwks := config.Auth.JWKS
	ua := req.HeaderParameter("user-agent")

	// Set arbitrary headers in context
	return metadata.NewContext(req.Request.Context(), map[string]string{
		sharedlib.MetaDataIP:   ip,
		sharedlib.MetaDataJwks: jwks,
		sharedlib.UserAgent:    ua,
	})
}

func writeError(rsp *restful.Response, errorcode *errorcode.Error, detail ...string) {
	e := utils.NewError(errorcode, detail...)
	rsp.WriteError(int(errorcode.Code), e)
}
