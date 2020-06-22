package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"paperManger/models"
)

type GraduateController struct {
	beego.Controller
}

func (c *GraduateController) PageGraduates() {
	gDao := &models.GraduatesDao{
		Db: models.Db,
	}
	pageNo := c.GetString("pageNo")
	var page *models.Page
	var err error
	if pageNo == "" {
		page, err = gDao.GetPageGraduates("1")
	} else {
		page, err = gDao.GetPageGraduates(pageNo)
	}

	if err != nil {
		fmt.Println(err)
		return
	}
	c.TplName = "graduate/page_graduates.html"
	c.Data["Page"] = page
}

func (c *GraduateController) ToUpdateOrAddGraduate() {
	gDao := &models.GraduatesDao{
		Db: models.Db,
	}

	id := c.GetString("id")
	graduate := &models.GraduatesInfo{
		Id: id,
	}
	if id == "" {
		//add user
		c.TplName = "graduate/graduate_edit.html"
	} else {
		//update user
		err := gDao.Read(graduate)
		if err != nil {
			fmt.Println(err)
			return
		}
		c.TplName = "graduate/graduate_edit.html"
		c.Data["Graduate"] = graduate
	}
}

func (c *GraduateController) UpdateOrAddGraduate() {
	graduate := &models.GraduatesInfo{
		Id:           c.GetString("id"),
		Name:         c.GetString("name"),
		GraduateTime: c.GetString("graduatetime"),
		TutorId:      c.GetString("tutorid"),
	}

	gDao := &models.GraduatesDao{
		Db: models.Db,
	}

	if err := gDao.Read(&models.GraduatesInfo{
		Id: graduate.Id,
	}); err == nil {
		//更新用户
		err = gDao.Update(graduate)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		//添加用户
		err = gDao.Insert(graduate)
		if err != nil {
			fmt.Println(err)
		}
	}
	c.Redirect("/getPageGraduates", 302)
}

func (c *GraduateController) DeleteGraduate() {
	id := c.GetString("id")
	gDao := &models.GraduatesDao{
		Db: models.Db,
	}
	err := gDao.Delete(id)
	if err != nil {
		fmt.Println(err)
	}
	c.Redirect("/getPageGraduates", 302)
}

func (c *GraduateController) IsLogin() (bool, *models.SessionValue) {
	uuid := c.Ctx.GetCookie("user")
	sessionInterface := c.GetSession(uuid)
	if sessionInterface == nil {
		return false, nil
	} else {
		return true, sessionInterface.(*models.SessionValue)
	}
}