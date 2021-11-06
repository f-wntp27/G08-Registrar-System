package entity

import (
	"time"

	"gorm.io/gorm"
)

type EnrollmentType struct {
	gorm.Model
	Name string

	EnrollmentItems []EnrollmentItem `gorm:"foreignKey:EnrollmentTypeID"`
}

type Enrollment struct {
	gorm.Model
	EnrollYear      uint
	EnrollTrimester uint
	EnrollDateTime  time.Time
	TotalCredit     uint

	OwnerID *uint
	Owner   StudentRecord `gorm:"references:ID"`

	EnrollmentItems []EnrollmentItem `gorm:"foreignKey:EnrollmentID; constraint:OnDelete:CASCADE"`
}

type EnrollmentItem struct {
	gorm.Model

	EnrollmentID *uint
	Enrollment   Enrollment `gorm:"references:ID"`

	EnrollmentTypeID *uint
	EnrollmentType   EnrollmentType `gorm:"references:ID"`

	ManageCourseID *uint
	ManageCourse   ManageCourse `gorm:"references:ID"`
}
