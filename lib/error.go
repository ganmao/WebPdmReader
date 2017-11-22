package WRPLibs

import "github.com/astaxie/beego"

func Check(e error) {
	if e != nil {
		// fmt.Printf("error: %v \n", e)
		beego.Error(e)
		panic(e)
	}
}
