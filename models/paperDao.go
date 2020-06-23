package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
)

type PaperDao struct {
	Db orm.Ormer
}

func (pDao *PaperDao) Insert(pInfo *PaperInfo) (err error) {
	_, err = pDao.Db.Insert(pInfo)
	return
}

func (pDao *PaperDao) Update(pInfo *PaperInfo) (err error) {
	_, err = pDao.Db.Update(pInfo)
	return
}

func (pDao *PaperDao) Read(pInfo *PaperInfo) (err error) {
	err = pDao.Db.Read(pInfo)
	return
}

func (pDao *PaperDao) Delete(pId string) (err error) {
	_, err = pDao.Db.Delete(&PaperInfo{
		PaperId: pId,
	})
	return
}

func (pDao *PaperDao) GetPagePapers(pageNo , search string) (page *Page, err error) {
	//将PageNo转化为int64类型
	iPageNo, err := strconv.ParseInt(pageNo, 10, 64)
	if err != nil {
		return nil, err
	}
	//获取数据库中论文的总数
	var totalPapers int64
	switch typeOfSearch(search){
	case "":
		err = pDao.Db.Raw("select count(*) from paper_info").QueryRow(&totalPapers)
	case "author":
		err = pDao.Db.Raw("select count(*) from paper_info where paper_author = ?",search).QueryRow(&totalPapers)
	case "name":
		err = pDao.Db.Raw("select count(*) from paper_info where paper_name like concat('%',?,'%')",search).QueryRow(&totalPapers)
	}
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

	var papers []*PaperInfo
	switch typeOfSearch(search){
	case "":
		_, err = pDao.Db.Raw("select paper_id,paper_name,paper_digest,paper_grade,paper_author,paper_version,paper_filepath "+
			"from paper_info limit ?,?", (iPageNo-1)*pageSize, pageSize).QueryRows(&papers)
	case "author":
		_, err = pDao.Db.Raw("select paper_id,paper_name,paper_digest,paper_grade,paper_author,paper_version,paper_filepath "+
			"from paper_info where paper_author = ? limit ?,?", search , (iPageNo-1)*pageSize, pageSize).QueryRows(&papers)
	case "name":
		_, err = pDao.Db.Raw("select paper_id,paper_name,paper_digest,paper_grade,paper_author,paper_version,paper_filepath "+
			"from paper_info where paper_name like concat('%',?,'%') limit ?,?", search , (iPageNo-1)*pageSize, pageSize).QueryRows(&papers)
	}


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

func (pDao *PaperDao) UpdatePaperFilePath(paperId string,paperFilePath string)(err error){
	_,err = pDao.Db.Raw("update paper_info set paper_filepath = ? where paper_id = ?",paperFilePath,paperId).Exec()
	return err
}

func (pDao *PaperDao) GetPaperFilePathById(paperId string)(filepath string,err error){
	err = pDao.Db.Raw("select paper_filepath from paper_info where paper_id = ?",paperId).QueryRow(&filepath)
	return
}

func typeOfSearch(search string) string{
	if search == ""{
		return ""
	}else if isAuthorId(search){
		return "author"
	}else{
		return "name"
	}
}

func isAuthorId(str string) bool{
	_,err := strconv.Atoi(str)
	return err == nil
}