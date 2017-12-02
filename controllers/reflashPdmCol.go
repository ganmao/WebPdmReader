package controllers

import (
	"WebPdmReader/lib"

	"github.com/astaxie/beego"
)

type PdmManager struct {
	beego.Controller
}

func (this *PdmManager) Get() {
	var inputMngCmd string = ""
	var IsNeedSubmit bool = true
	var IsCmdOutput bool = false
	var refreshLog []WPRLibs.RefreshPdmLog
	this.Ctx.Input.Bind(&inputMngCmd, "cmd")

	// TODO:form会传入但代码暂未获取IdxPath，PdmPath；等有需求再实现
	idxPath, err := beego.GetConfig("String", "IdxPath", "data/idx")
	if err == nil {
		beego.Debug("GET IdxPath : ", idxPath)
	}
	WPRLibs.Check(err)

	pdmPath, err := beego.GetConfig("String", "PdmPath", "data/idx")
	if err == nil {
		beego.Debug("GET PdmPath : ", pdmPath)
	}
	WPRLibs.Check(err)

	if inputMngCmd == "refresh" {
		refreshLog = WPRLibs.RefreshPdmTableIdx(pdmPath.(string), idxPath.(string))
		IsCmdOutput = true
		IsNeedSubmit = false
	}

	/* 模板赋值 */
	// --- contents
	this.Data["IsNeedSubmit"] = IsNeedSubmit
	this.Data["IsCmdOutput"] = IsCmdOutput
	this.Data["pdmPath"] = pdmPath
	this.Data["indexPath"] = idxPath
	this.Data["refreshLog"] = refreshLog

	// --- page header
	this.Data["pageHeader"] = "Manager Pdm"
	this.Data["pLevel"] = "Manager Pdm"
	this.Data["pLevelLink"] = "mng"
	this.Data["pHere"] = "Refresh Pdm Index"

	// 界面展示
	// 模板
	this.Layout = "layout.html"
	// 主要展示内容
	this.TplName = "pdmManager.tpl"

	// 子模板
	this.LayoutSections = make(map[string]string)
	// 页头
	this.LayoutSections["pageHeader"] = "pageHeader.tpl"
	// 当前页使用的script脚本
	this.LayoutSections["userScripts"] = "dataTables_script.tpl"
}
