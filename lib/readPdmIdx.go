package WPRLibs

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

type PrintTablesStruct struct {
	Tables []PrintTableStruct
}

type PrintTableStruct struct {
	TableName    string
	TableCode    string
	TableComment string
	TableColumns []PrintColumnStruct
}

type PrintColumnStruct struct {
	ColumnId                     string
	ColumnName                   string
	ColumnCode                   string
	ColumnDataType               string
	ColumnLength                 int
	ColumnIsPrimaryKey           bool
	ColumnMandatory              bool
	ColumnDefaultValue           string
	ColumnComment                string
	ColumnIndexCode              []string
	ColumnExtendedAttributesText string
}

type SearchIndexsStruct struct {
	Index []SearchIndexStruct
}

type SearchIndexStruct struct {
	TableCode   string
	TableName   string
	TableDomain string
}

func ReadPdmIdxFile(filePath string) []PrintTableStruct {
	var pts PrintTablesStruct

	idxFile, err := os.Open(filePath)
	defer idxFile.Close()
	Check(err)

	// 读取文件内容
	idxData, err := ioutil.ReadAll(idxFile)
	Check(err)

	// 解析文件内容
	err = xml.Unmarshal(idxData, &pts)
	Check(err)

	return pts.Tables
}

func ReadTabIdxFile(filePath string) []SearchIndexStruct {
	var s SearchIndexsStruct

	idxFile, err := os.Open(filePath)
	defer idxFile.Close()
	Check(err)

	// 读取文件内容
	idxData, err := ioutil.ReadAll(idxFile)
	Check(err)

	// 解析文件内容
	err = xml.Unmarshal(idxData, &s)
	Check(err)

	return s.Index
}
