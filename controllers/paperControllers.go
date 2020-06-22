package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"os"
	"paperManger/models"
	"strconv"
)

type PaperController struct {
	beego.Controller
}

func (c *PaperController) PagePapers() {
	pDao := &models.PaperDao{
		Db: models.Db,
	}
	pageNo := c.GetString("pageNo")
	searchStr := c.GetString("searchstr")
	var page *models.Page
	var err error
	if pageNo == "" {
		page, err = pDao.GetPagePapers("1",searchStr)
	} else {
		page, err = pDao.GetPagePapers(pageNo,searchStr)
	}

	if err != nil {
		fmt.Println(err)
		return
	}
	c.TplName = "paper/page_papers.html"
	c.Data["Page"] = page
	c.Data["SearchStr"] = searchStr
}

func (c *PaperController) ToUpdateOrAddPaper() {
	pDao := &models.PaperDao{
		Db: models.Db,
	}

	paperId := c.GetString("paperid")
	paper := &models.PaperInfo{
		PaperId: paperId,
	}
	if paperId == "" {
		//add user
		c.TplName = "paper/paper_edit.html"
	} else {
		//update user
		err := pDao.Read(paper)
		if err != nil {
			fmt.Println(err)
			return
		}
		c.TplName = "paper/paper_edit.html"
		c.Data["Paper"] = paper
	}
}

func (c *PaperController) UpdateOrAddPaper() {
	paper := &models.PaperInfo{
		PaperId:     c.GetString("paperid"),
		PaperName:   c.GetString("papername"),
		PaperDigest: c.GetString("paperdigest"),
		PaperAuthor: c.GetString("paperauthor"),
	}
	paper.PaperGrade, _ = strconv.Atoi(c.GetString("papergrade"))
	paper.PaperVersion, _ = strconv.Atoi(c.GetString("paperversion"))

	pDao := &models.PaperDao{
		Db: models.Db,
	}

	if err := pDao.Read(&models.PaperInfo{
		PaperId: paper.PaperId,
	}); err == nil {
		//更新用户
		err = pDao.Update(paper)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		//添加用户
		err = pDao.Insert(paper)
		if err != nil {
			fmt.Println(err)
		}
	}
	c.Redirect("/getPagePapers", 302)
}

func (c *PaperController) DeletePaper() {
	paperId := c.GetString("paperid")
	pDao := &models.PaperDao{
		Db: models.Db,
	}
	err := pDao.Delete(paperId)
	if err != nil {
		fmt.Println(err)
	}
	c.Redirect("/getPagePapers", 302)
}

func (c *PaperController) ToUploadFile() {
	uuid := c.Ctx.GetCookie("user")
	fmt.Println(c.GetSession(uuid))
	c.TplName = "paper/submitPaper.html"
	c.Data["IsAdmin"] = false
}

func (c *PaperController) UploadFile() {
	_, session := c.IsLogin()
	paper := &models.PaperInfo{
		PaperId:     c.GetString("paperid"),
		PaperName:   c.GetString("papername"),
		PaperAuthor: c.GetString("paperauthor"),
	}

	pDao := &models.PaperDao{
		Db: models.Db,
	}
	file, header, err := c.GetFile("paperfile")
	if err != nil {
		//文件上传失败
		fmt.Println(err)
		c.TplName = "paper/submitPaper.html"
		c.Data["err"] = "有错误"
		c.Data["msg"] = "论文上传失败！"
		return
	}
	defer file.Close()
	err = c.SaveToFile("paperfile", "/home/liu/paperManagerData/"+paper.PaperId+"_"+header.Filename)
	if err != nil {
		//文件保存失败，请求重新上传
		fmt.Println(err)
		c.TplName = "paper/submitPaper.html"
		c.Data["err"] = "有错误"
		c.Data["msg"] = "论文保存失败！请重新上传"
		return
	}
	if err := pDao.Read(&models.PaperInfo{
		PaperId: paper.PaperId,
	}); err == nil {
		//论文已经存在 不需要其他操作
	} else {
		//论文不存在，插入论文信息
		err := pDao.Insert(paper)
		if err != nil {
			fmt.Println(err)
			c.TplName = "paper/submitPaper.html"
			c.Data["err"] = "有错误"
			c.Data["msg"] = "论文信息保存失败！请重新提交表单！"
			err = os.Remove("static/paperfile/" + session.UserId + "_" + header.Filename)
			if err != nil {
				//删除文件失败
				fmt.Println(err)
			}
			return
		}
	}
	c.Redirect("/main", 302)
}

func (c *PaperController) IsLogin() (bool, *models.SessionValue) {
	uuid := c.Ctx.GetCookie("user")
	sessionInterface := c.GetSession(uuid)
	if sessionInterface == nil {
		return false, nil
	} else {
		return true, sessionInterface.(*models.SessionValue)
	}
}
