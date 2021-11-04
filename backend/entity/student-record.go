package entity

import (
	"gorm.io/gorm"
)

type Prefix struct {
	gorm.Model
	Value string

	StudentRecords []StudentRecord `gorm:"foreignKey:PrefixID"`
	StaffAccounts  []StaffAccount  `gorm:"foreignKey:PrefixID"`
}

type StaffAccount struct {
	gorm.Model

	PrefixID *uint
	Prefix   Prefix `gorm:"references:ID"`

	FirstName string
	LastName  string
	StaffCode string `gorm:"uniqueIndex"`
	Password  string

	// 1 Staff บันทึกได้หลาย Student
	StudentRecords []StudentRecord `gorm:"foreignKey:CreatorID"`
}

type Faculty struct {
	gorm.Model
	Name string

	Departments []Department `gorm:"foreignKey:FacultyID"`
}

type Department struct {
	gorm.Model
	Name string

	FacultyID *uint

	Faculty        Faculty         `gorm:"references:ID"`
	StudentRecords []StudentRecord `gorm:"foreignKey:DepartmentID"`
	Professors     []Professor     `gorm:"foreignKey:DepartmentID"`
}

type StudentRecord struct {
	gorm.Model

	FirstName   string
	LastName    string
	PersonalId  string `gorm:"uniqueIndex"`
	StudentCode string `gorm:"uniqueIndex"`
	Password    string

	// PrefixID ทำหน้าที่เป็น FK
	PrefixID *uint
	Prefix   Prefix `gorm:"references:ID"`

	// DepartmentID ทำหน้าที่เป็น FK
	DepartmentID *uint
	Department   Department `gorm:"references:ID"`

	// AdvisorID ทำหน้าที่เป็น FK
	AdvisorID *uint
	Advisor   Professor `gorm:"references:ID"`

	// CreatorID ทำหน้าที่เป็น FK
	CreatorID *uint
	Creator   StaffAccount `gorm:"references:ID"`

	// นักศึกษา 1 คน ลงทะเบียนได้หลายใบ
	Enrollments []Enrollment `gorm:"foreignKey:OwnerID"`

	// นักศึกษา 1 คน ยื่นคำร้องได้หลายใบ
	RequestRegisters []RequestRegister `gorm:"foreignKey:OwnerID"`
}
