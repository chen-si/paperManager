package models

import (
	"fmt"
	"testing"
)

func TestModel(t *testing.T) {
	//FailedPaper
	t.Run("FPUpdate", testFPUpdate)
	t.Run("FPRead", testFPRead)
	t.Run("FPPagePaper",testFPPagePaper)

	//Paper
	t.Run("PaperInsert", testPaperInsert)
	t.Run("PaperUpdate", testPaperUpdate)
	t.Run("PaperDelete", testPaperDelete)
	t.Run("PaperRead", testPaperRead)
	t.Run("PaperGetPagePapers", testPaperGetPagePapers)
	t.Run("UpdatePaperFilePath",testPaperUpdateFilePath)
	t.Run("GetPaperFilePath",testPaperGetPaperFilePath)

	//Graduate
	t.Run("GraduateInsert", testGraduateInsert)
	t.Run("GraduateUpdate", testGraduateUpdate)
	t.Run("GraduateDelete", testGraduateDelete)
	t.Run("GraduateRead", testGraduateRead)
	t.Run("GetGraPaperInfo",testGetGraPaperInfo)

	//Tutor
	t.Run("TutorInsert", testTutorInsert)
	t.Run("TutorUpdate", testTutorUpdate)
	t.Run("TutorDelete", testTutorDelete)
	t.Run("TutorRead", testTutorRead)
	t.Run("GetTutorGraInfo",testGetTutorGraInfo)
	t.Run("GetTutorGraPaperInfo",testGetTutorGraPaperInfo)

	//User
	t.Run("UserInsert", testUserDaoInsert)
	t.Run("UserDelete", testUserDaoDelete)
	t.Run("UserUpdate", testUserDaoUpdate)
	t.Run("UserRead", testUserDaoRead)
	t.Run("GetPageUsers", testGetPageUsers)
}

//Failed Paper

func testFPUpdate(t *testing.T) {
	fpDao := &FailedPaperDao{
		Db: Db,
	}
	fp := &FailedPaper{
		PaperId:         "555666",
		ReasonForFailed: "一塌糊涂",
	}
	fmt.Println(fpDao.Update(fp))
}

func testFPRead(t *testing.T) {
	fpDao := &FailedPaperDao{
		Db: Db,
	}
	fp := &FailedPaper{
		PaperId: "555666",
	}
	fmt.Println(fpDao.Read(fp))
	fmt.Println(fp)
}

func testFPPagePaper(t *testing.T){
	fpDao := &FailedPaperDao{
		Db: Db,
	}

	page,err := fpDao.GetFailedPaperInfoFromView("1")
	fmt.Println(err)
	fmt.Println(page)
	fmt.Println(page.FailedPapers[0])
	fmt.Println(page.FailedPapers[1])
}

//Paper

func testPaperInsert(t *testing.T) {
	p := &PaperDao{
		Db: Db,
	}

	pInfo := &PaperInfo{
		PaperId:      "555666",
		PaperName:    "论自由主义",
		PaperDigest:  "自由令人无限向往",
		PaperGrade:   5,
		PaperVersion: 1,
		PaperAuthor:  "123456",
	}
	fmt.Println(p.Insert(pInfo))
}

func testPaperUpdate(t *testing.T) {
	p := &PaperDao{
		Db: Db,
	}

	pInfo := &PaperInfo{
		PaperId:      "456789",
		PaperName:    "论危险主义",
		PaperDigest:  "危险令人恐惧",
		PaperGrade:   5,
		PaperVersion: 3,
		PaperAuthor:  "123456",
	}
	fmt.Println(p.Update(pInfo))
}

func testPaperRead(t *testing.T) {
	p := &PaperDao{
		Db: Db,
	}

	pInfo := &PaperInfo{
		PaperId: "555666",
	}
	fmt.Println(p.Read(pInfo))
	fmt.Println(pInfo)
}

func testPaperDelete(t *testing.T) {
	p := &PaperDao{
		Db: Db,
	}
	fmt.Println(p.Delete("555666"))
}

func testPaperGetPagePapers(t *testing.T) {
	p := &PaperDao{
		Db: Db,
	}

	page, err := p.GetPagePapers("1","")
	fmt.Println(err)
	fmt.Println(page)
	fmt.Println(page.Papers[0])
	fmt.Println(page.Papers[1])
}

func testPaperUpdateFilePath(t *testing.T){
	p := &PaperDao{
		Db: Db,
	}
	fmt.Println(p.UpdatePaperFilePath("555666","test.txt"))
}

func testPaperGetPaperFilePath(t *testing.T){
	p := &PaperDao{
		Db: Db,
	}
	fmt.Println(p.GetPaperFilePathById("555666"))
}


//Graduate

func testGetGraPaperInfo(t *testing.T){
	gDao := &GraduatesDao{
		Db: Db,
	}
	gpInfo,err := gDao.GetGraPaperInfo("101510")
	fmt.Println(err)
	for _,v := range gpInfo{
		fmt.Println(v)
	}
}

func testGraduateInsert(t *testing.T) {
	g := &GraduatesDao{
		Db: Db,
	}
	gInfo := &GraduatesInfo{
		Id:           "123456",
		Name:         "liuyang",
		GraduateTime: "2018",
		TutorId:      "111555",
	}
	fmt.Println(g.Insert(gInfo))
}

func testGraduateUpdate(t *testing.T) {
	g := &GraduatesDao{
		Db: Db,
	}
	gInfo := &GraduatesInfo{
		Id:           "123456",
		Name:         "liuyang",
		GraduateTime: "2020",
		TutorId:      "111555",
	}
	fmt.Println(g.Update(gInfo))
}

func testGraduateRead(t *testing.T) {
	g := &GraduatesDao{
		Db: Db,
	}
	gInfo := &GraduatesInfo{
		Id: "123456",
	}
	fmt.Println(g.Read(gInfo))
	fmt.Println(gInfo)
}

func testGraduateDelete(t *testing.T) {
	g := &GraduatesDao{
		Db: Db,
	}
	fmt.Println(g.Delete("123456"))
}

//Tutor

func testGetTutorGraInfo(t *testing.T){
	tDao := &TutorDao{
		Db: Db,
	}
	tgInfo,err := tDao.GetTutorGraInfo("111555")
	fmt.Println(err)
	for _,v := range tgInfo{
		fmt.Println(v)
	}
}

func testGetTutorGraPaperInfo(t *testing.T){
	tDao := &TutorDao{
		Db: Db,
	}
	tgpInfo,err := tDao.GetTutorGraPaperInfo("111555")
	fmt.Println(err)
	for _,v := range tgpInfo{
		fmt.Println(v)
	}
}

func testTutorDelete(t *testing.T) {
	tu := &TutorDao{
		Db: Db,
	}
	fmt.Println(tu.Delete("111555"))
}

func testTutorUpdate(t *testing.T) {
	tu := &TutorDao{
		Db: Db,
	}
	tutor := &TutorsInfo{
		TutorId:   "111555",
		TutorName: "liugan",
		College:   "EE",
	}
	fmt.Println(tu.Update(tutor))
}

func testTutorRead(t *testing.T) {
	tu := &TutorDao{
		Db: Db,
	}
	tutor := &TutorsInfo{
		TutorId:   "111555",
		TutorName: "",
		College:   "",
	}
	fmt.Println(tu.Read(tutor))
	fmt.Println(tutor)
}

func testTutorInsert(t *testing.T) {
	tu := &TutorDao{
		Db: Db,
	}
	tutor := &TutorsInfo{
		TutorId:   "111555",
		TutorName: "liugang",
		College:   "CS",
	}
	fmt.Println(tu.Insert(tutor))
}

//User

func testUserDaoDelete(t *testing.T) {
	u := &UserDao{
		Db: Db,
	}
	err := u.Delete("123456")
	fmt.Println(err)
}

func testUserDaoInsert(t *testing.T) {
	u := &UserDao{
		Db: Db,
	}
	user := UserInfo{
		UserId:   "123456",
		UserPwd:  "123456",
		UserName: "liu",
		UserRole: "admin",
	}
	err := u.Insert(&user)
	fmt.Println(err)
}

func testUserDaoUpdate(t *testing.T) {
	u := &UserDao{
		Db: Db,
	}
	user := UserInfo{
		UserId:   "123456",
		UserPwd:  "123456",
		UserName: "yang",
		UserRole: "normal",
	}
	err := u.Update(&user)
	fmt.Println(err)
}

func testUserDaoRead(t *testing.T) {
	u := &UserDao{
		Db: Db,
	}
	user := UserInfo{
		UserId: "123456",
	}
	err := u.Read(&user)
	fmt.Println(err)
	fmt.Println(user)
}

func testGetPageUsers(t *testing.T) {
	u := &UserDao{
		Db: Db,
	}
	page, err := u.GetPageUsers("1")
	fmt.Println(err)
	fmt.Println(page)
	fmt.Println(page.Users)
	fmt.Println(page.Users[0])
}
