package models

import (
	"fmt"
	"testing"
)

func TestModel(t *testing.T) {
	//FailedPaper
	t.Run("FPUpdate",testFPUpdate)
	t.Run("FPRead",testFPRead)

	//Paper
	t.Run("PaperInsert",testPaperInsert)
	t.Run("PaperUpdate",testPaperUpdate)
	t.Run("PaperDelete",testPaperDelete)
	t.Run("PaperRead",testPaperRead)
	t.Run("PaperGetPagePapers",testPaperGetPagePapers)

	//Graduate
	t.Run("GraduateInsert",testGraduateInsert)
	t.Run("GraduateUpdate",testGraduateUpdate)
	t.Run("GraduateDelete",testGraduateDelete)
	t.Run("GraduateRead",testGraduateRead)

	//Tutor
	t.Run("TutorInsert",testTutorInsert)
	t.Run("TutorUpdate",testTutorUpdate)
	t.Run("TutorDelete",testTutorDelete)
	t.Run("TutorRead",testTutorRead)

	//User
	t.Run("UserInsert", testUserDaoInsert)
	t.Run("UserDelete", testUserDaoDelete)
	t.Run("UserUpdate", testUserDaoUpdate)
	t.Run("UserRead", testUserDaoRead)
}

func testFPUpdate(t *testing.T){
	fpDao := &FailedPaperDao{
		db: db,
	}
	fp := &FailedPaper{
		PaperId:         "555666",
		ReasonForFailed: "一塌糊涂",
	}
	fmt.Println(fpDao.Update(fp))
}

func testFPRead(t *testing.T){
	fpDao := &FailedPaperDao{
		db: db,
	}
	fp := &FailedPaper{
		PaperId:         "555666",
	}
	fmt.Println(fpDao.Read(fp))
	fmt.Println(fp)
}

func testPaperInsert(t *testing.T){
	p := &PaperDao{
		Db: db,
	}

	pInfo := &PaperInfo{
		PaperId:      "456789",
		PaperName:    "论危险主义",
		PaperDigest:  "危险令人恐惧",
		PaperGrade:   5,
		PaperVersion: 1,
		PaperAuthor:  "123456",
	}
	fmt.Println(p.Insert(pInfo))
}

func testPaperUpdate(t *testing.T){
	p := &PaperDao{
		Db: db,
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

func testPaperRead(t *testing.T){
	p := &PaperDao{
		Db: db,
	}

	pInfo := &PaperInfo{
		PaperId:      "555666",
	}
	fmt.Println(p.Read(pInfo))
	fmt.Println(pInfo)
}

func testPaperDelete(t *testing.T){
	p := &PaperDao{
		Db: db,
	}
	fmt.Println(p.Delete("456789"))
}

func testPaperGetPagePapers(t *testing.T){
	p := &PaperDao{
		Db: db,
	}

	page,err := p.GetPagePapers("1")
	fmt.Println(err)
	fmt.Println(page)
	fmt.Println(page.Papers[0])
	fmt.Println(page.Papers[1])
}


func testGraduateInsert(t *testing.T){
	g := &GraduatesDao{
		db: db,
	}
	gInfo := &GraduatesInfo{
		Id:           "123456",
		Name:         "liuyang",
		GraduateTime: "2018",
		TutorId:      "111555",
	}
	fmt.Println(g.Insert(gInfo))
}

func testGraduateUpdate(t *testing.T){
	g := &GraduatesDao{
		db: db,
	}
	gInfo := &GraduatesInfo{
		Id:           "123456",
		Name:         "liuyang",
		GraduateTime: "2020",
		TutorId:      "111555",
	}
	fmt.Println(g.Update(gInfo))
}

func testGraduateRead(t *testing.T){
	g := &GraduatesDao{
		db: db,
	}
	gInfo := &GraduatesInfo{
		Id:           "123456",
	}
	fmt.Println(g.Read(gInfo))
	fmt.Println(gInfo)
}

func testGraduateDelete(t *testing.T){
	g := &GraduatesDao{
		db: db,
	}
	fmt.Println(g.Delete("123456"))
}

func testTutorDelete(t *testing.T){
	tu := &TutorDao{
		db: db,
	}
	fmt.Println(tu.Delete("111555"))
}

func testTutorUpdate(t *testing.T){
	tu := &TutorDao{
		db: db,
	}
	tutor := &TutorsInfo{
		TutorId:   "111555",
		TutorName: "liugan",
		College:   "EE",
	}
	fmt.Println(tu.Update(tutor))
}

func testTutorRead(t *testing.T){
	tu := &TutorDao{
		db: db,
	}
	tutor := &TutorsInfo{
		TutorId:   "111555",
		TutorName: "",
		College:   "",
	}
	fmt.Println(tu.Read(tutor))
	fmt.Println(tutor)
}

func testTutorInsert(t *testing.T){
	tu := &TutorDao{
		db: db,
	}
	tutor := &TutorsInfo{
		TutorId:   "111555",
		TutorName: "liugang",
		College:   "CS",
	}
	fmt.Println(tu.Insert(tutor))
}

func testUserDaoDelete(t *testing.T) {
	u := &UserDao{
		db: db,
	}
	err := u.Delete("123456")
	fmt.Println(err)
}

func testUserDaoInsert(t *testing.T) {
	u := &UserDao{
		db: db,
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
		db: db,
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
		db: db,
	}
	user := UserInfo{
		UserId: "123456",
	}
	err := u.Read(&user)
	fmt.Println(err)
	fmt.Println(user)
}
