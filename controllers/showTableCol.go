package controllers

import (
	"WPdmReader/lib"

	"github.com/astaxie/beego"
)

type ShowTabController struct {
	beego.Controller
}

func (this *ShowTabController) Get() {
	var IsShowTableDetail bool = false
	var IsShowTableList bool = false
	var pdmFileName string
	var tableName string

	// 获取索引文件名
	this.Ctx.Input.Bind(&pdmFileName, "name")
	if pdmFileName != "" {
		IsShowTableList = true
		beego.Debug("Get Pdm File Name :", pdmFileName)
	}

	// 获取表名称
	this.Ctx.Input.Bind(&tableName, "tab")
	if tableName != "" {
		IsShowTableDetail = true
		beego.Debug("Get Tab Name :", tableName)
	}

	/* 设置table List */
	if IsShowTableList {
		// TODO:看一下这两个替换进行优化
		tabIdxPath, err := beego.GetConfig("String", "IdxPath", "data/idx")
		beego.Debug("GET Config IdxPath : ", tabIdxPath)
		WRPLibs.Check(err)
		pdmIdxFile := tabIdxPath.(string) + WRPLibs.PATH_SPLIT + pdmFileName + ".xml"
		beego.Debug("Set Index File Name : ", pdmIdxFile)

		tables := WRPLibs.ReadPdmIdxFile(pdmIdxFile)

		/* 模板赋值 */
		// --- contents
		this.Data["tables"] = tables
		this.Data["pdmFileName"] = pdmFileName

		// --- page header
		this.Data["pageHeader"] = "Show Table List"
		this.Data["pDescription"] = pdmFileName
		this.Data["pLevel"] = "PDMLIST"
		this.Data["pLevelLink"] = "pdm"
		this.Data["pHere"] = "TABLIST"

		/* 设置table detail */
		if IsShowTableDetail {
			for _, tab := range tables {
				if tableName == tab.TableCode {
					this.Data["table"] = tab
					IsShowTableList = false
					// --- page header
					this.Data["pageHeader"] = "Show Table Detail"
					this.Data["pDescription"] = pdmFileName + " > " + tab.TableCode
					this.Data["pLevel"] = "PDM LIST"
					this.Data["pLevelLink"] = "pdm"
					this.Data["pLevel1"] = "TABLE LIST"
					this.Data["pLevelLink1"] = "tab?name=" + pdmFileName
					this.Data["pHere"] = "TABLIST"
					break
				}
			}
		}
	}

	this.Data["IsShowTableList"] = IsShowTableList
	this.Data["IsShowTableDetail"] = IsShowTableDetail

	// 界面展示
	// 模板
	this.Layout = "layout.html"
	// 主要展示内容
	this.TplName = "showTabList.tpl"

	// 子模板
	this.LayoutSections = make(map[string]string)
	// 页头
	this.LayoutSections["pageHeader"] = "pageHeader.tpl"
	// 当前页使用的script脚本
	this.LayoutSections["userScripts"] = "dataTables_script.tpl"
}
