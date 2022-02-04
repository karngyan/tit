package main

import (
	"fmt"
	"net/http"
	"time"
	"titd/models"
	"titd/modules/global"
	_ "titd/routers"
	"titd/settings"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
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

	// Models init - includes the tweet scheduler
	models.Init()

	// cors setup
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Content-Type", "Origin", "Authorization", "Access-Control-Allow-Origin", "X-User-ApiKey", "X-Token", "X-Ca-Token"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		MaxAge:           24 * time.Hour, // caps for 2 hours in chromium
	}), beego.WithReturnOnOutput(false))

	logs.Info("run mode =", beego.BConfig.RunMode)

	if beego.BConfig.RunMode == "prod" {
		// only for prod
		logs.Info("api server started at %v..", time.Now())
		logs.Info(fmt.Sprintf("starting titui on http://0.0.0.0:%s", "3001"))
		go serveTitUI("3001")
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

func serveTitUI(port string) {
	http.Handle("/", http.FileServer(http.Dir("./titui/dist")))
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		panic(err)
	}
}
