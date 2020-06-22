package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"paperManger/models"
	"paperManger/utils"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) ToLogin() {
	flag, _ := c.IsLogin()
	if flag {
		//已经登录了
		c.Redirect("/main", 302)
	} else {
		c.TplName = "user/login.html"
	}
}

func (c *UserController) Login() {
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
		uuid := utils.CreateUUID()
		session := &models.SessionValue{
			SessionId: uuid,
			UserId:    user.UserId,
			UserName:  user.UserName,
			UserRole:  user.UserRole,
		}
		c.SetSession(uuid, session)
		c.Ctx.SetCookie("user", uuid, 36000, "/")
		c.Redirect("/main", 302)
	} else {
		//用户名或密码不正确！
		c.TplName = "user/login.html"
		c.Data["err"] = "用户名或密码不正确！"
	}
}

func (c *UserController) Logout() {
	uuid := c.Ctx.GetCookie("user")
	c.DelSession(uuid)
	c.Redirect("/", 302)
}

func (c *UserController) PageUsers() {
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

func (c *UserController) DeleteUser() {
	userId := c.GetString("userid")
	uDao := &models.UserDao{
		Db: models.Db,
	}
	err := uDao.Delete(userId)
	if err != nil {
		fmt.Println(err)
	}
	c.Redirect("/getPageUsers", 302)
}

func (c *UserController) ToUpdateOrAddUser() {
	uDao := &models.UserDao{
		Db: models.Db,
	}

	userId := c.GetString("userid")
	user := &models.UserInfo{
		UserId: userId,
	}
	if userId == "" {
		//add user
		c.TplName = "user/user_edit.html"
	} else {
		//update user
		err := uDao.Read(user)
		fmt.Println(user)
		if err != nil {
			fmt.Println(err)
			return
		}
		c.TplName = "user/user_edit.html"
		c.Data["User"] = user
	}
}

func (c *UserController) UpdateOrAddUser() {
	user := &models.UserInfo{
		UserId:   c.GetString("userid"),
		UserPwd:  c.GetString("userpwd"),
		UserName: c.GetString("username"),
		UserRole: c.GetString("userrole"),
	}

	uDao := &models.UserDao{
		Db: models.Db,
	}

	if err := uDao.Read(&models.UserInfo{
		UserId: user.UserId,
	}); err == nil {
		//更新用户
		err = uDao.Update(user)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		//添加用户
		err = uDao.Insert(user)
		if err != nil {
			fmt.Println(err)
		}
	}
	c.Redirect("/getPageUsers", 302)
}

func (c *UserController) IsLogin() (bool, *models.SessionValue) {
	uuid := c.Ctx.GetCookie("user")
	sessionInterface := c.GetSession(uuid)
	if sessionInterface == nil {
		return false, nil
	} else {
		return true, sessionInterface.(*models.SessionValue)
	}
}

func (c *UserController) MainPage() {
	flag, session := c.IsLogin()
	if !flag {
		c.Redirect("/", 302)
		return
	}
	c.TplName = "main_page.html"
	c.Data["Name"] = session.UserName
	c.Data["Role"] = session.UserRole
	c.Data["Id"] = session.UserId
}
