package controllers

import (
	"WPdmReader/lib"
	"strings"

	"github.com/astaxie/beego"
)

type SearchTable struct {
	beego.Controller
}

func (this *SearchTable) Get() {
	// 获取索引文件名
	tabIdxPath, err := beego.GetConfig("String", "IdxPath", "data/idx")
	if err == nil {
		beego.Debug("GET IdxPath : ", tabIdxPath)
	}
	WRPLibs.Check(err)
	tabIdxFile := tabIdxPath.(string) + WRPLibs.PATH_SPLIT + "idx_tab.xml"

	beego.Debug("tabIdxFile = ", tabIdxFile)
	allIndexs := WRPLibs.ReadTabIdxFile(tabIdxFile)

	// 根据传入表名过滤
	var inputTableName string
	var indexs []WRPLibs.SearchIndexStruct
	this.Ctx.Input.Bind(&inputTableName, "tn")
	inputTableName = strings.ToUpper(inputTableName)
	if inputTableName != "" {
		for _, tn := range allIndexs {
			if strings.Contains(tn.TableCode, inputTableName) || strings.Contains(tn.TableName, inputTableName) {
				indexs = append(indexs, tn)
			} else {
				continue
			}
		}
	} else {
		indexs = allIndexs
	}

	/* 模板赋值 */
	// --- contents
	this.Data["indexs"] = indexs

	// --- page header
	this.Data["pageHeader"] = "Show Table Idex List"
	this.Data["pLevel"] = "Search Table"
	this.Data["pLevelLink"] = "st"
	this.Data["pHere"] = "Table Index LIST"

	// 界面展示
	// 模板
	this.Layout = "layout.html"
	// 主要展示内容
	this.TplName = "showTableIndex.tpl"

	// 子模板
	this.LayoutSections = make(map[string]string)
	// 页头
	this.LayoutSections["pageHeader"] = "pageHeader.tpl"
	// 当前页使用的script脚本
	this.LayoutSections["userScripts"] = "dataTables_script.tpl"
}
