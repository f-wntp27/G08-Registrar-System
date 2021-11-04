package controller

import (
	"net/http"

	"sa-project-g08/backend/entity"

	"github.com/gin-gonic/gin"
)

// GET /enrollment_types
func ListEnrollmentTypes(c *gin.Context) {
	var enrollTypes []entity.EnrollmentType
	if err := entity.DB().Table("enrollment_types").Scan(&enrollTypes).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": enrollTypes})
}

// GET /manage_courses/course/:id
func ListManageCoursesFromCourseID(c *gin.Context) {
	id := c.Param("id")
	var manageCourse []entity.ManageCourse
	if err := entity.DB().Table("manage_courses").
		Preload("Course").Preload("Professor").Preload("Room").Where("course_id = ?", id).Find(&manageCourse).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": manageCourse})
}

// POST /enrollment
func CreateEnrollment(c *gin.Context) {
	var student entity.StudentRecord
	var enrollment entity.Enrollment

	// ผลลัพธ์ที่ได้จากข้อ 9 จะถูก bind เข้าตัวแปร enrollment
	if err := c.ShouldBindJSON(&enrollment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 10: ค้นหา student ด้วย id
	if tx := entity.DB().Where("id = ?", enrollment.OwnerID).First(&student); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "student not found"})
		return
	}

	// 11: สร้าง EnrollResult
	en := entity.Enrollment{
		EnrollYear:      enrollment.EnrollYear,
		EnrollTrimester: enrollment.EnrollTrimester,
		EnrollDateTime:  enrollment.EnrollDateTime,
		TotalCredit:     enrollment.TotalCredit,
		Owner:           student,
	}

	// 12: บันทึก
	if err := entity.DB().Create(&en).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// วนลูปเพื่อบันทึกรายวิชาใน 1 ใบลงทะเบียน
	for _, enc := range enrollment.EnrollmentItems {
		var enrollType entity.EnrollmentType
		var manageCourse entity.ManageCourse
		var thisEnrollment entity.Enrollment

		if tx := entity.DB().Where("owner_id = ? AND enroll_year = ? AND enroll_trimester = ? AND enroll_date_time = ?",
			enrollment.OwnerID, enrollment.EnrollYear, enrollment.EnrollTrimester, enrollment.EnrollDateTime).First(&thisEnrollment); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "enrollment not found"})
			return
		}

		// 13: ค้นหา ManageCourse ด้วย id
		if tx := entity.DB().Where("id = ?", enc.ManageCourseID).First(&manageCourse); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "manage course not found"})
			return
		}

		// 14: ค้นหา EnrollmentType ด้วย id
		if tx := entity.DB().Where("id = ?", enc.EnrollmentTypeID).First(&enrollType); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "course type not found"})
			return
		}

		// 15: สร้าง EnrollmentItem
		cel := entity.EnrollmentItem{
			ManageCourse:   manageCourse,
			EnrollmentType: enrollType,
			Enrollment:     en,
		}

		// 16: บันทึก
		if err := entity.DB().Create(&cel).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
	entity.DB().Table("enrollment_items").Where("enrollment_id = ?", en.ID).Find(&en.EnrollmentItems)
	c.JSON(http.StatusOK, gin.H{"data": en})
}

// GET /enrollments
func ListEnrollments(c *gin.Context) {
	var enrollments []entity.Enrollment
	if err := entity.DB().Table("enrollments").
		Preload("StudentRecord").
		Preload("EnrollmentItems").
		Preload("EnrollmentItems.EnrollmentType").
		Preload("EnrollmentItems.ManageCourse").
		Preload("EnrollmentItems.ManageCourse.Course").Scan(&enrollments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": enrollments})
}

// GET /enrollments/student/:id
func ListEnrollmentsFromStudentID(c *gin.Context) {
	id := c.Param("id")
	var enrollments []entity.Enrollment
	if err := entity.DB().Table("enrollments").
		Preload("Owner").Preload("EnrollmentItems").
		Preload("EnrollmentItems.EnrollmentType").
		Preload("EnrollmentItems.ManageCourse").
		Preload("EnrollmentItems.ManageCourse.Course").Where("owner_id = ?", id).Find(&enrollments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": enrollments})
}
