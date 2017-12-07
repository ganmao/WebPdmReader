package WPRLibs

import (
	"testing"

	"github.com/astaxie/beego"
)

func Test_ListCurrDir(t *testing.T) {
	var isFind bool = false
	fList, err := ListCurrDir("./", ".go")
	Check(err)

	for _, f := range fList {
		t.Logf("找到文件 %v", f.FName)
		if f.FName == "procFiles_test.go" ||
			f.FName == "procFiles.go" {
			isFind = true
			break
		}

	}

	if !isFind {
		t.Error("没有找到指定的文件: ", "procFiles_test.go ", "procFiles.go")
	}
}

func Benchmark_ListCurrDir(b *testing.B) {
	b.StopTimer() //调用该函数停止压力测试的时间计数
	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能
	beego.SetLevel(beego.LevelError)
	b.StartTimer() //重新开始时间

	for i := 0; i < b.N; i++ {
		fList, err := ListCurrDir("./", ".go")
		Check(err)
		if len(fList) > 0 {
			b.Logf("执行第【%d】次", i)
		}
	}
}

func Test_UpdatePdmIndexes(t *testing.T) {
	var isSuccess bool = false
	ulog, err := UpdatePdmIndexes("../data/pdm", "../data/idx")
	Check(err)
	// 对输出数据进行效验
	for _, l := range ulog {
		if l.UpdateFileName == "../data/pdm/WPR_PhysicalDataModel_Test.pdm" {
			isSuccess = true
			break
		}
		t.Log("执行日志文件：", l.UpdateFileName)
	}

	t.Log("执行日志信息：", ulog)

	if !isSuccess {
		t.Error("执行错误！")
	}
}
