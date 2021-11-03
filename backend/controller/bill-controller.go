package controller

import (
	"net/http"

	"sa-project-g08/backend/entity"

	"github.com/gin-gonic/gin"
)

//BILL
// POST /bill //รับข้อมูลมาจาก Frontend มาบันทึกลง DB
func CreateBill(c *gin.Context) {

	var bill entity.Bill
	//ตัวเเปรเช็คบิลซ้ำ
	var check_bill entity.Bill
	var payment_types entity.PaymentType
	var place entity.Place
	//ดึงของใบลงทะเบียนเรียน
	var enrollment entity.Enrollment

	//ขั้นตอนที่ 8  = บันทึกใบชำระเงินค่าลงทะเบียนเรียน
	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร bill จาก communication diagram
	if err := c.ShouldBindJSON(&bill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//ค้นหา id ของ enrollment ที่ชำระ bill แล้ว ถ้ามีใน bill ส่ง bill ที่มี enrollment ที่ซ้ำออกมา
	if tx := entity.DB().Where("enrollment_id = ?", bill.EnrollmentID).Find(&check_bill); tx.RowsAffected != 0 {
		c.JSON(http.StatusOK, gin.H{"billDuplicate": check_bill})
		return
	}

	// 9: ค้นหา payment_types ด้วย id
	if tx := entity.DB().Where("id = ?", bill.PaymentTypeID).First(&payment_types); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "PaymentType not found"})
		return
	}

	// 10: ค้นหา place ด้วย id
	if tx := entity.DB().Where("id = ?", bill.PlaceID).First(&place); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Place not found"})
		return
	}

	// 11: ค้นหา registrations ด้วย id
	if tx := entity.DB().Where("id = ?", bill.EnrollmentID).First(&enrollment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enrollment not found"})
		return
	}

	//คำนวณหน่วยกิต
	var calculate = (enrollment.TotalCredit * 800)

	// 12: สร้าง bill
	bi := entity.Bill{
		PaymentType: payment_types, // โยงความสัมพันธ์กับ Entity payment_types
		Place:       place,         // โยงความสัมพันธ์กับ Entity place
		Enrollment:  enrollment,    // โยงความสัมพันธ์กับ Entity enrollment
		BillTime:    bill.BillTime, // ตั้งค่าฟิลด์ BillTime
		TotalPrice:  calculate,
	}

	// 13: บันทึก
	if err := entity.DB().Create(&bi).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": bi})
}

// GET /bills
func ListBills(c *gin.Context) {
	var bill []entity.Bill
	if err := entity.DB().Preload("PaymentType").
		Preload("Place").
		Preload("Enrollment").
		Preload("Enrollment.StudentRecord").
		Raw("SELECT * FROM bills").Find(&bill).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bill})
}

// GET /bills/:id  //ตาม Login
//ค้นหาบิลตามคนที่ลงทะเบียนเข้ามา
func ListIDBill(c *gin.Context) {
	id := c.Param("id")
	var enrollment []entity.Enrollment
	var bill []entity.Bill

	if tx := entity.DB().Where("student_record_id = ?", id).Find(&enrollment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "enrollment not found"})
		return
	}

	for _, enroll := range enrollment {
		var b entity.Bill
		if tx := entity.DB().Preload("Place").
			Preload("PaymentType").
			Preload("Enrollment").
			Preload("Enrollment.StudentRecord").
			Where("enrollment_id = ?", enroll.ID).Find(&b); tx.RowsAffected == 0 {
			continue
		} else {
			bill = append(bill, b)
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": bill})
}

//PAYMENT TYPE
// GET /payment_types  //การส่งข้อมูลไปที่ frontend
func GetPaymentType(c *gin.Context) {
	var paymentType []entity.PaymentType
	if err := entity.DB().Raw("SELECT * FROM payment_types").Scan(&paymentType).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": paymentType})
}

//PLACE
// GET /places //การส่งข้อมูลไปที่ frontend
func GetPlace(c *gin.Context) {
	var place []entity.Place
	if err := entity.DB().Raw("SELECT * FROM places").Scan(&place).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": place})
}

//ENROLLMENT
// GET /bills/enrollments/:id    ----> list
func ListEnrollmentForBill(c *gin.Context) {
	id := c.Param("id")
	var enrollment []entity.Enrollment
	if err := entity.DB().Raw("SELECT * FROM enrollments WHERE student_record_id = ? ORDER BY enroll_year, enroll_trimester ASC ", id).Scan(&enrollment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": enrollment})
}
