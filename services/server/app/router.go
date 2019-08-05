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
			POST("/pupil").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.UpdatePupil))

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
			POST("/teacher").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.UpdateTeacher))

	webservice.Route(
		webservice.
			POST("/teachers").
			Consumes("multipart/form-data").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.UpdateTeachers))

	webservice.Route(
		webservice.
			GET("/attendances").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.GetAttendances))

	webservice.Route(
		webservice.
			POST("/attendance").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.UpdateAttendance))

	webservice.Route(
		webservice.
			POST("/attendances").
			Consumes("multipart/form-data").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.UpdateAttendances))

	webservice.Route(
		webservice.
			GET("/physiques").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.GetPhysiques))

	webservice.Route(
		webservice.
			POST("/physique").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.UpdatePhysique))

	webservice.Route(
		webservice.
			POST("/physiques").
			Consumes("multipart/form-data").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.UpdatePhysiques))

	webservice.Route(
		webservice.
			GET("/masters").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.GetMasters))

	webservice.Route(
		webservice.
			GET("/menus").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.GetMenus))

	webservice.Route(
		webservice.
			GET("/recipes").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.GetRecipes))

	webservice.Route(
		webservice.
			GET("/ingredients").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.GetIngredients))

	// shame on it since front end autocomplete package doesn't allow headers so can't pass JWT
	webservice.Route(
		webservice.
			GET("/ingredient/names").
			To(handlers.GetIngredientNames))

	webservice.Route(
		webservice.
			POST("/ingredient").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.UpdateIngredient))

	webservice.Route(
		webservice.
			POST("/ingredients").
			Consumes("multipart/form-data").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.UpdateIngredients))

	webservice.Route(
		webservice.
			GET("/procurements").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.GetProcurements))

	webservice.Route(
		webservice.
			POST("/procurement").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.UpdateProcurement))

	webservice.Route(
		webservice.
			GET("/profile/prev").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.GetPrevProfile))

	webservice.Route(
		webservice.
			GET("/profile/next").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.GetNextProfile))

	webservice.Route(
		webservice.
			GET("/profile").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.GetProfile))

	webservice.Route(
		webservice.
			GET("/profiles").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.GetProfiles))

	webservice.Route(
		webservice.
			POST("/profile/create").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.CreateProfile))

	webservice.Route(
		webservice.
			POST("/profile/delete").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.DeleteProfile))

	webservice.Route(
		webservice.
			POST("/profile").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.UpdateProfile))

	webservice.Route(
		webservice.
			GET("/ebooks").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.GetEbooks))

	webservice.Route(
		webservice.
			POST("/ebook").
			Filter(middlewares.JwtAuthenticate).
			To(handlers.UpdateEbook))

	container.Add(webservice)

	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"Content-Type", "Accept", "Authorization"},
		AllowedMethods: []string{"GET", "POST"},
		AllowedDomains: []string{"http://localhost:8080", "http://localhost:4040"},
		CookiesAllowed: false,
		Container:      container}
	container.Filter(cors.Filter)

	container.Filter(container.OPTIONSFilter)

	return container
}
