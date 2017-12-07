package WPRLibs

/* 功能说明：
 *  定义所有用到的结构体
 *
 *
 */
import (
	"encoding/xml"
	"os"
	"time"
)

type PdmRootModel struct {
	XMLName xml.Name `xml:"Model"`
	Root    PdmRoot  `xml:"RootObject"`
}

type PdmRoot struct {
	XMLName  xml.Name    `xml:"RootObject"`
	Id       string      `xml:"Id,attr"`
	Children PdmChildren `xml:"Children"`
}

type PdmChildren struct {
	XMLName xml.Name `xml:"Children"`
	Model   PdmModel `xml:"Model"`
}

// TODO: 看看是否可以用这个结构来简化
// type PdmComAttr struct {
// 	Id               string `xml:"Id,attr"`
// 	ObjectID         string `xml:"ObjectID"`
// 	CreationDate     string `xml:"CreationDate"`
// 	Creator          string `xml:"Creator"`
// 	ModificationDate string `xml:"ModificationDate"`
// 	Modifier         string `xml:"Modifier"`
// }

type PdmModel struct {
	XMLName          xml.Name      `xml:"Model"`
	Id               string        `xml:"Id,attr"`
	ObjectID         string        `xml:"ObjectID"`
	Name             string        `xml:"Name"`
	Code             string        `xml:"Code"`
	CreationDate     string        `xml:"CreationDate"`
	Creator          string        `xml:"Creator"`
	ModificationDate string        `xml:"ModificationDate"`
	Modifier         string        `xml:"Modifier"`
	Users            PdmUsers      `xml:"Users"`
	Tables           PdmTables     `xml:"Tables"`
	References       PdmReferences `xml:"References"`
}

type PdmUsers struct {
	XMLName xml.Name      `xml:"Users"`
	Users   []PdmUserInfo `xml:"User"`
}

type PdmUserInfo struct {
	XMLName          xml.Name `xml:"User"`
	Id               string   `xml:"Id,attr"`
	ObjectID         string   `xml:"ObjectID"`
	Name             string   `xml:"Name"`
	Code             string   `xml:"Code"`
	CreationDate     string   `xml:"CreationDate"`
	Creator          string   `xml:"Creator"`
	ModificationDate string   `xml:"ModificationDate"`
	Modifier         string   `xml:"Modifier"`
}

type PdmTables struct {
	XMLName xml.Name   `xml:"Tables"`
	Tables  []PdmTable `xml:"Table"`
}

type PdmTable struct {
	XMLName          xml.Name      `xml:"Table"`
	Id               string        `xml:"Id,attr"`
	ObjectID         string        `xml:"ObjectID"`
	Name             string        `xml:"Name"`
	Code             string        `xml:"Code"`
	CreationDate     string        `xml:"CreationDate"`
	Creator          string        `xml:"Creator"`
	ModificationDate string        `xml:"ModificationDate"`
	Modifier         string        `xml:"Modifier"`
	Comment          string        `xml:"Comment"`
	Columns          PdmColumns    `xml:"Columns"`
	Keys             PdmKeys       `xml:"Keys"`
	Indexes          PdmIndexes    `xml:"Indexes"`
	Owner            PdmTableOwner `xml:"Owner"`
	PrimaryKey       PdmPrimaryKey `xml:"PrimaryKey>Key"`
}

type PdmColumns struct {
	XMLName xml.Name    `xml:"Columns"`
	Columns []PdmColumn `xml:"Column"`
}

type PdmColumn struct {
	XMLName                xml.Name `xml:"Column"`
	Id                     string   `xml:"Id,attr"`
	ObjectID               string   `xml:"ObjectID"`
	Name                   string   `xml:"Name"`
	Code                   string   `xml:"Code"`
	CreationDate           string   `xml:"CreationDate"`
	Creator                string   `xml:"Creator"`
	ModificationDate       string   `xml:"ModificationDate"`
	Modifier               string   `xml:"Modifier"`
	Comment                string   `xml:"Comment"`
	LowValue               string   `xml:"LowValue"`
	HighValue              string   `xml:"HighValue"`
	DefaultValue           string   `xml:"DefaultValue"`
	DataType               string   `xml:"DataType"`
	Length                 string   `xml:"Length"`
	ColumnMandatory        string   `xml:"Column.Mandatory"`
	ExtendedAttributesText string   `xml:"ExtendedAttributesText"`
}

type PdmKeys struct {
	XMLName xml.Name `xml:"Keys"`
	Keys    []PdmKey `xml:"Key"`
}

type PdmKey struct {
	XMLName          xml.Name      `xml:"Key"`
	Id               string        `xml:"Id,attr"`
	ObjectID         string        `xml:"ObjectID"`
	Name             string        `xml:"Name"`
	Code             string        `xml:"Code"`
	CreationDate     string        `xml:"CreationDate"`
	Creator          string        `xml:"Creator"`
	ModificationDate string        `xml:"ModificationDate"`
	Modifier         string        `xml:"Modifier"`
	KeyColumns       PdmKeyColumns `xml:"Key.Columns"`
}

type PdmKeyColumns struct {
	XMLName    xml.Name       `xml:"Key.Columns"`
	KeyColumns []PdmComColumn `xml:"Column"`
}

type PdmComColumn struct {
	XMLName xml.Name `xml:"Column"`
	Ref     string   `xml:"Ref,attr"`
}

type PdmIndexes struct {
	XMLName xml.Name   `xml:"Indexes"`
	Indexes []PdmIndex `xml:"Index"`
}

type PdmIndex struct {
	XMLName          xml.Name        `xml:"Index"`
	Id               string          `xml:"Id,attr"`
	ObjectID         string          `xml:"ObjectID"`
	Name             string          `xml:"Name"`
	Code             string          `xml:"Code"`
	CreationDate     string          `xml:"CreationDate"`
	Creator          string          `xml:"Creator"`
	ModificationDate string          `xml:"ModificationDate"`
	Modifier         string          `xml:"Modifier"`
	IndexColumns     PdmIndexColumns `xml:"IndexColumns"`
}

type PdmIndexColumns struct {
	XMLName      xml.Name         `xml:"IndexColumns"`
	IndexColumns []PdmIndexColumn `xml:"IndexColumn"`
}

type PdmIndexColumn struct {
	XMLName          xml.Name     `xml:"IndexColumn"`
	Id               string       `xml:"Id,attr"`
	ObjectID         string       `xml:"ObjectID"`
	CreationDate     string       `xml:"CreationDate"`
	Creator          string       `xml:"Creator"`
	ModificationDate string       `xml:"ModificationDate"`
	Modifier         string       `xml:"Modifier"`
	IdxColumn        PdmIdxColumn `xml:"Column"`
}

type PdmIdxColumn struct {
	Columns []PdmComColumn `xml:"Column"`
}

type PdmTableOwner struct {
	XMLName xml.Name        `xml:"Owner"`
	User    PdmTabOwnerUser `xml:"User"`
}

type PdmTabOwnerUser struct {
	XMLName xml.Name `xml:"User"`
	Ref     string   `xml:"Ref,attr"`
}

type PdmPrimaryKey struct {
	Key string `xml:"Ref,attr"`
}

type PdmReferences struct {
	XMLName    xml.Name       `xml:"References"`
	References []PdmReference `xml:"Reference"`
}

type PdmReference struct {
	XMLName          xml.Name       `xml:"Reference"`
	Id               string         `xml:"Id,attr"`
	ObjectID         string         `xml:"ObjectID"`
	Name             string         `xml:"Name"`
	Code             string         `xml:"Code"`
	CreationDate     string         `xml:"CreationDate"`
	Creator          string         `xml:"Creator"`
	ModificationDate string         `xml:"ModificationDate"`
	Modifier         string         `xml:"Modifier"`
	Cardinality      string         `xml:"Cardinality"`
	UpdateConstraint string         `xml:"UpdateConstraint"`
	DeleteConstraint string         `xml:"DeleteConstraint"`
	ParentTable      PdmParentTable `xml:"ParentTable>Table"`
	ChildTable       PdmChildTable  `xml:"ChildTable>Table"`
	ParentKey        PdmParentKey   `xml:"ParentKey>Key"`
	Joins            PdmJoins       `xml:"Joins"`
}

type PdmParentTable struct {
	Ref string `xml:"Ref,attr"`
}

type PdmChildTable struct {
	Ref string `xml:"Ref,attr"`
}

type PdmParentKey struct {
	Ref string `xml:"Key;attr"`
}

type PdmJoins struct {
	XMLName xml.Name         `xml:"Joins"`
	Joins   PdmReferenceJoin `xml:"ReferenceJoin"`
}

type PdmReferenceJoin struct {
	XMLName          xml.Name     `xml:"ReferenceJoin"`
	Id               string       `xml:"Id,attr"`
	ObjectID         string       `xml:"ObjectID"`
	CreationDate     string       `xml:"CreationDate"`
	Creator          string       `xml:"Creator"`
	ModificationDate string       `xml:"ModificationDate"`
	Modifier         string       `xml:"Modifier"`
	Object1          PdmComColumn `xml:"Object1>Column"`
	Object2          PdmComColumn `xml:"Object2>Column"`
}

// 刷新Pdm日志，用于展示刷新进度
type UpdatePdmFileLog struct {
	UpdateFileName   string
	UpdateTabNum     int
	UpdateTabColNum  int
	UpdateTime       time.Time
	UpdateFileStatus os.FileInfo
}

// 解析输出结构，
type OutputTables struct {
	Tables []OutputTable
}

type OutputTable struct {
	TableName    string
	TableCode    string
	TableComment string // `xml:",innerxml"`
	TableOwner   string
	TableColumns []OutputColumn
}

type OutputColumn struct {
	ColumnId                     string
	ColumnName                   string
	ColumnCode                   string
	ColumnDataType               string
	ColumnLength                 int
	ColumnIsPrimaryKey           bool
	ColumnMandatory              bool
	ColumnDefaultValue           string
	ColumnComment                string // `xml:",innerxml"`
	ColumnIndexCode              []string
	ColumnkeyCode                []string
	ColumnExtendedAttributesText string
}

// 搜索使用的索引
type SearchIndex struct {
	TableIndex  []SearchTableIndex
	ColumnIndex []SearchColumnIndex
}

type SearchTableIndex struct {
	TableCode string
	TableName string
	PdmFile   string
}

type SearchColumnIndex struct {
	ColumnCode string
	ColumnName string
	TableCode  string
	TableName  string
	PdmFile    string
}
