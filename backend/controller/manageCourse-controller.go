package controller

import (
	"net/http"

	"sa-project-g08/backend/entity"

	"github.com/gin-gonic/gin"
)

//ManageCourses
// POST /ManageCourses
func CreateManageCourse(c *gin.Context) {
	var manageCourse entity.ManageCourse
	var professor entity.Professor
	var course entity.Course
	var ta entity.TA
	var room entity.Room

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร manageCourse
	if err := c.ShouldBindJSON(&manageCourse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา course ด้วย id
	if tx := entity.DB().Where("id = ?", manageCourse.CourseID).First(&course); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "course not found"})
		return
	}

	// 10: ค้นหา ta ด้วย id
	if tx := entity.DB().Where("id = ?", manageCourse.TAID).First(&ta); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ta not found"})
		return
	}

	// 11: ค้นหา room ด้วย id
	if tx := entity.DB().Where("id = ?", manageCourse.RoomID).First(&room); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room not found"})
		return
	}

	// 12: สร้าง manageCourse
	wv := entity.ManageCourse{
		Professor:        professor,                     // โยงความสัมพันธ์กับ Entity Teacher
		Course:           course,                        // โยงความสัมพันธ์กับ Entity Course
		TA:               ta,                            // โยงความสัมพันธ์กับ Entity TA
		Room:             room,                          // โยงความสัมพันธ์กับ Entity Room
		ManageCourseTime: manageCourse.ManageCourseTime, // ตั้งค่าฟิลด์ ManageCourseTime
		Group:            manageCourse.Group,            // ตั้งค่าฟิลด์ Group
		Trimester:        manageCourse.Trimester,        // ตั้งค่าฟิลด์ Trimester
		TeachingTime:     manageCourse.TeachingTime,     // ตั้งค่าฟิลด์ Teaching_Time
		UngraduatedYear:  manageCourse.UngraduatedYear,  // ตั้งค่าฟิลด์ Ungraduared_year
	}

	// 13: บันทึก
	if err := entity.DB().Create(&wv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": wv})
}

// GET /manageCourses/:id
func GetManageCourse(c *gin.Context) {
	var manageCourse entity.ManageCourse
	id := c.Param("id")
	if err := entity.DB().Preload("Professor").Preload("Course").Preload("TA").Preload("Room").Raw("SELECT * FROM manageCourses WHERE id = ?", id).Find(&manageCourse).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": manageCourse})
}

// GET /manageCourses
func ListManageCourses(c *gin.Context) {
	var manageCourses []entity.ManageCourse
	if err := entity.DB().Preload("Professor").Preload("Course").Preload("TA").Preload("Room").Raw("SELECT * FROM manageCourses").Find(&manageCourses).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": manageCourses})
}

// DELETE /managCourses/:id
func DeleteManagCourse(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM managCourses WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "managecourse not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

//Course
// POST /courses
func CreateCourse(c *gin.Context) {
	var course entity.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&course).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": course})
}

// GET /course/:id
func GetCourse(c *gin.Context) {
	var course entity.Course
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM courses WHERE id = ?", id).Scan(&course).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": course})
}

// GET /courses
func ListCourses(c *gin.Context) {
	var courses []entity.Course
	if err := entity.DB().Raw("SELECT * FROM courses").Scan(&courses).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": courses})
}

//Room
// POST /rooms
func CreateRoom(c *gin.Context) {
	var room entity.Room
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&room).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": room})
}

// GET /room/:id
func GetRoom(c *gin.Context) {
	var room entity.Room
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM rooms WHERE id = ?", id).Scan(&room).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": room})
}

// GET /rooms
func ListRooms(c *gin.Context) {
	var rooms []entity.Room
	if err := entity.DB().Raw("SELECT * FROM rooms").Scan(&rooms).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rooms})
}

//TA
// POST /tas
func CreateTA(c *gin.Context) { //Ta
	var ta entity.TA
	if err := c.ShouldBindJSON(&ta); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&ta).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ta})
}

// GET /ta/:id
func GetTA(c *gin.Context) {
	var ta entity.TA

	id := c.Param("id")
	if err := entity.DB().Preload("Owner").Raw("SELECT * FROM tas WHERE id = ?", id).Find(&ta).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ta})
}

// GET /tas
func ListTAs(c *gin.Context) {
	var tas []entity.TA
	if err := entity.DB().Preload("Owner").Raw("SELECT * FROM tas").Find(&tas).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tas})
}

// GET /professor/:id
func GetProfessor(c *gin.Context) {
	var professor entity.Professor
	id := c.Param("id")
	if err := entity.DB().Preload("Owner").Raw("SELECT * FROM professors WHERE id = ?", id).Find(&professor).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": professor})
}

// GET /professors
func ListProfessors(c *gin.Context) {
	var professors []entity.Professor
	if err := entity.DB().Preload("Owner").Raw("SELECT * FROM professors").Find(&professors).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": professors})
}
