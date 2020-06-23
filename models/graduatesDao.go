package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
)

type GraduatesDao struct {
	Db orm.Ormer
}

func (gDao *GraduatesDao) Insert(gInfo *GraduatesInfo) (err error) {
	_, err = gDao.Db.Insert(gInfo)
	return
}

func (gDao *GraduatesDao) Update(gInfo *GraduatesInfo) (err error) {
	_, err = gDao.Db.Update(gInfo)
	return
}

func (gDao *GraduatesDao) Read(gInfo *GraduatesInfo) (err error) {
	err = gDao.Db.Read(gInfo)
	return
}

func (gDao *GraduatesDao) Delete(gId string) (err error) {
	_, err = gDao.Db.Delete(&GraduatesInfo{
		Id: gId,
	})
	return
}

func (gDao *GraduatesDao) GetPageGraduates(pageNo string) (page *Page, err error) {
	//将PageNo转化为int64类型
	iPageNo, err := strconv.ParseInt(pageNo, 10, 64)
	if err != nil {
		return nil, err
	}
	//获取数据库中论文的总数
	var totalPapers int64
	err = gDao.Db.Raw("select count(*) from graduates_info").QueryRow(&totalPapers)
	if err != nil {
		return
	}

	//设置每页显示4条记录
	var pageSize int64 = 4
	//获取总页数
	var totalPageNo int64
	if totalPapers%pageSize == 0 {
		totalPageNo = totalPapers / pageSize
	} else {
		totalPageNo = totalPapers/pageSize + 1
	}

	var graduates []*GraduatesInfo
	_, err = gDao.Db.Raw("select id,name,graduate_time,tutor_id "+
		"from graduates_info limit ?,?", (iPageNo-1)*pageSize, pageSize).QueryRows(&graduates)

	if err != nil {
		return
	}

	page = &Page{
		Papers:      nil,
		Graduates:   graduates,
		Tutors:      nil,
		PageStyle:   GraduatesPageStyle,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalPapers,
	}
	return page, nil
}

func (gDao *GraduatesDao) GetGraPaperInfo(id string)([]*GraPaperInfo,error){
	var gpInfo []*GraPaperInfo
	_,err := gDao.Db.Raw("select gi.id id,gi.name name,gi.graduate_time graduate_time,gi.tutor_id tutor_id," +
		" pi.paper_id paper_id,pi.paper_name paper_name, pi.paper_grade paper_grade,pi.paper_digest paper_digest " +
		"from graduates_info gi join paper_info pi on gi.id = pi.paper_author where gi.id = ?",id).QueryRows(&gpInfo)
	return gpInfo,err
}
