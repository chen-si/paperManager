package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) ToMainPage() {
	c.TplName = "main_page.html"
}
