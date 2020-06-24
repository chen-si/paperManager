package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"paperManger/models"
)

type TutorController struct {
	beego.Controller
}

func (c *TutorController) PageTutors() {
	flag, session := c.IsLogin()
	if !flag {
		//未登录状态
		c.Redirect("/", 302)
		return
	}
	tDao := &models.TutorDao{
		Db: models.Db,
	}
	pageNo := c.GetString("pageNo")
	var page *models.Page
	var err error
	if pageNo == "" {
		page, err = tDao.GetPageTutors("1")
	} else {
		page, err = tDao.GetPageTutors(pageNo)
	}

	if err != nil {
		fmt.Println(err)
		return
	}
	c.TplName = "tutor/page_tutors.html"
	c.Data["Page"] = page
	c.sendUserType(session)
}

func (c *TutorController) ToUpdateOrAddTutor() {
	//判断登录状态
	flag, _ := c.IsLogin()
	if !flag {
		c.Redirect("/", 302)
		return
	}

	tDao := &models.TutorDao{
		Db: models.Db,
	}

	tutorId := c.GetString("tutorid")
	tutor := &models.TutorsInfo{
		TutorId: tutorId,
	}
	if tutorId == "" {
		//add user
		c.TplName = "tutor/tutor_edit.html"
	} else {
		//update user
		err := tDao.Read(tutor)
		if err != nil {
			fmt.Println(err)
			return
		}
		c.TplName = "tutor/tutor_edit.html"
		c.Data["Tutor"] = tutor
	}
}

func (c *TutorController) UpdateOrAddTutor() {
	tutor := &models.TutorsInfo{
		TutorId:   c.GetString("tutorid"),
		TutorName: c.GetString("tutorname"),
		College:   c.GetString("college"),
	}

	tDao := &models.TutorDao{
		Db: models.Db,
	}

	if err := tDao.Read(&models.TutorsInfo{
		TutorId: tutor.TutorId,
	}); err == nil {
		//更新用户
		err = tDao.Update(tutor)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		//添加用户
		err = tDao.Insert(tutor)
		if err != nil {
			fmt.Println(err)
		}
	}
	c.Redirect("/getPageTutors", 302)
}

func (c *TutorController) DeleteTutor() {
	tutorId := c.GetString("tutorid")
	tDao := &models.TutorDao{
		Db: models.Db,
	}
	err := tDao.Delete(tutorId)
	if err != nil {
		fmt.Println(err)
	}
	c.Redirect("/getPageTutors", 302)
}

func (c *TutorController) GetTutorGraInfo() {
	//判断登录状态
	flag, _ := c.IsLogin()
	if !flag {
		c.Redirect("/", 302)
		return
	}
	id := c.GetString("tutorid")

	tDao := &models.TutorDao{
		Db: models.Db,
	}

	tgInfo, err := tDao.GetTutorGraInfo(id)
	if err != nil {
		fmt.Println(err)
		c.Redirect("/", 302)
		return
	}
	c.TplName = "tutor/tutor_gra_info.html"
	c.Data["TutorGraInfos"] = tgInfo
}

func (c *TutorController) GetTutorGraPaperInfo() {
	//判断登录状态
	flag, _ := c.IsLogin()
	if !flag {
		c.Redirect("/", 302)
		return
	}
	id := c.GetString("tutorid")

	tDao := &models.TutorDao{
		Db: models.Db,
	}

	tgpInfo, err := tDao.GetTutorGraPaperInfo(id)
	if err != nil {
		fmt.Println(err)
		c.Redirect("/", 302)
		return
	}
	c.TplName = "tutor/tutor_gra_paper_info.html"
	c.Data["TutorGraPaperInfos"] = tgpInfo
}

func (c *TutorController) IsLogin() (bool, *models.SessionValue) {
	uuid := c.Ctx.GetCookie("user")
	sessionInterface := c.GetSession(uuid)
	if sessionInterface == nil {
		return false, nil
	} else {
		return true, sessionInterface.(*models.SessionValue)
	}
}

func (c *TutorController) sendUserType(value *models.SessionValue) {
	switch value.UserRole {
	case models.AdminUser:
		c.Data["IsAdmin"] = true
		c.Data["IsNormal"] = false
		c.Data["IsGraduate"] = false
		c.Data["IsTutor"] = false
	case models.NormalUser:
		c.Data["IsAdmin"] = false
		c.Data["IsNormal"] = true
		c.Data["IsGraduate"] = false
		c.Data["IsTutor"] = false
	case models.GraduateUser:
		c.Data["IsAdmin"] = false
		c.Data["IsNormal"] = false
		c.Data["IsGraduate"] = true
		c.Data["IsTutor"] = false
	case models.TutorUser:
		c.Data["IsAdmin"] = false
		c.Data["IsNormal"] = false
		c.Data["IsGraduate"] = false
		c.Data["IsTutor"] = true
	}
}
