package routers

import (
	"github.com/astaxie/beego"
	"paperManger/controllers"
)

func init() {
	beego.Router("/getPagePapers", &controllers.MainController{})

	//User
	beego.Router("/", &controllers.UserController{},"get:ToLogin;post:Login")
	beego.Router("/getPageUsers", &controllers.UserController{},"*:PageUsers")
	beego.Router("/deleteUser", &controllers.UserController{},"*:DeleteUser")
	beego.Router("/updateOrAddUser", &controllers.UserController{},"get:ToUpdateOrAddUser;post:UpdateOrAddUser")

	//Tutor
	beego.Router("/getPageTutors", &controllers.TutorController{},"*:PageTutors")
	beego.Router("/deleteTutor", &controllers.TutorController{},"*:DeleteTutor")
	beego.Router("/updateOrAddTutor", &controllers.TutorController{},"get:ToUpdateOrAddTutor;post:UpdateOrAddTutor")
}
