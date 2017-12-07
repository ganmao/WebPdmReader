package controllers

import (
	"WebPdmReader/lib"
	"strings"

	"github.com/astaxie/beego"
)

type SearchTable struct {
	beego.Controller
}

type SearchShow struct {
	TableCode   string
	TableName   string
	TableDomain string
	ColumnCode  string
	ColumnName  string
}

func (this *SearchTable) Get() {
	// 获取索引文件名
	tabIdxPath, err := beego.GetConfig("String", "IdxPath", "data/idx")
	if err == nil {
		beego.Debug("GET IdxPath : ", tabIdxPath)
	}
	WPRLibs.Check(err)

	tabIdxFile := tabIdxPath.(string) + WPRLibs.PATH_SPLIT + "SearchIndexes.xml"
	beego.Debug("tabIdxFile = ", tabIdxFile)
	tabIndexs := WPRLibs.ReadPdmIdxFile(tabIdxFile)

	indexs := []SearchShow{}

	// 根据传入表名过滤
	var inTab string
	this.Ctx.Input.Bind(&inTab, "tn")
	inTab = strings.Replace(inTab, "+", "", -1)
	inTab = strings.ToUpper(inTab)

	if inTab != "" {
		// 过滤表名
		for _, tn := range tabIndexs.TableIndex {
			if strings.Contains(strings.ToUpper(tn.TableCode), strings.ToUpper(inTab)) ||
				strings.Contains(strings.ToUpper(tn.TableName), strings.ToUpper(inTab)) {
				idx := SearchShow{}
				idx.TableCode = tn.TableCode
				idx.TableName = tn.TableName
				idx.TableDomain = tn.PdmFile

				indexs = append(indexs, idx)
			} else {
				continue
			}
		}
	}

	// 根据传入字段过滤
	var inOutCol string
	this.Ctx.Input.Bind(&inOutCol, "cn")
	inOutCol = strings.Replace(inOutCol, "+", "", -1)
	inOutCol = strings.ToUpper(inOutCol)

	if inOutCol != "" {
		// 过滤表名
		for _, cn := range tabIndexs.ColumnIndex {
			if strings.Contains(strings.ToUpper(cn.ColumnCode), strings.ToUpper(inOutCol)) ||
				strings.Contains(strings.ToUpper(cn.ColumnName), strings.ToUpper(inOutCol)) {
				idx := SearchShow{}
				idx.TableCode = cn.TableCode
				idx.TableName = cn.TableName
				idx.TableDomain = cn.PdmFile
				idx.ColumnCode = cn.ColumnCode
				idx.ColumnName = cn.ColumnName

				indexs = append(indexs, idx)
			} else {
				continue
			}
		}
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
	this.TplName = "searchTableList.tpl"

	// 子模板
	this.LayoutSections = make(map[string]string)
	// 页头
	this.LayoutSections["pageHeader"] = "pageHeader.tpl"
	// 当前页使用的script脚本
	this.LayoutSections["userScripts"] = "dataTables_script.tpl"
}
