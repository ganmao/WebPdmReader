package controllers

import (
	"WebPdmReader/lib"

	"github.com/astaxie/beego"
)

type ShowPdmController struct {
	beego.Controller
}

func (this *ShowPdmController) Get() {
	// 获取查询pdm文件列表
	pdmPath, err := beego.GetConfig("String", "PdmPath", "data/pdm")
	if err == nil {
		beego.Debug("GET PdmPath : ", pdmPath)
	}
	WPRLibs.Check(err)

	fList, err := WPRLibs.ListCurrDir(pdmPath.(string), "pdm")
	WPRLibs.Check(err)

	// 数据赋值
	// --- contents
	this.Data["fileList"] = fList

	// --- page header
	this.Data["pageHeader"] = "Show PDM List"
	this.Data["pDescription"] = "当前目录下的pdm文件"
	this.Data["pLevel"] = "HOME"
	this.Data["pLevelLink"] = "/"
	this.Data["pHere"] = "PDMLIST"

	// 界面展示
	// 模板
	this.Layout = "layout.html"
	// 主要展示内容
	this.TplName = "showPdmList.tpl"

	// 子模板
	this.LayoutSections = make(map[string]string)
	// 页头
	this.LayoutSections["pageHeader"] = "pageHeader.tpl"
	// 当前页使用的script脚本
	this.LayoutSections["userScripts"] = "dataTables_script.tpl"
}
