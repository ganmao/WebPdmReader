package WRPLibs

import (
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

const (
	IS_DIRECTORY = iota
	IS_REGULAR
	IS_SYMLINK
)

type PdmFile struct {
	FType  int
	FName  string
	FSize  int64
	FMtime time.Time
	FMode  os.FileMode
	FPath  string
}

// 获取指定目录下的所有文件，不进入下一级目录搜索，可以匹配后缀过滤。
func ListDir(dirPth string, suffix string) (files []PdmFile, err error) {
	beego.Debug("dirPth", dirPth)
	beego.Debug("suffix", suffix)
	files = make([]PdmFile, 0, 50)
	p := PdmFile{}

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
