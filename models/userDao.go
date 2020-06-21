package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
)

type UserDao struct {
	Db orm.Ormer
}

func (uDao *UserDao) Insert(userInfo *UserInfo) (err error) {
	_, err = uDao.Db.Insert(userInfo)
	return
}

func (uDao *UserDao) Delete(userId string) (err error) {
	_, err = uDao.Db.Delete(&UserInfo{
		UserId: userId,
	})

	return
}

func (uDao *UserDao) Update(userInfo *UserInfo) (err error) {
	_, err = uDao.Db.Update(userInfo)
	return
}

func (uDao *UserDao) Read(userInfo *UserInfo) (err error) {
	err = uDao.Db.Read(userInfo)
	return
}

func (uDao *UserDao) GetPageUsers(pageNo string) (page *Page, err error) {
	//将PageNo转化为int64类型
	iPageNo, err := strconv.ParseInt(pageNo, 10, 64)
	if err != nil {
		return nil, err
	}
	//获取数据库中论文的总数
	var totalPapers int64
	err = uDao.Db.Raw("select count(*) from user_info").QueryRow(&totalPapers)
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

	var users []*UserInfo
	_, err = uDao.Db.Raw("select `user_id`,`user_name`,`user_pwd`,`user_role` "+
		"from user_info limit ?,?", (iPageNo-1)*pageSize, pageSize).QueryRows(&users)

	if err != nil {
		return
	}

	page = &Page{
		Papers:      nil,
		Graduates:   nil,
		Tutors:      nil,
		Users:       users,
		PageStyle:   UserPageStyle,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalPapers,
	}
	return page, nil
}
