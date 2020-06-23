package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
)

type FailedPaperDao struct {
	Db orm.Ormer
}

func (fpDao *FailedPaperDao) Update(fp *FailedPaper) (err error) {
	_, err = fpDao.Db.Update(fp)
	return
}

func (fpDao *FailedPaperDao) Read(fp *FailedPaper) (err error) {
	err = fpDao.Db.Read(fp)
	return
}

func (fpDao *FailedPaperDao) GetFailedPaperInfoFromView(pageNo string)(page *Page,err error){
	//将PageNo转化为int64类型
	iPageNo, err := strconv.ParseInt(pageNo, 10, 64)
	if err != nil {
		return nil, err
	}
	//获取数据库中未通过论文的总数
	var totalPapers int64
	err = fpDao.Db.Raw("select count(*) from gra_failed_paper_info").QueryRow(&totalPapers)
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

	var fpInfo []*FailedPaperInfo
	_, err = fpDao.Db.Raw("select name,graduate_time,paper_name,paper_id,digest,reason,suggestions "+
		"from gra_failed_paper_info limit ?,?", (iPageNo-1)*pageSize, pageSize).QueryRows(&fpInfo)

	if err != nil {
		return
	}

	page = &Page{
		FailedPapers: fpInfo,
		PageStyle:   FailedPaperPageStyle,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalPapers,
	}
	return page, nil
}
