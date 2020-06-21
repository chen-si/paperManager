package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"paperManger/models"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "user/login.html"
}

func (c *LoginController) Post() {
	userId := c.GetString("username")
	userPwd := c.GetString("password")
	uDao := &models.UserDao{
		Db: models.Db,
	}

	user := models.UserInfo{
		UserId: userId,
	}
	//从数据库中读取对应ID的用户信息
	_ = uDao.Read(&user)
	//if err != nil{
	//	log.Fatal(err)
	//}
	if user.UserPwd == userPwd {
		//登录验证成功
		c.TplName = "user/login_success.html"
		c.Data["name"] = user.UserName
	} else {
		//用户名或密码不正确！
		c.TplName = "user/login.html"
		c.Data["err"] = "用户名或密码不正确！"
	}
}

type UserManagerController struct {
	beego.Controller
}

func (c *UserManagerController) Get() {
	uDao := &models.UserDao{
		Db: models.Db,
	}
	pageNo := c.GetString("pageNo")
	var page *models.Page
	var err error
	if pageNo == "" {
		page, err = uDao.GetPageUsers("1")
	} else {
		page, err = uDao.GetPageUsers(pageNo)
	}

	if err != nil {
		fmt.Println(err)
		return
	}
	c.TplName = "user/page_users.html"
	c.Data["Page"] = page
}

type DeleteUserController struct {
	beego.Controller
}

func (c *DeleteUserController) Get() {
	userId := c.GetString("userid")
	uDao := &models.UserDao{
		Db: models.Db,
	}
	err := uDao.Delete(userId)
	if err != nil{
		fmt.Println(err)
	}
	c.Redirect("/getPageUsers",302)
}

type UpdateOrAddUserController struct{
	beego.Controller
}

func (c *UpdateOrAddUserController) Get(){
	uDao := &models.UserDao{
		Db: models.Db,
	}
	
	userId := c.GetString("userid")
	user := &models.UserInfo{
		UserId: userId,
	}
	if userId == ""{
		//add user
		c.TplName = "user/user_edit.html"
	}else{
		//update user
		err := uDao.Read(user)
		fmt.Println(user)
		if err != nil{
			fmt.Println(err)
			return
		}
		c.TplName = "user/user_edit.html"
		c.Data["User"] = user
	}
}

func (c *UpdateOrAddUserController) Post(){
	user := &models.UserInfo{
		UserId:   c.GetString("userid"),
		UserPwd:  c.GetString("username"),
		UserName: c.GetString("userpwd"),
		UserRole: c.GetString("userrole"),
	}
	
	uDao := &models.UserDao{
		Db: models.Db,
	}
	
	if err := uDao.Read(&models.UserInfo{
		UserId: user.UserId,
	}); err == nil{
		//更新用户
		err = uDao.Update(user)
		if err != nil{
			fmt.Println(err)
		}
	}else{
		//添加用户
		err = uDao.Insert(user)
		if err != nil{
			fmt.Println(err)
		}
	}
	c.Redirect("/getPageUsers",302)
}