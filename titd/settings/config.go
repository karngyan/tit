package settings

import (
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

var (
	DbPath  string
	LogFile string
)

func LoadConfig() (err error) {
	if DbPath, err = beego.AppConfig.String("dbPath"); err != nil {
		return err
	}
	logs.Info("DbPath:", DbPath)

	if LogFile, err = beego.AppConfig.String("logFile"); err != nil {
		return err
	}
	logs.Info("LogFile:", LogFile)

	return err
}
