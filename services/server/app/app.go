package app

import (
	"github.com/ilovelili/dongfeng-core-proxy/services/utils" // k8s registry plugin
	_ "github.com/micro/go-plugins/registry/kubernetes"
	web "github.com/micro/go-web"
)

// App app. They call me God Object so I guess I am cool
type App struct {
	Router  *Router
	Service web.Service
}

// Bootstarp Bootstarp the service
func (app *App) Bootstarp() error {
	myapp, err := app.init()

	if err != nil {
		return err
	}

	myapp.Service.Handle("/", myapp.Router.Route())
	return myapp.Service.Run()
}

// init init the app
func (app *App) init() (application *App, err error) {
	if application, err = app.initializeRouter(); err != nil {
		return app, err
	}

	if application, err = app.initializeProxyService(); err != nil {
		return app, err
	}

	return application, err
}

// initializeRouter Initialize router
func (app *App) initializeRouter() (*App, error) {
	app.Router = NewRouter()
	return app, nil
}

// initializeProxyService init reverse proxy service with router
func (app *App) initializeProxyService() (*App, error) {
	config := utils.GetConfig()

	service := web.NewService(
		web.Name(config.ServiceNames.CoreProxy),
		web.RegisterTTL(config.ServiceMeta.GetRegistryTTL()),
		web.RegisterInterval(config.ServiceMeta.GetRegistryHeartbeat()),
		web.Version(config.ServiceMeta.GetVersion()),
	)

	err := service.Init()
	app.Service = service
	return app, err
}
