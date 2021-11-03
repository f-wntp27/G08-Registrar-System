package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"sa-project-g08/backend/entity"

	"github.com/gin-gonic/gin"
)

// GET : /requestregister/type
// GET : /requestregister/type?id=1
func ListRequestType(c *gin.Context) {
	id := c.DefaultQuery("id", "")

	if id == "" {
		var dataType []entity.RequestType
		entity.DB().Model(&entity.RequestType{}).Find(&dataType)

		c.JSON(200, dataType)
		return
	} else {
		var dataType entity.RequestType
		entity.DB().Model(&entity.RequestType{}).Find(&dataType, entity.DB().Where("id = ?", id))

		c.JSON(200, dataType)
		return
	}

}

// GET : /requestregister/status
// GET : /requestregister/status?id=1
func ListRequestStatus(c *gin.Context) {
	id := c.DefaultQuery("id", "")

	if id == "" {
		var dataStatus []entity.RequestStatus
		entity.DB().Model(&entity.RequestStatus{}).Find(&dataStatus)

		c.JSON(200, dataStatus)
		return
	} else {
		var dataStatus entity.RequestStatus
		entity.DB().Model(&entity.RequestStatus{}).Find(&dataStatus, entity.DB().Where("id = ?", id))

		c.JSON(200, dataStatus)
		return
	}

}

// POST : /requestregister
func CreateRequestRegister(c *gin.Context) {
	userCode, _ := c.Get("userCode")

	var studentRecord entity.StudentRecord
	entity.DB().Model(&entity.StudentRecord{}).Where("student_code = ?", userCode).First(&studentRecord)

	var requestRegister entity.RequestRegister
	var manageCourse entity.ManageCourse
	var requestType entity.RequestType
	var requestStatus entity.RequestStatus

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร requestRegister
	if err := c.ShouldBindJSON(&requestRegister); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data_in, err := json.MarshalIndent(requestRegister, "", "    ")
	if err != nil {
		panic("fail")
	}
	fmt.Println(string(data_in))

	// 9: ค้นหา ManageCourse ด้วย id
	if tx := entity.DB().Model(&entity.ManageCourse{}).Joins("Course").First(&manageCourse, entity.DB().Model(&entity.ManageCourse{}).Where("`manage_courses`.`id` = ?", requestRegister.ManageCourseID)); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Course not found"})
		return
	}

	// 10: ค้นหา RequestType ด้วย id
	if tx := entity.DB().Model(&entity.RequestType{}).First(&requestType, entity.DB().Where("id = ?", requestRegister.RequestTypeID)); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "RequestType not found"})
		return
	}

	// 11: ค้นหา RequestStatus ด้วย id
	if tx := entity.DB().Model(&entity.RequestStatus{}).First(&requestStatus, entity.DB().Where("id = ?", requestRegister.RequestStatusID)); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "RequestStatus not found"})
		return
	}

	// 12: สร้าง RequestRegister
	rr := entity.RequestRegister{
		ManageCourse:  manageCourse,                // โยงความสัมพันธ์กับ Entity ManageCourse
		Owner:         studentRecord,               // โยงความสัมพันธ์กับ Entity StudentRecord
		RequestType:   requestType,                 // โยงความสัมพันธ์กับ Entity RequestType
		RequestStatus: requestStatus,               // โยงความสัมพันธ์กับ Entity RequestStatus
		RequestTime:   requestRegister.RequestTime, // ตั้งค่าฟิลด์ RequestTime
	}

	data_out, err := json.MarshalIndent(rr, "", "    ")
	if err != nil {
		panic("fail")
	}
	fmt.Println(string(data_out))

	// 13: บันทึก
	if err := entity.DB().Create(&rr).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": rr})

}

// GET : /requestregisters
func ListRequestRegister(c *gin.Context) {
	userCode, _ := c.Get("userCode")

	var studentRecord entity.StudentRecord
	entity.DB().Model(&entity.StudentRecord{}).Where("student_code = ?", userCode).First(&studentRecord)

	var RequestRegisters []entity.RequestRegister
	entity.DB().Model(&entity.RequestRegister{}).
		Joins("ManageCourse").
		Joins("RequestType").
		Joins("RequestStatus").
		Joins("Owner").
		Find(&RequestRegisters, entity.DB().Where("owner_id = ?", studentRecord.ID))
	for i, v := range RequestRegisters {
		entity.DB().Model(&entity.Teacher{}).Where("id = ?", v.ManageCourse.TeacherID).Scan(&RequestRegisters[i].ManageCourse.Teacher)
		entity.DB().Model(&entity.TA{}).Where("id = ?", v.ManageCourse.TaID).Scan(&RequestRegisters[i].ManageCourse.Ta)
		entity.DB().Model(&entity.Room{}).Where("id = ?", v.ManageCourse.RoomID).Scan(&RequestRegisters[i].ManageCourse.Room)
		entity.DB().Model(&entity.Course{}).Where("id = ?", v.ManageCourse.CourseID).Scan(&RequestRegisters[i].ManageCourse.Course)
	}

	c.JSON(200, RequestRegisters)

}

// DELETE : /requestregister?id=5004
func DeleteRequestRegister(c *gin.Context) {
	id := c.DefaultQuery("id", "")

	userCode, _ := c.Get("userCode")

	var studentRecord entity.StudentRecord
	entity.DB().Model(&entity.StudentRecord{}).Where("student_code = ?", userCode).First(&studentRecord)

	var data entity.RequestRegister
	sql := "DELETE FROM request_registers WHERE id = ? and owner_id = ?"
	entity.DB().Model(&entity.RequestRegister{}).Raw(sql, id, studentRecord.ID).Scan(&data)

	c.JSON(200, data)
}
