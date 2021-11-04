package entity

import (
	"time"

	"gorm.io/gorm"
)

type Place struct {
	gorm.Model
	Name string

	Bills []Bill `gorm:"foreignKey:PlaceID"`
}

type PaymentType struct {
	gorm.Model
	Name string

	Bills []Bill `gorm:"foreignKey:PaymentTypeID"`
}

type Bill struct {
	//มี ID ที่เป็น primarykey
	gorm.Model

	BillTime time.Time

	EnrollmentID *uint      `gorm:"uniqueIndex"`
	Enrollment   Enrollment `gorm:"references:ID; constraint:OnDelete:CASCADE"`
	//ลบใบลงทะเบียน บิลหาย (1 - 1)

	PaymentTypeID *uint
	PaymentType   PaymentType `gorm:"references:ID"`

	PlaceID *uint
	Place   Place `gorm:"references:ID"`

	TotalPrice uint
}
