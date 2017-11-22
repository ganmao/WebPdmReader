package routers

import (
	"WPdmReader/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.ShowPdmController{})
	beego.Router("/pdm", &controllers.ShowPdmController{})
	beego.Router("/tab", &controllers.ShowTabController{})
	beego.Router("/st", &controllers.SearchTable{})
	beego.Router("/mng", &controllers.PdmManager{})
}
