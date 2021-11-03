package entity

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-64.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(
		// Student Record subsystem
		&Prefix{},
		&Faculty{},
		&Department{},
		&StudentRecord{},

		// Manage Course subsystem
		&Professor{},
		&TA{},
		&Room{},
		&Course{},
		&ManageCourse{},

		// Enrollment Register subsystem
		&Enrollment{},
		&EnrollmentType{},
		&EnrollmentItem{},

		// Bill subsystem
		&Place{},
		&PaymentType{},
		&Bill{},

		// Request Register subsystem
		&RequestStatus{},
		&RequestType{},
		&RequestRegister{},
	)

	db = database

	// เตรียมข้อมูลสำหรับระบบย่อย
	Init_Student()
	Init_ManageCourse()
	Init_Enrollment()
	Init_Bill()
	Init_RequestRegister()
}
