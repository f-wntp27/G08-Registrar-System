package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

var State = map[string]bool{
	"Staff":           false,
	"Teacher":         false,
	"Student":         false,
	"ManageCourse":    false,
	"Enrollment":      false,
	"Bill":            false,
	"RequestRegister": false,
}

// Prefix data
var mister = Prefix{
	Value: "นาย",
}
var mistress = Prefix{
	Value: "นาง",
}
var miss = Prefix{
	Value: "นางสาว",
}

// Generate Password
func SetupPassword(pwd string) string {
	var password, _ = bcrypt.GenerateFromPassword([]byte(pwd), 14)
	return string(password)
}

// Staff data
var admin01 = StaffAccount{
	PrefixID:  &mister.ID,
	FirstName: "กอเอ้ย",
	LastName:  "กอไก่",
	StaffCode: "admin01",
	Password:  SetupPassword("admin"),
}
var admin02 = StaffAccount{
	PrefixID:  &miss.ID,
	FirstName: "ขอไข่",
	LastName:  "ในเล้า",
	StaffCode: "admin02",
	Password:  SetupPassword("admin"),
}

func Init_Staff() {
	// Prefix
	db.Model(&Prefix{}).Create(&mister)
	db.Model(&Prefix{}).Create(&mistress)
	db.Model(&Prefix{}).Create(&miss)

	// StaffAccount
	db.Model(&StaffAccount{}).Create(&admin01)
	db.Model(&StaffAccount{}).Create(&admin02)

	State["Staff"] = true
}

// Faculty data
var faculty_engineer = Faculty{
	Name: "วิศวกรรม",
}

// Department data
var department_cpe = Department{
	Name:      "วิศวกรรมคอมพิวเตอร์",
	FacultyID: &faculty_engineer.ID,
}

// Professor data
var teacher_AA = Professor{
	TeacherName:   "ดร.ศรัญญา กาญจนวัฒนา",
	TeacherEmail:  "AA@gmail.com",
	ProfessorCode: "R1234567",
	Password:      SetupPassword("123456"),
	DepartmentID:  &department_cpe.ID,
}
var teacher_BB = Professor{
	TeacherName:   "ดร.ชาญวิทย์ แก้วกสิ",
	TeacherEmail:  "BB@gmail.com",
	ProfessorCode: "R555552",
	Password:      SetupPassword("123456"),
	DepartmentID:  &department_cpe.ID,
}

func Init_Teacher() {
	// Faculty data
	db.Model(&Faculty{}).Create(&faculty_engineer)

	// Department data
	db.Model(&Department{}).Create(&department_cpe)

	// TeacherRecord
	db.Model(&Professor{}).Create(&teacher_AA)
	db.Model(&Professor{}).Create(&teacher_BB)

	State["Teacher"] = true
}

// StudentRecord data
var student_B6200001 = StudentRecord{
	PrefixID:     &mister.ID,
	FirstName:    "เด็กดี",
	LastName:     "มาเรียน",
	PersonalId:   "123xx688xxxxx",
	StudentCode:  "B6200001",
	Password:     SetupPassword("B6200001"),
	DepartmentID: &department_cpe.ID,
	AdvisorID:    &teacher_AA.ID,
	CreatorID:    &admin01.ID,
}
var student_B6200002 = StudentRecord{
	PrefixID:     &miss.ID,
	FirstName:    "เด็กขยัน",
	LastName:     "ใจดี",
	PersonalId:   "0000000000000",
	StudentCode:  "B6200002",
	Password:     SetupPassword("B6200002"),
	DepartmentID: &department_cpe.ID,
	AdvisorID:    &teacher_BB.ID,
	CreatorID:    &admin02.ID,
}

func Init_Student() {
	// StaffAccount
	if !State["Staff"] {
		Init_Staff()
	}

	// TeacherRecord
	if !State["Teacher"] {
		Init_Teacher()
	}

	// StudentRecord
	db.Model(&StudentRecord{}).Create(&student_B6200001)
	db.Model(&StudentRecord{}).Create(&student_B6200002)

	State["Student"] = true
}

// TA data
var ta_ta01 = TA{
	TaCode: "TA01",
	Name:   "AAaa",
}
var ta_ta02 = TA{
	TaCode: "TA02",
	Name:   "BBbb",
}

// Room data
var room_01 = Room{
	Number:       1,
	StudentCount: 300,
}
var room_02 = Room{
	Number:       2,
	StudentCount: 1500,
}

// Course data
var course_sa = Course{
	CourseCode: "523331",
	Name:       "SYSTEM ANALYSIS AND DESIGN",
	Credit:     4,
}
var course_hh = Course{
	CourseCode: "IST20 1503",
	Name:       "HOLISTIC HEALTH",
	Credit:     2,
}

// ManageCourse data
var manage_sa = ManageCourse{
	Group:            1,
	TeachingTime:     2,
	UngraduatedYear:  2,
	Trimester:        1,
	ManageCourseTime: time.Date(2021, 1, 1, 6, 0, 0, 0, time.UTC),
	CourseID:         &course_sa.ID,
	RoomID:           &room_01.ID,
	ProfessorID:      &teacher_AA.ID,
	TAID:             &ta_ta01.ID,
}
var manage_hh = ManageCourse{
	Group:            1,
	TeachingTime:     2,
	UngraduatedYear:  2,
	Trimester:        1,
	ManageCourseTime: time.Date(2021, 1, 2, 6, 0, 0, 0, time.UTC),
	CourseID:         &course_hh.ID,
	RoomID:           &room_02.ID,
	ProfessorID:      &teacher_BB.ID,
	TAID:             &ta_ta02.ID,
}

func Init_ManageCourse() {
	// Teacher
	if !State["Teacher"] {
		Init_Teacher()
	}

	// TA
	db.Model(&TA{}).Create(&ta_ta01)
	db.Model(&TA{}).Create(&ta_ta02)

	// Room
	db.Model(&Room{}).Create(&room_01)
	db.Model(&Room{}).Create(&room_02)

	// Course
	db.Model(&Course{}).Create(&course_sa)
	db.Model(&Course{}).Create(&course_hh)

	// ManageCourse
	db.Model(&ManageCourse{}).Create(&manage_sa)
	db.Model(&ManageCourse{}).Create(&manage_hh)

	State["ManageCourse"] = true
}

// EnrollmentType data
var enrollmentType_Credit = EnrollmentType{
	Name: "Credit",
}
var enrollmentType_Audit = EnrollmentType{
	Name: "Audit",
}

// Enrollment data
var enroll_01 = Enrollment{
	EnrollDateTime:  time.Date(2021, 1, 2, 6, 0, 0, 0, time.UTC),
	EnrollYear:      2020,
	EnrollTrimester: 1,
	TotalCredit:     course_sa.Credit + course_hh.Credit,
	OwnerID:         &student_B6200001.ID,
}
var enroll_02 = Enrollment{
	EnrollDateTime:  time.Date(2021, 1, 5, 6, 0, 0, 0, time.UTC),
	EnrollYear:      202,
	EnrollTrimester: 1,
	TotalCredit:     course_sa.Credit,
	OwnerID:         &student_B6200002.ID,
}

// EnrollmentItem data
var enroll_01_item_01 = EnrollmentItem{
	EnrollmentID:     &enroll_01.ID,
	ManageCourseID:   &manage_sa.ID,
	EnrollmentTypeID: &enrollmentType_Credit.ID,
}
var enroll_01_item_02 = EnrollmentItem{
	EnrollmentID:     &enroll_01.ID,
	ManageCourseID:   &manage_hh.ID,
	EnrollmentTypeID: &enrollmentType_Credit.ID,
}
var enroll_02_item_01 = EnrollmentItem{
	EnrollmentID:     &enroll_02.ID,
	ManageCourseID:   &manage_sa.ID,
	EnrollmentTypeID: &enrollmentType_Credit.ID,
}

func Init_Enrollment() {
	// StudentRecord
	if !State["Student"] {
		Init_Student()
	}

	// ManageCourse
	if !State["ManageCourse"] {
		Init_ManageCourse()
	}

	// CourseType
	db.Model(&EnrollmentType{}).Create(&enrollmentType_Credit)
	db.Model(&EnrollmentType{}).Create(&enrollmentType_Audit)

	// Enrollment
	db.Model(&Enrollment{}).Create(&enroll_01)
	db.Model(&Enrollment{}).Create(&enroll_02)

	// EnrollmentItem
	db.Model(&EnrollmentItem{}).Create(&enroll_01_item_01)
	db.Model(&EnrollmentItem{}).Create(&enroll_01_item_02)
	db.Model(&EnrollmentItem{}).Create(&enroll_02_item_01)

	State["Enrollment"] = true
}

// Place data
var place01 = Place{
	Name: "ธนาคาร A",
}
var place02 = Place{
	Name: "ธนาคาร B",
}
var place03 = Place{
	Name: "ธนาคาร C",
}

// PaymentType data
var paymentType_01 = PaymentType{
	Name: "หักผ่านบัญชี",
}
var paymentType_02 = PaymentType{
	Name: "ทุน",
}
var paymentType_03 = PaymentType{
	Name: "เงินสด",
}

// Bill data - DIDN'T SETUP

func Init_Bill() {
	// Enrollment
	if !State["Enrollment"] {
		Init_Enrollment()
	}

	// Place
	db.Model(&Place{}).Create(&place01)
	db.Model(&Place{}).Create(&place02)
	db.Model(&Place{}).Create(&place03)

	// PaymentType
	db.Model(&PaymentType{}).Create(&paymentType_01)
	db.Model(&PaymentType{}).Create(&paymentType_02)
	db.Model(&PaymentType{}).Create(&paymentType_03)

	// Bill - DIDN'T SETUP

	State["Bill"] = true
}

// RequestStatus data
var requestStatus_wait = RequestStatus{
	Name: "รอดำเนินการ",
}
var requestStatus_approve = RequestStatus{
	Name: "อนุมัติ",
}
var requestStatus_non_approve = RequestStatus{
	Name: "ไม่อนุมัติ",
}

// RequestType data
var requestType_inc = RequestType{
	Name: "เพิ่มรายวิชา",
}
var requestType_dec = RequestType{
	Name: "ลดรายวิชา",
}

// RequestRegister data
var req_B6200001_01 = RequestRegister{
	ManageCourseID:  &manage_sa.ID,
	RequestTypeID:   &requestType_inc.ID,
	RequestStatusID: &requestStatus_approve.ID,
	OwnerID:         &student_B6200001.ID,
	RequestTime:     time.Date(2021, 1, 2, 10, 0, 0, 0, time.UTC),
}
var req_B6200001_02 = RequestRegister{
	ManageCourseID:  &manage_hh.ID,
	RequestTypeID:   &requestType_inc.ID,
	RequestStatusID: &requestStatus_wait.ID,
	OwnerID:         &student_B6200001.ID,
	RequestTime:     time.Date(2021, 1, 3, 9, 0, 0, 0, time.UTC),
}
var req_B6200002_01 = RequestRegister{
	ManageCourseID:  &manage_hh.ID,
	RequestTypeID:   &requestType_inc.ID,
	RequestStatusID: &requestStatus_wait.ID,
	OwnerID:         &student_B6200002.ID,
	RequestTime:     time.Date(2021, 1, 3, 11, 0, 0, 0, time.UTC),
}

func Init_RequestRegister() {
	// StudentRecord
	if !State["Student"] {
		Init_Student()
	}

	// ManageCourse
	if !State["ManageCourse"] {
		Init_ManageCourse()
	}

	// RequestStatus
	db.Model(&RequestStatus{}).Create(&requestStatus_wait)
	db.Model(&RequestStatus{}).Create(&requestStatus_approve)
	db.Model(&RequestStatus{}).Create(&requestStatus_non_approve)

	// RequestType
	db.Model(&RequestType{}).Create(&requestType_inc)
	db.Model(&RequestType{}).Create(&requestType_dec)

	// RequestRegister
	db.Model(&RequestRegister{}).Create(&req_B6200001_01)
	db.Model(&RequestRegister{}).Create(&req_B6200001_02)
	db.Model(&RequestRegister{}).Create(&req_B6200002_01)

	State["RequestRegister"] = true
}
