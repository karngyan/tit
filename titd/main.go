package main

import (
	"fmt"
	"time"
	"titd/modules/global"
	_ "titd/routers"
	"titd/settings"

	"github.com/beego/beego/v2/core/logs"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {

	defer func() {
		logs.Warn("shutting down api_server at: %v", time.Now().Unix())
	}()

	var err error

	// load settings
	if err = settings.LoadConfig(); err != nil {
		panic(err)
	}

	// setup logger
	setupLogger()

	// Global init
	global.Init()

	// swagger docs
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/v1/swagger"] = "swagger"

	logs.Info("run mode =", beego.BConfig.RunMode)

	if beego.BConfig.RunMode == "prod" {
		logs.Info("api server started at %v..", time.Now())
	}

	beego.Run()
}

func setupLogger() {
	logs.Async()
	if beego.BConfig.RunMode == "prod" {
		_ = logs.SetLogger(logs.AdapterFile, fmt.Sprintf(`{"filename": "%s"}`, settings.LogFile))
	} else {
		_ = logs.SetLogger(logs.AdapterConsole)
	}
}
