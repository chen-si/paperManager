package routers

import (
	"paperManger/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/getPagePapers", &controllers.MainController{})
}
