package models

import "github.com/astaxie/beego/orm"

type UserDao struct {
	db orm.Ormer
}

func (userDao *UserDao) Insert(userInfo *UserInfo) (err error) {
	_, err = userDao.db.Insert(userInfo)
	return
}

func (userDao *UserDao) Delete(userId string) (err error) {
	_, err = userDao.db.Delete(&UserInfo{
		UserId: userId,
	})

	return
}

func (userDao *UserDao) Update(userInfo *UserInfo) (err error) {
	_, err = userDao.db.Update(userInfo)
	return
}

func (userDao *UserDao) Read(userInfo *UserInfo) (err error) {
	err = userDao.db.Read(userInfo)
	return
}
