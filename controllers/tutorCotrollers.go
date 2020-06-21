package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"paperManger/models"
)

type TutorController struct{
	beego.Controller
}

func (c *TutorController) PageTutors() {
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
}

func (c *TutorController) ToUpdateOrAddTutor() {
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