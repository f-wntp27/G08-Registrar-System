package controller

import (
	"net/http"

	"sa-project-g08/backend/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
		EnrollDateTime:  enrollment.EnrollDateTime.Local(),
		TotalCredit:     enrollment.TotalCredit,
		Owner:           student,
	}

	// 12: บันทึก
	if err := entity.DB().Create(&en).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// วนลูปเพื่อค้นหารายวิชาใน 1 ใบลงทะเบียน
	entity.DB().Transaction(func(tx *gorm.DB) error {
		var items []entity.EnrollmentItem
		for _, eni := range enrollment.EnrollmentItems {
			var enrollType entity.EnrollmentType
			var manageCourse entity.ManageCourse

			// 13: ค้นหา ManageCourse ด้วย id
			if tx1 := entity.DB().Where("id = ?", eni.ManageCourseID).First(&manageCourse); tx.RowsAffected == 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "manage course not found"})
				return tx1.Error
			}

			// 14: ค้นหา EnrollmentType ด้วย id
			if tx1 := entity.DB().Where("id = ?", eni.EnrollmentTypeID).First(&enrollType); tx.RowsAffected == 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "course type not found"})
				return tx1.Error
			}

			// 15: สร้าง EnrollmentItem
			it := entity.EnrollmentItem{
				ManageCourse:   manageCourse,
				EnrollmentType: enrollType,
				Enrollment:     en,
			}

			items = append(items, it)

		}

		// 16: บันทึก
		if err := entity.DB().Create(&items).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return err
		}

		entity.DB().Table("enrollment_items").Where("enrollment_id = ?", en.ID).Find(&en.EnrollmentItems)
		c.JSON(http.StatusOK, gin.H{"data": en})

		return nil
	})

}

// GET /enrollments
func ListEnrollments(c *gin.Context) {
	var enrollments []entity.Enrollment
	if err := entity.DB().Table("enrollments").
		Preload("Owner").
		Preload("EnrollmentItems").
		Preload("EnrollmentItems.EnrollmentType").
		Preload("EnrollmentItems.ManageCourse").
		Preload("EnrollmentItems.ManageCourse.Course").Find(&enrollments).Error; err != nil {
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
