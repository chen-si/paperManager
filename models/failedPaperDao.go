package models

import "github.com/astaxie/beego/orm"

type FailedPaperDao struct {
	db orm.Ormer
}


func (fpDao *FailedPaperDao) Update(fp *FailedPaper)(err error){
	_,err = fpDao.db.Update(fp)
	return
}

func (fpDao *FailedPaperDao) Read(fp *FailedPaper)(err error){
	err = fpDao.db.Read(fp)
	return
}