package WPRLibs

/* 功能说明：
 *  负责处理文件相应的读取，写入，查找
 *
 *
 */
import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

const (
	IS_DIRECTORY = iota
	IS_REGULAR
	IS_SYMLINK
)

type pFile struct {
	FType  int
	FName  string
	FSize  int64
	FMtime time.Time
	FMode  os.FileMode
	FPath  string
}

var ostype = os.Getenv("GOOS") // 获取系统类型
var listfile []pFile           //获取文件列表

/* 功能说明：
 *  获取指定目录下的所有文件，不进入下一级目录搜索，可以匹配后缀过滤。
 *
 *
 */
func ListCurrDir(dirPth string, suffix string) (files []pFile, err error) {
	beego.Debug("dirPth", dirPth)
	beego.Debug("suffix", suffix)
	files = make([]pFile, 0, 50)
	p := pFile{}

	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			p.FType = IS_DIRECTORY
			continue
		} else if (fi.Mode() & os.ModeSymlink) > 0 {
			p.FType = IS_REGULAR
			continue
		} else {
			p.FType = IS_SYMLINK
		}

		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			p.FName = fi.Name()
			p.FMtime = fi.ModTime()
			p.FSize = fi.Size()
			p.FMode = fi.Mode()
			p.FPath = dirPth + PATH_SPLIT + fi.Name()
			files = append(files, p)
		}
	}
	beego.Debug("files", files)
	return files, nil
}

/* 功能说明：
 *  获取指定目录下的所有文件，针对子目录一起搜索，可以匹配后缀过滤。
 *
 *
 */
func ListSubDir(dirPth string, suffix string) (files []pFile, err error) {
	beego.Debug("dirPth", dirPth)
	beego.Debug("suffix", suffix)
	listfile = []pFile{}

	//var strRet string
	err = filepath.Walk(dirPth, getFileInfo)
	Check(err)

	return listfile, err
}

/* 功能说明：
 *  根据walk扫描到的文件进行处理，放入listfile中返回。
 *
 *
 */
func getFileInfo(path string, f os.FileInfo, err error) error {
	var strRet string

	if f == nil {
		return err
	}
	if f.IsDir() {
		return nil
	}

	//用strings.HasSuffix(src, suffix)//判断src中是否包含 suffix结尾
	ok := strings.HasSuffix(strings.ToUpper(strRet), ".PDM")
	if ok {
		var fInfo pFile
		fInfo.FMode = f.Mode()
		fInfo.FName = f.Name()
		fInfo.FMtime = f.ModTime()
		fInfo.FPath = path
		fInfo.FSize = f.Size()

		if f.IsDir() { // 忽略目录
			fInfo.FType = IS_DIRECTORY
		} else if (f.Mode() & os.ModeSymlink) > 0 {
			fInfo.FType = IS_REGULAR
		} else {
			fInfo.FType = IS_SYMLINK
		}

		listfile = append(listfile, fInfo) //将目录push到listfile []string中
		beego.Debug("add file :", path)
	}

	return nil
}

func writeXml(filename string, s string) {
	var err error
	var f *os.File

	//打开文件
	f, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	Check(err)
	defer f.Close()

	n, err := io.WriteString(f, s) //写入文件(字符串)
	Check(err)
	beego.Debug("写入[", f.Name(), "]文件[", n, "]字节")
}

/* 功能说明：
 *  更新指定目录下PDM文件相应索引
 *
 *
 */
func UpdatePdmIndexes(pdmPath string, idxPath string) (uplog []UpdatePdmFileLog, err error) {
	var sIdx SearchIndex
	uplog = []UpdatePdmFileLog{}
	// 打开目录，查找所有需要处理的文件
	fileList, err := ListCurrDir(pdmPath, "pdm")
	Check(err)

	// 针对每个文件进行解析，写出查找内容到xml文件
	for _, f := range fileList {
		pdmStruct, ulog, err := decodePDMFile(f.FPath)
		Check(err)

		outstruct, err := pdmStruct.ExchangeToOutput(&ulog)

		// 针对文件内的表生成表索引
		for _, t := range outstruct.Tables {
			sTabIdx := SearchTableIndex{}
			sTabIdx.PdmFile = f.FName
			sTabIdx.TableCode = t.TableCode
			sTabIdx.TableName = t.TableName

			sIdx.TableIndex = append(sIdx.TableIndex, sTabIdx)

			// 针对文件内的表生成字段索引
			for _, c := range t.TableColumns {
				sColIdx := SearchColumnIndex{}
				sColIdx.PdmFile = f.FName
				sColIdx.TableCode = t.TableCode
				sColIdx.TableName = t.TableName
				sColIdx.ColumnCode = c.ColumnCode
				sColIdx.ColumnName = c.ColumnName

				sIdx.ColumnIndex = append(sIdx.ColumnIndex, sColIdx)
			}
		}

		// 转为XML格式输出
		outString, err := xml.MarshalIndent(outstruct, "", "  ")
		// 写出文件
		writeXml(idxPath+PATH_SPLIT+f.FName+".xml", string(outString))
		// 日志汇总
		uplog = append(uplog, ulog)
	}

	// 转为XML格式输出
	outIdxString, err := xml.MarshalIndent(sIdx, "", "  ")
	// 写出文件
	writeXml(idxPath+PATH_SPLIT+"SearchIndexes.xml", string(outIdxString))

	return uplog, nil
}

func ReadPdmIdxFile(filePath string) SearchIndex {
	var sIdx SearchIndex

	idxFile, err := os.Open(filePath)
	defer idxFile.Close()
	Check(err)

	// 读取文件内容
	idxData, err := ioutil.ReadAll(idxFile)
	Check(err)

	// 解析文件内容
	err = xml.Unmarshal(idxData, &sIdx)
	Check(err)

	return sIdx
}

func ReadPdmXmlFile(filePath string) []OutputTable {
	var oTab OutputTables

	idxFile, err := os.Open(filePath)
	defer idxFile.Close()
	Check(err)

	// 读取文件内容
	idxData, err := ioutil.ReadAll(idxFile)
	Check(err)

	// 解析文件内容
	err = xml.Unmarshal(idxData, &oTab)
	Check(err)

	return oTab.Tables
}
