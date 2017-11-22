package main

import (
	"WPdmReader/lib"
	_ "WPdmReader/routers"

	"github.com/astaxie/beego"
)

func main() {
	// 设置日志级别
	logLevel, err := beego.GetConfig("String", "logLevel", "error")
	WRPLibs.Check(err)

	logPath, err := beego.GetConfig("String", "LogPath", "logs")
	WRPLibs.Check(err)

	logfile := `{"filename":"` + logPath.(string) + WRPLibs.PATH_SPLIT + `WPdmReader.log"}`

	beego.SetLogger("file", logfile)

	switch logLevel {
	case "error":
		beego.SetLevel(beego.LevelError)
	case "info":
		beego.SetLevel(beego.LevelInformational)
	case "debug":
		beego.SetLevel(beego.LevelDebug)
	default:
		beego.SetLevel(beego.LevelError)
	}

	beego.Run()
}
