package main

import (
	"github.com/ilovelili/dongfeng/core-proxy/services/proxy/app"
	"github.com/ilovelili/dongfeng/sharedlib"
)

func main() {
	app := &app.App{}

	if err := app.Bootstarp(); err != nil {
		sharedlib.NewStdOutLogger().Panic(err.Error())
	}
}
