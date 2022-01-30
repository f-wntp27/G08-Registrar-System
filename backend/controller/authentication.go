package controller

import (
	"net/http"

	"sa-project-g08/backend/entity"
	"sa-project-g08/backend/service"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginPayload struct {
	UserCode string `json:"userCode"`
	Password string `json:"password"`
}

/* --- Student Response --- */
type StudentResponse struct {
	ID          uint
	Prefix      entity.Prefix
	FirstName   string
	LastName    string
	StudentCode string
	Department  entity.Department
	Advisor     ProfessorResponse
}

type LoginStudentResponse struct {
	Token   string          `json:"token"`
	Student StudentResponse `json:"student"`
}

// POST /student/login
func LoginStudent(c *gin.Context) {
	var payload LoginPayload
	var student entity.StudentRecord

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Preload("Prefix").Preload("Department").Preload("Advisor").
		Raw("SELECT * FROM student_records WHERE student_code = ?", payload.UserCode).Find(&student); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "student not found"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(student.Password), []byte(payload.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user credentials"})
		return
	}

	// กำหนดค่า SecretKey, Issuer และระยะเวลาหมดอายุของ Token สามารถกำหนดเองได้
	// SecretKey ใช้สำหรับการ sign ข้อความเพื่อบอกว่าข้อความมาจากตัวเราแน่นอน
	// Issuer เป็น unique id ที่เอาไว้ระบุตัว client
	// ExpirationHours เป็นเวลาหมดอายุของ token

	jwtWrapper := service.JwtWrapper{
		SecretKey:       "SvNQpBN8y3qlVrsGAYYWoJJk56LtzFHx",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}

	signedToken, err := jwtWrapper.GenerateToken(student.StudentCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error signing token"})
		return
	}

	tokenResponse := LoginStudentResponse{
		Token: signedToken,
		Student: StudentResponse{
			ID:          student.ID,
			Prefix:      student.Prefix,
			FirstName:   student.FirstName,
			LastName:    student.LastName,
			StudentCode: student.StudentCode,
			Department:  student.Department,
			Advisor: ProfessorResponse{
				ID:            student.Advisor.ID,
				TeacherName:   student.Advisor.TeacherName,
				TeacherEmail:  student.Advisor.TeacherEmail,
				ProfessorCode: student.Advisor.ProfessorCode,
			},
		},
	}

	c.JSON(http.StatusOK, gin.H{"data": tokenResponse})
}

/* --- Staff Response --- */
type StaffResponse struct {
	ID        uint
	Prefix    entity.Prefix
	FirstName string
	LastName  string
	StaffCode string
}

type LoginStaffResponse struct {
	Token string        `json:"token"`
	Staff StaffResponse `json:"staff"`
}

// POST /staff/login
func LoginStaff(c *gin.Context) {
	var payload LoginPayload
	var staff entity.StaffAccount

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Preload("Prefix").
		Raw("SELECT * FROM staff_accounts WHERE staff_code = ?", payload.UserCode).Find(&staff); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "staff not found"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(staff.Password), []byte(payload.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user credentials"})
		return
	}

	// กำหนดค่า SecretKey, Issuer และระยะเวลาหมดอายุของ Token สามารถกำหนดเองได้
	// SecretKey ใช้สำหรับการ sign ข้อความเพื่อบอกว่าข้อความมาจากตัวเราแน่นอน
	// Issuer เป็น unique id ที่เอาไว้ระบุตัว client
	// ExpirationHours เป็นเวลาหมดอายุของ token

	jwtWrapper := service.JwtWrapper{
		SecretKey:       "SvNQpBN8y3qlVrsGAYYWoJJk56LtzFHx",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}

	signedToken, err := jwtWrapper.GenerateToken(staff.StaffCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error signing token"})
		return
	}

	tokenResponse := LoginStaffResponse{
		Token: signedToken,
		Staff: StaffResponse{
			ID:        staff.ID,
			Prefix:    staff.Prefix,
			FirstName: staff.FirstName,
			LastName:  staff.LastName,
			StaffCode: staff.StaffCode,
		},
	}

	c.JSON(http.StatusOK, gin.H{"data": tokenResponse})
}

/* --- Teacher Response --- */
type ProfessorResponse struct {
	ID            uint
	TeacherName   string
	TeacherEmail  string
	ProfessorCode string
}

type LoginProfessorResponse struct {
	Token     string            `json:"token"`
	Professor ProfessorResponse `json:"professor"`
}

// POST /teacher/login
func LoginProfessor(c *gin.Context) {
	var payload LoginPayload
	var teacher entity.Professor

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().
		Raw("SELECT * FROM professors WHERE professor_id = ?", payload.UserCode).Find(&teacher); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "teacher not found"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(teacher.Password), []byte(payload.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user credentials"})
		return
	}

	// กำหนดค่า SecretKey, Issuer และระยะเวลาหมดอายุของ Token สามารถกำหนดเองได้
	// SecretKey ใช้สำหรับการ sign ข้อความเพื่อบอกว่าข้อความมาจากตัวเราแน่นอน
	// Issuer เป็น unique id ที่เอาไว้ระบุตัว client
	// ExpirationHours เป็นเวลาหมดอายุของ token

	jwtWrapper := service.JwtWrapper{
		SecretKey:       "SvNQpBN8y3qlVrsGAYYWoJJk56LtzFHx",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}

	signedToken, err := jwtWrapper.GenerateToken(teacher.ProfessorCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error signing token"})
		return
	}

	tokenResponse := LoginProfessorResponse{
		Token: signedToken,
		Professor: ProfessorResponse{
			ID:            teacher.ID,
			TeacherName:   teacher.TeacherName,
			TeacherEmail:  teacher.TeacherEmail,
			ProfessorCode: teacher.ProfessorCode,
		},
	}

	c.JSON(http.StatusOK, gin.H{"data": tokenResponse})
}
