package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"paperManger/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
	pDao := &models.PaperDao{
		Db: models.Db,
	}
	page, err := pDao.GetPagePapers("1")
	if err != nil {
		fmt.Println(err)
		return
	}
	c.TplName = "book_manager.html"
	c.Data["Page"] = page
}
