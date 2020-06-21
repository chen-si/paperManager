package models

import "github.com/astaxie/beego/orm"

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
