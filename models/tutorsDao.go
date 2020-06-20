package models

import "github.com/astaxie/beego/orm"

type TutorDao struct{
	db orm.Ormer
}

func (tutorDao *TutorDao) Insert(tutor *TutorsInfo) (err error){
	_,err = tutorDao.db.Insert(tutor)
	return
}

func (tutorDao *TutorDao) Update(tutor *TutorsInfo) (err error){
	_,err = tutorDao.db.Update(tutor)
	return
}

func (tutorDao *TutorDao) Read(tutor *TutorsInfo) (err error){
	err = tutorDao.db.Read(tutor)
	return
}

func (tutorDao *TutorDao) Delete(tutorId string) (err error){
	_,err = tutorDao.db.Delete(&TutorsInfo{
		TutorId: tutorId,
	})
	return
}