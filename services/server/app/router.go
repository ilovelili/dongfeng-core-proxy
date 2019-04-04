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

	webservice.Route(
		webservice.
			GET("/").
			To(handlers.HealthCheck))

	// todo: add admin role middleware
	webservice.Route(
		webservice.
			GET("/dashboard").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.Dashboard))

	webservice.Route(
		webservice.
			POST("/login").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.Login))

	webservice.Route(
		webservice.
			POST("/logout").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.Logout))

	webservice.Route(
		webservice.
			POST("/user/upload").
			Consumes("multipart/form-data").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.UploadAvatar))

	webservice.Route(
		webservice.
			PUT("/user/update").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.UpdateUser))

	webservice.Route(
		webservice.
			POST("/notifications").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.UpdateNotifications))

	webservice.Route(
		webservice.
			GET("/classes").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.GetClasses))

	webservice.Route(
		webservice.
			POST("/classes").
			Consumes("multipart/form-data").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.UpdateClasses))

	webservice.Route(
		webservice.
			GET("/pupils").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.GetPupils))

	webservice.Route(
		webservice.
			POST("/pupils").
			Consumes("multipart/form-data").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.UpdatePupils))

	webservice.Route(
		webservice.
			GET("/teachers").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.GetTeachers))

	webservice.Route(
		webservice.
			POST("/teachers").
			Consumes("multipart/form-data").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.UpdateTeachers))

	webservice.Route(
		webservice.
			POST("/physique").
			Consumes("multipart/form-data").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.UploadPhysique))

	webservice.Route(
		webservice.
			GET("/ingredient/{ingredient}").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.GetIngredient).
			Param(webservice.PathParameter("ingredient", "ingredient name").DataType("string")))

	webservice.Route(
		webservice.
			POST("/ingredient").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.UpdateIngredient))

	webservice.Route(
		webservice.
			GET("/recipe/{recipe}").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.GetRecipe).
			Param(webservice.PathParameter("recipe", "recipe name").DataType("string")))

	webservice.Route(
		webservice.
			POST("/recipe").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.UpdateRecipe))

	webservice.Route(
		webservice.
			GET("/menu/{from}/{to}").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.GetMenu).
			Param(webservice.PathParameter("from", "date from").DataType("string")).
			Param(webservice.PathParameter("to", "date to").DataType("string")))

	webservice.Route(
		webservice.
			GET("/procurement/{from}/{to}/{id}").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.GetProcurement).
			Param(webservice.PathParameter("from", "from").DataType("string")).
			Param(webservice.PathParameter("to", "to").DataType("string")).
			Param(webservice.PathParameter("id", "meal id").DataType("int")))

	container.Add(webservice)

	// Add container filter to respond to OPTIONS
	container.Filter(container.OPTIONSFilter)
	return container
}
