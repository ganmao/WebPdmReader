package main

import (
	"WebPdmReader/lib"
	_ "WebPdmReader/routers"

	"github.com/astaxie/beego"
)

func main() {
	// 设置日志级别
	logLevel, err := beego.GetConfig("String", "logLevel", "error")
	WPRLibs.Check(err)

	logPath, err := beego.GetConfig("String", "LogPath", "logs")
	WPRLibs.Check(err)

	logfile := `{"filename":"` + logPath.(string) + WPRLibs.PATH_SPLIT + `WPdmReader.log"}`

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
