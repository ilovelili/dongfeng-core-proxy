package handlers

import (
	"sync"

	api "github.com/ilovelili/dongfeng/core-proxy/services/proto"
	"github.com/ilovelili/dongfeng/core-proxy/services/utils"

	"github.com/micro/go-micro/client"
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
	instance *Client
	once     sync.Once
)

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
