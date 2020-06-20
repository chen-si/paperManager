package models

import "github.com/astaxie/beego/orm"

type GraduatesDao struct{
	db orm.Ormer
}

func (gDao *GraduatesDao)Insert(gInfo *GraduatesInfo)(err error){
	_,err = gDao.db.Insert(gInfo)
	return
}

func (gDao *GraduatesDao)Update(gInfo *GraduatesInfo)(err error){
	_,err = gDao.db.Update(gInfo)
	return
}

func (gDao *GraduatesDao)Read(gInfo *GraduatesInfo)(err error){
	err = gDao.db.Read(gInfo)
	return
}

func (gDao *GraduatesDao)Delete(gId string)(err error){
	_,err = gDao.db.Delete(&GraduatesInfo{
		Id: gId,
	})
	return
}