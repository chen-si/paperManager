package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"os"
	"paperManger/models"
	"strconv"
)

const FileDir = "/home/liu/paperManagerData/"

type PaperController struct {
	beego.Controller
}

func (c *PaperController) ToFailedPaper() {
	flag, session := c.IsLogin()
	if !flag {
		c.Redirect("/", 302)
		return
	}

	fpDao := &models.FailedPaperDao{
		Db: models.Db,
	}

	pageNo := c.GetString("pageNo")
	var page *models.Page
	var err error
	if pageNo == "" {
		page, err = fpDao.GetFailedPaperInfoFromView("1")
	} else {
		page, err = fpDao.GetFailedPaperInfoFromView(pageNo)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	c.TplName = "paper/failed_paper.html"
	c.Data["Page"] = page
	c.sendUserType(session)
}

func (c *PaperController) ToUpdateFailedPaper() {
	flag, _ := c.IsLogin()

	if !flag {
		c.Redirect("/", 302)
		return
	}

	fpDao := &models.FailedPaperDao{
		Db: models.Db,
	}

	paperId := c.GetString("paperid")
	fp := &models.FailedPaper{
		PaperId: paperId,
	}
	//update failed paper
	err := fpDao.Read(fp)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.TplName = "paper/failed_paper_edit.html"
	c.Data["FP"] = fp
}

func (c *PaperController) UpdateFailedPaper() {
	fp := &models.FailedPaper{
		PaperId:         c.GetString("paperid"),
		ReasonForFailed: c.GetString("reason"),
		Suggestions:     c.GetString("suggestions"),
	}

	fpDao := &models.FailedPaperDao{
		Db: models.Db,
	}

	err := fpDao.Update(fp)
	if err != nil {
		fmt.Println(err)
	}
	c.Redirect("/failedPapers", 302)
}

func (c *PaperController) PagePapers() {
	flag, session := c.IsLogin()
	if !flag {
		c.Redirect("/", 302)
		return
	}
	pDao := &models.PaperDao{
		Db: models.Db,
	}
	pageNo := c.GetString("pageNo")
	searchStr := c.GetString("searchstr")
	var page *models.Page
	var err error
	if pageNo == "" {
		page, err = pDao.GetPagePapers("1", searchStr)
	} else {
		page, err = pDao.GetPagePapers(pageNo, searchStr)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	c.TplName = "paper/page_papers.html"
	c.Data["Page"] = page
	c.Data["SearchStr"] = searchStr

	c.sendUserType(session)
}

func (c *PaperController) ToUpdateOrAddPaper() {
	flag, session := c.IsLogin()

	if !flag {
		c.Redirect("/", 302)
		return
	}

	pDao := &models.PaperDao{
		Db: models.Db,
	}

	paperId := c.GetString("paperid")
	paper := &models.PaperInfo{
		PaperId: paperId,
	}
	if paperId == "" {
		//add paper
		c.TplName = "paper/paper_edit.html"
	} else {
		//update paper
		err := pDao.Read(paper)
		if err != nil {
			fmt.Println(err)
			return
		}
		if session.UserRole == models.GraduateUser && session.UserId != paper.PaperAuthor {
			//不能修改其他人的论文信息
			c.PagePapers()
			c.Data["err"] = "不能修改其他人的论文信息"
			return
		}
		c.TplName = "paper/paper_edit.html"
		c.Data["Paper"] = paper
		//发送用户类型给模板文件
		c.sendUserType(session)
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
	flag, _ := c.IsLogin()
	if !flag {
		c.Redirect("/", 302)
		return
	}
	c.TplName = "paper/submit_paper.html"
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
		c.TplName = "paper/submit_paper.html"
		c.Data["err"] = "有错误"
		c.Data["msg"] = "论文上传失败！"
		return
	}
	defer file.Close()
	err = c.SaveToFile("paperfile", FileDir+paper.PaperId+"_"+session.UserId+"_"+header.Filename)
	if err != nil {
		//文件保存失败，请求重新上传
		fmt.Println(err)
		c.TplName = "paper/submit_paper.html"
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
			c.TplName = "paper/submit_paper.html"
			c.Data["err"] = "有错误"
			c.Data["msg"] = "论文信息保存失败！请重新提交表单！"
			err = os.Remove(FileDir + paper.PaperId + "_" + session.UserId + "_" + header.Filename)
			if err != nil {
				//删除文件失败
				fmt.Println(err)
			}
			return
		}
	}
	err = pDao.UpdatePaperFilePath(paper.PaperId, FileDir+paper.PaperId+"_"+session.UserId+"_"+header.Filename)
	if err != nil {
		fmt.Println(err)
		c.TplName = "paper/submit_paper.html"
		c.Data["err"] = "有错误"
		c.Data["msg"] = "论文路径保存失败！请重新提交表单！"
		err = os.Remove(FileDir + paper.PaperId + "_" + header.Filename)
		if err != nil {
			//删除文件失败
			fmt.Println(err)
		}
		return
	}
	c.Redirect("/main", 302)
}

func (c *PaperController) DownloadFile() {
	paperId := c.GetString("paperid")

	pDao := &models.PaperDao{
		Db: models.Db,
	}

	filepath, err := pDao.GetPaperFilePathById(paperId)

	if err != nil {
		fmt.Println(err)
		c.Redirect("/getPagePapers", 302)
		return
	}
	c.Ctx.Output.Download(filepath)
	c.Redirect("/getPagePapers", 302)
	return
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

func (c *PaperController) sendUserType(value *models.SessionValue) {
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
