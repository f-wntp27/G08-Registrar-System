package main

import (
	"sa-project-g08/backend/controller"
	"sa-project-g08/backend/entity"
	"sa-project-g08/backend/middlewares"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	entity.SetupDatabase()

	r := gin.Default()

	r.Use(CORSMiddleware())

	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{
			// Student Record subsystem
			protected.POST("/prefixes", controller.CreatePrefix)
			protected.GET("/prefix/:id", controller.GetPrefix)
			protected.GET("/prefixes", controller.ListPrefixes)

			protected.POST("/departments", controller.CreateDepartment)
			protected.GET("/departments/faculty/:id", controller.ListDepartmentByFacultyID)
			protected.GET("/departments", controller.ListDepartments)

			protected.POST("/faculties", controller.CreateFaculty)
			protected.GET("/faculty/:id", controller.GetFaculty)
			protected.GET("/faculties", controller.ListFaculties)

			protected.POST("/student_records", controller.CreateStudentRecord)
			protected.GET("/student_records/:id", controller.GetStudentRecord)
			protected.GET("/student_records", controller.ListStudentRecords)

			// Manage Course subsystem

			// Enrollment Registration subsystem
			protected.GET("/enrollment_types", controller.ListEnrollmentTypes)
			protected.GET("/manage_courses/course/:id", controller.ListManageCoursesFromCourseID)

			protected.POST("/enrollment", controller.CreateEnrollment)

			protected.GET("/enrollments", controller.ListEnrollments)
			protected.GET("/enrollments/student/:id", controller.ListEnrollmentsFromStudentID)

			// Bill subsystem
			protected.POST("/bill", controller.CreateBill)

			protected.GET("/bills", controller.ListBills)
			protected.GET("/bills/:id", controller.ListIDBill)

			protected.GET("/payment_types", controller.GetPaymentType)
			protected.GET("/places", controller.GetPlace)

			protected.GET("/bills/enrollments/:id", controller.ListEnrollmentForBill)

			// Request Register subsystem
			protected.GET("/requestregister/type", controller.ListRequestType)
			protected.GET("/requestregister/status", controller.ListRequestStatus)

			protected.POST("/requestregister", controller.CreateRequestRegister)
			protected.GET("/requestregisters", controller.ListRequestRegister)
			protected.DELETE("/requestregister", controller.DeleteRequestRegister)
		}
	}

	// Login
	r.POST("/student/login", controller.LoginStudent)
	r.POST("/staff/login", controller.LoginStaff)
	r.POST("/teacher/login", controller.LoginTeacher)

	// Run server
	r.Run()
}
