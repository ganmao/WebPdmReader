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

	// TODO: 如何将这个日志的路径和文件名采用变量传入？
	beego.SetLogger("file", `{"filename":"logs/WPdmReader.log","level":5,"maxlines":0,"maxsize":20,"daily":true,"maxdays":10}`)
	beego.SetLogFuncCall(false)

	switch logLevel {
	case "error":
		beego.SetLevel(beego.LevelError)
		beego.SetLogFuncCall(true)
	case "info":
		beego.SetLevel(beego.LevelInformational)
	case "debug":
		beego.SetLevel(beego.LevelDebug)
		beego.SetLogFuncCall(true)
	default:
		beego.SetLevel(beego.LevelError)
	}

	beego.Run()
}
