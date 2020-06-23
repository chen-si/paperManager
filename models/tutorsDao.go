package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
)

type TutorDao struct {
	Db orm.Ormer
}

func (tDao *TutorDao) Insert(tutor *TutorsInfo) (err error) {
	_, err = tDao.Db.Insert(tutor)
	return
}

func (tDao *TutorDao) Update(tutor *TutorsInfo) (err error) {
	_, err = tDao.Db.Update(tutor)
	return
}

func (tDao *TutorDao) Read(tutor *TutorsInfo) (err error) {
	err = tDao.Db.Read(tutor)
	return
}

func (tDao *TutorDao) Delete(tutorId string) (err error) {
	_, err = tDao.Db.Delete(&TutorsInfo{
		TutorId: tutorId,
	})
	return
}

func (tDao *TutorDao) GetPageTutors(pageNo string) (page *Page, err error) {
	//将PageNo转化为int64类型
	iPageNo, err := strconv.ParseInt(pageNo, 10, 64)
	if err != nil {
		return nil, err
	}
	//获取数据库中论文的总数
	var totalPapers int64
	err = tDao.Db.Raw("select count(*) from tutors_info").QueryRow(&totalPapers)
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

	var tutors []*TutorsInfo
	_, err = tDao.Db.Raw("select tutor_id,tutor_name,college "+
		"from tutors_info limit ?,?", (iPageNo-1)*pageSize, pageSize).QueryRows(&tutors)

	if err != nil {
		return
	}

	page = &Page{
		Papers:      nil,
		Graduates:   nil,
		Tutors:      tutors,
		PageStyle:   TutorPagesStyle,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalPapers,
	}
	return page, nil
}

func (tDao *TutorDao) GetTutorGraPaperInfo(id string)([]*TutorGraPaperInfo,error){
	var tgpInfo []*TutorGraPaperInfo
	_,err := tDao.Db.Raw("select tutor_id, tutor_name, college, gra_id as id, graduate_name as name, " +
		" graduate_time, paper_name, paper_digest " +
	" from tutor_gra_paper where tutor_id = ?",id).QueryRows(&tgpInfo)
	return tgpInfo,err
}

func (tDao *TutorDao) GetTutorGraInfo(id string)([]*TutorGraInfo,error){
	var tgInfo []*TutorGraInfo
	_,err := tDao.Db.Raw("select ti.tutor_id ,ti.tutor_name, ti.college,gi.id , gi.name ,gi.graduate_time " +
	" from tutors_info ti join graduates_info gi on ti.tutor_id = gi.tutor_id where ti.tutor_id = ?",id).QueryRows(&tgInfo)
	return tgInfo,err
}