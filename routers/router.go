package routers

import (
	"github.com/astaxie/beego"
	"paperManger/controllers"
)

func init() {
	beego.Router("/", &controllers.LoginController{})
	beego.Router("/getPagePapers", &controllers.MainController{})
	beego.Router("/getPageUsers", &controllers.UserManagerController{})
	beego.Router("/deleteUser", &controllers.DeleteUserController{})
	beego.Router("/updateOrAddUser", &controllers.UpdateOrAddUserController{})
}
