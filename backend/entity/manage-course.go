package entity

import (
	"time"

	"gorm.io/gorm"
)

type Professor struct {
	gorm.Model
	TeacherName   string
	TeacherEmail  string
	ProfessorCode string `gorm:"uniqueIndex"`
	Password      string

	DepartmentID *uint
	Department   Department `gorm:"references:ID"`

	StudentRecords []StudentRecord `gorm:"foreignKey:AdvisorID"`

	ManageCourses []ManageCourse `gorm:"foreignKey:ProfessorID"`
}

type TA struct {
	gorm.Model
	TaCode        string
	Name          string
	ManageCourses []ManageCourse `gorm:"foreignKey:TAID"`
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

	ProfessorID *uint
	Professor   Professor `gorm:"references:id"`

	TAID *uint
	TA   TA `gorm:"references:id"`

	EnrollmentItems []EnrollmentItem `gorm:"foreignKey:ManageCourseID"`

	RequestRegisters []RequestRegister `gorm:"foreignKey:ManageCourseID"`
}
