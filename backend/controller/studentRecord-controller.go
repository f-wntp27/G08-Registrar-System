package controller

import (
	"net/http"
	"sa-project-g08/backend/entity"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

//Prefix

// POST /prefixes
func CreatePrefix(c *gin.Context) {
	var prefix entity.Prefix
	if err := c.ShouldBindJSON(&prefix); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&prefix).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": prefix})
}

// GET /prefix/:id
func GetPrefix(c *gin.Context) {
	var prefix entity.Prefix
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM prefixes WHERE id = ?", id).Find(&prefix).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": prefix})
}

// GET /prefixes
func ListPrefixes(c *gin.Context) {
	var prefixes []entity.Prefix
	if err := entity.DB().Raw("SELECT * FROM prefixes").Find(&prefixes).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": prefixes})
}

//Department

// POST /departments
func CreateDepartment(c *gin.Context) {
	var department entity.Department
	if err := c.ShouldBindJSON(&department); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&department).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": department})
}

// GET /departments/faculty/:id
func ListDepartmentByFacultyID(c *gin.Context) {
	id := c.Param("id")
	var department []entity.Department
	if err := entity.DB().Preload("Faculty").Raw("SELECT * FROM departments WHERE faculty_id = ?", id).Find(&department).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": department})
}

// GET /departments
func ListDepartments(c *gin.Context) {
	var departments []entity.Department
	if err := entity.DB().Preload("Faculty").Raw("SELECT * FROM departments").Find(&departments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": departments})
}

//Faculty

// POST /faculties
func CreateFaculty(c *gin.Context) {
	var faculty entity.Faculty
	if err := c.ShouldBindJSON(&faculty); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&faculty).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": faculty})
}

// GET /faculty/:id
func GetFaculty(c *gin.Context) {
	var faculty entity.Faculty
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM faculties WHERE id = ?", id).Find(&faculty).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": faculty})
}

// GET /faculties
func ListFaculties(c *gin.Context) {
	var faculties []entity.Faculty
	if err := entity.DB().Raw("SELECT * FROM faculties").Find(&faculties).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": faculties})
}

//StudentRecord

// POST /student_records
func CreateStudentRecord(c *gin.Context) {
	var prefix entity.Prefix
	var department entity.Department
	var advisor entity.Professor
	var studentrecord entity.StudentRecord

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร studentRecord
	if err := c.ShouldBindJSON(&studentrecord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา prefix ด้วย id
	if tx := entity.DB().Where("id = ?", studentrecord.PrefixID).First(&prefix); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "prefix not found"})
		return
	}

	// 10: ค้นหา department ด้วย id
	if tx := entity.DB().Where("id = ?", studentrecord.DepartmentID).First(&department); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "department not found"})
		return
	}

	// 11: ค้นหา advisor ด้วย id
	if tx := entity.DB().Where("id = ?", studentrecord.AdvisorID).First(&advisor); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "advisor not found"})
		return
	}

	// 12: สร้าง StudentRecord
	sr := entity.StudentRecord{
		Prefix:      prefix,                    // โยงความสัมพันธ์กับ Entity Prefix
		FirstName:   studentrecord.FirstName,   // ตั้งค่าฟิลด์ FirstName
		LastName:    studentrecord.LastName,    // ตั้งค่าฟิลด์ LastName
		PersonalId:  studentrecord.PersonalId,  // ตั้งค่าฟิลด์ PersonalId
		StudentCode: studentrecord.StudentCode, // ตั้งค่าฟิลด์ StudentCode
		Password:    studentrecord.StudentCode, // ตั้งค่าฟิลด์ Password
		Department:  department,                // โยงความสัมพันธ์กับ Entity Department
		Advisor:     advisor,                   // โยงความสัมพันธ์กับ Entity Advisor
	}

	// เข้ารหัสลับรหัสผ่านที่ผู้ใช้กรอกก่อนบันทึกลงฐานข้อมูล
	bytes, err := bcrypt.GenerateFromPassword([]byte(sr.Password), 14)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error hashing password"})
		return
	}
	sr.Password = string(bytes)

	// 13: บันทึก
	if err := entity.DB().Create(&sr).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": sr})
}

// GET /student_records/:id
func GetStudentRecord(c *gin.Context) {
	var studentrecord entity.StudentRecord
	id := c.Param("id")
	if err := entity.DB().Preload("Prefix").Preload("Department").Preload("Department.Faculty").Preload("Advisor").Raw("SELECT * FROM student_records WHERE id = ?", id).Find(&studentrecord).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": studentrecord})
}

// GET /student_records
func ListStudentRecords(c *gin.Context) {
	var studentrecords []entity.StudentRecord
	if err := entity.DB().Preload("Prefix").Preload("Department").Preload("Department.Faculty").Preload("Advisor").Raw("SELECT * FROM student_records").Find(&studentrecords).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": studentrecords})
}
