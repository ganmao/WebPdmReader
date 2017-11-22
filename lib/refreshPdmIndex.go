package WRPLibs

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

type RefreshPdmLog struct {
	RefreshFileName   string
	RefreshTabNum     int
	RefreshTabColumns int
	// RefreshTime       string
	RefreshTime   time.Time
	RefreshStatus os.FileInfo
}

type TablesStruct struct {
	// Description string   `xml:",innerxml"`
	Tables []TableStruct `xml:"RootObject>Children>Model>Tables>Table"`
}

type TableStruct struct {
	XMLName      xml.Name         `xml:"Table"`
	TableName    string           `xml:"Name"`
	TableCode    string           `xml:"Code"`
	TableComment string           `xml:"Comment"`
	TableColumns []ColumnStruct   `xml:"Columns>Column"`
	TableKeys    []TableKeyStruct `xml:"Keys>Key"`
	TableIndexs  []IndexStruct    `xml:"Indexes>Index"`
}

type ColumnStruct struct {
	XMLName                      xml.Name `xml:"Column"`
	ColumnId                     string   `xml:"Id,attr"`
	ColumnName                   string   `xml:"Name"`
	ColumnCode                   string   `xml:"Code"`
	ColumnComment                string   `xml:"Comment"`
	ColumnDefaultValue           string   `xml:"DefaultValue"`
	ColumnDataType               string   `xml:"DataType"`
	ColumnLength                 string   `xml:"Length"`
	ColumnMandatory              string   `xml:"Column.Mandatory"`
	ColumnExtendedAttributesText string   `xml:"ExtendedAttributesText"`
}

type TableKeyStruct struct {
	XMLName    xml.Name          `xml:"Key"`
	KeyId      string            `xml:"Id,attr"`
	KeyName    string            `xml:"Name"`
	KeyCode    string            `xml:"Code"`
	KeyColumns []KeyColumnStruct `xml:"Key.Columns>Column"`
}

type KeyColumnStruct struct {
	RefColumn string `xml:"Ref,attr"`
}

type IndexStruct struct {
	XMLName      xml.Name            `xml:"Index"`
	IndexId      string              `xml:"Id,attr"`
	IndexName    string              `xml:"Name"`
	IndexCode    string              `xml:"Code"`
	IndexColumns []IndexColumnStruct `xml:"IndexColumns>IndexColumn"`
}

type IndexColumnStruct struct {
	XMLName               xml.Name       `xml:"IndexColumn"`
	IndexColumnId         string         `xml:"Id,attr"`
	IndexAscending        string         `xml:"Ascending"`
	IndexColumnExpression string         `xml:"IndexColumn.Expression"`
	IndexColumns          []IdxColStruct `xml:"Column>Column"`
}

type IdxColStruct struct {
	RefColumn string `xml:"Ref,attr"`
}

func (ts *TablesStruct) Exchange() (ops PrintTablesStruct, colNum int) {
	colNum = 0
	for _, t := range ts.Tables {

		pts := PrintTableStruct{}
		// 表基本信息赋值
		pts.TableName = t.TableName
		pts.TableCode = t.TableCode
		pts.TableComment = strings.Replace(t.TableComment, "&#xA;", "<br>", -1)

		for _, c := range t.TableColumns {
			tableColumn := PrintColumnStruct{}
			// 字段信息赋值
			tableColumn.ColumnId = c.ColumnId
			tableColumn.ColumnName = c.ColumnName
			tableColumn.ColumnCode = c.ColumnCode
			tableColumn.ColumnDataType = c.ColumnDataType
			tableColumn.ColumnLength, _ = strconv.Atoi(c.ColumnLength)
			tableColumn.ColumnDefaultValue = c.ColumnDefaultValue
			tableColumn.ColumnMandatory = (c.ColumnMandatory == "1")
			tableColumn.ColumnComment = strings.Replace(c.ColumnComment, "&#xA;", "<br>", -1)
			// pcs.ColumnExtendedAttributesText = c.ColumnExtendedAttributesText

			// 设置主键
			// fmt.Println("ColumnId", tableColumn.ColumnId, "t.TableKeys", t.TableKeys)
			if len(t.TableKeys) > 0 {
				for _, key := range t.TableKeys {
					// fmt.Println("tableColumn.ColumnId", tableColumn.ColumnId, "key.KeyId", key.KeyId)
					for _, ref := range key.KeyColumns {
						if ref.RefColumn == tableColumn.ColumnId {
							tableColumn.ColumnIsPrimaryKey = true
							goto FIND_KEY
						}
					}
				}
			FIND_KEY:
			}

			// 设置索引
			if len(t.TableIndexs) > 0 {
				indexCode := []string{}
				for _, idx := range t.TableIndexs {
					for _, idxCol := range idx.IndexColumns {
						if len(idxCol.IndexColumnExpression) > 0 {
							indexCode = append(indexCode, idxCol.IndexColumnExpression)
							goto NEXT_INDEX
						} else if len(idxCol.IndexColumns) > 0 {
							for _, idxRef := range idxCol.IndexColumns {
								if idxRef.RefColumn == tableColumn.ColumnId {
									indexCode = append(indexCode, idx.IndexName)
									goto NEXT_INDEX
								}
							}
						} else {
							fmt.Println("tableColumns error")
						}
					}
				NEXT_INDEX:
				}
				tableColumn.ColumnIndexCode = indexCode
			}
			pts.TableColumns = append(pts.TableColumns, tableColumn)
			colNum++
		}
		ops.Tables = append(ops.Tables, pts)
	}

	return ops, colNum
}

func analysisFile(filename string) (outStruct PrintTablesStruct, rp RefreshPdmLog, err error) {
	rp = RefreshPdmLog{}
	rp.RefreshTime = time.Now()
	// 打开文件
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error: %v", err)
		return PrintTablesStruct{}, rp, err
	}
	defer file.Close()
	rp.RefreshFileName = file.Name()
	rp.RefreshStatus, err = file.Stat()
	Check(err)

	// 读取文件内容
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return PrintTablesStruct{}, rp, err
	}

	// 解析文件内容
	v := TablesStruct{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return PrintTablesStruct{}, rp, err
	}
	rp.RefreshTabNum = len(v.Tables)

	// 转为内部结构定义
	interTableStatus, num := v.Exchange()
	rp.RefreshTabColumns = num

	return interTableStatus, rp, nil
}

func writeXml(filename string, s string) {
	var err error
	var f *os.File
	f, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666) //打开文件
	defer f.Close()

	Check(err)
	n, err := io.WriteString(f, s) //写入文件(字符串)
	Check(err)
	fmt.Printf("写入 %s 文件 %d 字节\n", f.Name(), n)
}

func writeIdx(filename string, s SearchIndexsStruct) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666) //打开文件
	Check(err)
	defer f.Close()
	// fmt.Println("s:", s)

	o, err := xml.MarshalIndent(s, "", "  ")
	Check(err)
	// fmt.Println("o:", string(o))

	_, err = io.WriteString(f, string(o)+"\n") //写入文件(字符串)
	Check(err)
}

func RefreshPdmTableIdx(pdmPath string, idxPath string) []RefreshPdmLog {
	var SISS SearchIndexsStruct
	var RPLS []RefreshPdmLog
	// 打开目录
	fileList, err := ListDir(pdmPath, "pdm")
	Check(err)

	for _, f := range fileList {
		o, r, err := analysisFile(f.FPath)
		Check(err)

		// 转为XML格式输出
		output, err := xml.MarshalIndent(o, "", "  ")
		// 写出文件
		writeXml(idxPath+PATH_SPLIT+f.FName+".xml", string(output))

		// 写入索引
		for _, p := range o.Tables {
			var sis SearchIndexStruct
			sis.TableDomain = f.FName
			sis.TableCode = p.TableCode
			sis.TableName = p.TableName
			SISS.Index = append(SISS.Index, sis)

			// fmt.Println("SIS:", SIS)
		}
		RPLS = append(RPLS, r)
	}
	writeIdx(idxPath+PATH_SPLIT+"idx_tab.xml", SISS)

	return RPLS
}
