package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db orm.Ormer

const (
	AdminUser  = "admin"
	NormalUser = "normal"
)

const (
	PaperPageStyle     = "PaperPage"
	GraduatesPageStyle = "GraduatesPage"
	TutorPagesStyle    = "TutorPage"
	UserPageStyle      = "UserPage"
)

type UserInfo struct {
	UserId   string `orm:"pk"`
	UserPwd  string
	UserName string
	UserRole string
}

type GraduatesInfo struct {
	Id           string `orm:"pk"`
	Name         string
	GraduateTime string
	TutorId      string
}

type TutorsInfo struct {
	TutorId   string `orm:"pk"`
	TutorName string
	College   string
}

type PaperInfo struct {
	PaperId      string `orm:"pk"`
	PaperName    string
	PaperDigest  string
	PaperGrade   int
	PaperVersion int
	PaperAuthor  string
}

type FailedPaper struct {
	PaperId         string `orm:"pk"`
	ReasonForFailed string
}

type Page struct {
	Papers      []*PaperInfo
	Graduates   []*GraduatesInfo
	Tutors      []*TutorsInfo
	Users       []*UserInfo
	PageStyle   string
	PageNo      int64 //当前页码
	PageSize    int64 //每页显示的条数
	TotalPageNo int64 //总页数 计算得到
	TotalRecord int64 //总记录数 查询数据库得到
	IsLogin     bool
	Username    string
}

type SessionValue struct {
	SessionId string
	UserId    string
	UserName  string
	UserRole  string
}

//IsHasPrev
func (p *Page) IsHasPrev() bool {
	return p.PageNo > 1
}

//IsHasNext
func (p *Page) IsHasNext() bool {
	return p.PageNo < p.TotalPageNo
}

//GetPrevPageNo
func (p *Page) GetPrevPageNo() int64 {
	if p.IsHasPrev() {
		return p.PageNo - 1
	} else {
		return 1
	}
}

//GetNextPageNo
func (p *Page) GetNextPageNo() int64 {
	if p.IsHasNext() {
		return p.PageNo + 1
	} else {
		return p.TotalPageNo
	}
}

func init() {
	orm.Debug = true

	//链接到数据库
	err := orm.RegisterDataBase("default", "mysql", "root:1234w5asd@tcp(localhost:3306)/paper_manage_system?charset=utf8", 30)
	if err != nil {
		log.Fatal(err)
		return
	}

	orm.RegisterModel(new(UserInfo), new(GraduatesInfo), new(PaperInfo), new(TutorsInfo), new(FailedPaper))
	Db = orm.NewOrm()
}
