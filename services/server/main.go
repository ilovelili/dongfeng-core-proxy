package main

import (
	"fmt"
	"time"

	"github.com/ilovelili/dongfeng-core-proxy/services/server/app"
	logger "github.com/ilovelili/dongfeng-logger"
)

func main() {
	app := &app.App{}

	if err := app.Bootstarp(); err != nil {
		errorlog := &logger.Log{
			Category: "ErrorLog",
			Content:  fmt.Sprintf("Core proxy bootstrap failed: %s\n", err.Error()),
			Time:     time.Now(),
		}

		errorlog.ErrorLog(logger.CoreProxy)
	}
}
