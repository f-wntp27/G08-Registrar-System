package entity

import (
	"time"

	"gorm.io/gorm"
)

type TA struct {
	gorm.Model
	TaCode        string
	Name          string
	ManageCourses []ManageCourse `gorm:"foreignKey:TaID"`
}

type Room struct {
	gorm.Model
	Number        uint
	StudentCount  uint
	ManageCourses []ManageCourse `gorm:"foreignKey:RoomID"`
}

type Course struct {
	gorm.Model
	CourseCode    string
	Name          string
	Credit        uint
	ManageCourses []ManageCourse `gorm:"foreignKey:CourseID"`
}

type ManageCourse struct {
	gorm.Model
	Group            uint
	TeachingTime     uint
	UngraduatedYear  uint
	Trimester        uint
	ManageCourseTime time.Time

	CourseID *uint
	Course   Course `gorm:"references:id"`

	RoomID *uint
	Room   Room `gorm:"references:id"`

	TeacherID *uint
	Teacher   Teacher `gorm:"references:id"`

	TaID *uint
	Ta   TA `gorm:"references:id"`

	EnrollmentItems []EnrollmentItem `gorm:"foreignKey:ManageCourseID"`

	RequestRegisters []RequestRegister `gorm:"foreignKey:ManageCourseID"`
}
