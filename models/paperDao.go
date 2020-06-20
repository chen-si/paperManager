package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
)

type PaperDao struct{
	Db orm.Ormer
}

func (pDao *PaperDao)Insert(pInfo *PaperInfo)(err error){
	_,err = pDao.Db.Insert(pInfo)
	return
}

func (pDao *PaperDao)Update(pInfo *PaperInfo)(err error){
	_,err = pDao.Db.Update(pInfo)
	return
}

func (pDao *PaperDao)Read(pInfo *PaperInfo)(err error){
	err = pDao.Db.Read(pInfo)
	return
}

func (pDao *PaperDao)Delete(pId string)(err error){
	_,err = pDao.Db.Delete(&PaperInfo{
		PaperId: pId,
	})
	return
}

func (pDao *PaperDao)GetPagePapers(pageNo string)(page *Page,err error){
	//将PageNo转化为int64类型
	iPageNo, err := strconv.ParseInt(pageNo, 10, 64)
	if err != nil {
		return nil, err
	}
	//获取数据库中论文的总数
	var totalPapers int64
	err = pDao.Db.Raw("select count(*) from paper_info").QueryRow(&totalPapers)
	if err != nil{
		return
	}

	//设置每页显示4条记录
	var pageSize int64 = 4
	//获取总页数
	var totalPageNo int64
	if totalPapers % pageSize == 0 {
		totalPageNo = totalPapers / pageSize
	} else {
		totalPageNo = totalPapers / pageSize + 1
	}

	var papers []*PaperInfo
	_, err = pDao.Db.Raw("select paper_id,paper_name,paper_digest,paper_grade,paper_author,paper_version " +
		"from paper_info limit ?,?",(iPageNo - 1) * pageSize,pageSize).QueryRows(&papers)

	if err != nil {
		return
	}

	page = &Page{
		Papers:      papers,
		Graduates:   nil,
		Tutors:      nil,
		PageStyle:   PaperPageStyle,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalPapers,
	}
	return page, nil
}

