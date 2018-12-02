package app

import (
	"net/http"

	restful "github.com/emicklei/go-restful"
	"github.com/ilovelili/dongfeng-core-proxy/services/server/handlers"
	"github.com/ilovelili/dongfeng-core-proxy/services/server/middlewares"
)

// Router restful router wrapper
type Router struct {
	basepath string
}

// NewRouter init router
func NewRouter() *Router {
	return &Router{
		basepath: "/api",
	}
}

// Route route mapping
func (r *Router) Route() http.Handler {
	webservice := new(restful.WebService)
	container := restful.NewContainer()
	webservice.
		Path(r.basepath).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	webservice.Route(webservice.GET("/").To(handlers.HealthCheck))

	// todo: add admin role middleware
	webservice.Route(webservice.GET("/dashboard").Filter(middlewares.JwtAuthenticate).To(handlers.Dashboard))
	webservice.Route(webservice.POST("/login").Filter(middlewares.JwtAuthenticate).To(handlers.Login))
	webservice.Route(webservice.POST("/logout").Filter(middlewares.JwtAuthenticate).To(handlers.Logout))
	webservice.Route(webservice.POST("/user/upload").Consumes("multipart/form-data").Filter(middlewares.JwtAuthenticate).To(handlers.UploadAvatar))
	webservice.Route(webservice.PUT("/user/update").Filter(middlewares.JwtAuthenticate).To(handlers.UpdateUser))

	container.Add(webservice)

	// Add container filter to respond to OPTIONS
	container.Filter(container.OPTIONSFilter)

	return container
}
