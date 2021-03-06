package routers

import (
	"github.com/astaxie/beego"
	"paperManger/controllers"
)

func init() {
	//Paper
	beego.Router("/getPagePapers", &controllers.PaperController{}, "*:PagePapers")
	beego.Router("/deletePaper", &controllers.PaperController{}, "*:DeletePaper")
	beego.Router("/updateOrAddPaper", &controllers.PaperController{}, "get:ToUpdateOrAddPaper;post:UpdateOrAddPaper")
	beego.Router("/submitPaper", &controllers.PaperController{}, "get:ToUploadFile;post:UploadFile")
	beego.Router("/downloadPaper", &controllers.PaperController{}, "*:DownloadFile")
	beego.Router("/failedPapers", &controllers.PaperController{}, "*:ToFailedPaper")
	beego.Router("/updateFailedPaper", &controllers.PaperController{}, "post:UpdateFailedPaper;get:ToUpdateFailedPaper")

	//User
	beego.Router("/", &controllers.UserController{}, "get:ToLogin;post:Login")
	beego.Router("/logout", &controllers.UserController{}, "*:Logout")
	beego.Router("/main", &controllers.UserController{}, "*:MainPage")
	beego.Router("/getPageUsers", &controllers.UserController{}, "*:PageUsers")
	beego.Router("/deleteUser", &controllers.UserController{}, "*:DeleteUser")
	beego.Router("/updateOrAddUser", &controllers.UserController{}, "get:ToUpdateOrAddUser;post:UpdateOrAddUser")

	//Tutor
	beego.Router("/getPageTutors", &controllers.TutorController{}, "*:PageTutors")
	beego.Router("/deleteTutor", &controllers.TutorController{}, "*:DeleteTutor")
	beego.Router("/updateOrAddTutor", &controllers.TutorController{}, "get:ToUpdateOrAddTutor;post:UpdateOrAddTutor")
	beego.Router("/getTutorGraInfo", &controllers.TutorController{}, "*:GetTutorGraInfo")
	beego.Router("/getTutorGraPaperInfo", &controllers.TutorController{}, "*:GetTutorGraPaperInfo")

	//Graduate
	beego.Router("/getPageGraduates", &controllers.GraduateController{}, "*:PageGraduates")
	beego.Router("/deleteGraduate", &controllers.GraduateController{}, "*:DeleteGraduate")
	beego.Router("/updateOrAddGraduate", &controllers.GraduateController{}, "get:ToUpdateOrAddGraduate;post:UpdateOrAddGraduate")
	beego.Router("/getGraPaperInfo", &controllers.GraduateController{}, "*:GetGraPaperInfo")
}
