package entity

import (
	"time"

	"gorm.io/gorm"
)

type RequestStatus struct {
	gorm.Model
	Name string `json:"name"`

	RequestRegisters []RequestRegister `json:"-" gorm:"foreignKey:RequestStatusID"`
}

type RequestType struct {
	gorm.Model
	Name string `json:"name"`

	RequestRegisters []RequestRegister `json:"-" gorm:"foreignKey:RequestTypeID"`
}

type RequestRegister struct {
	gorm.Model
	RequestTime time.Time `json:"requestTime"`

	ManageCourseID *uint        `json:"manageCourseID"`
	ManageCourse   ManageCourse `json:"manageCourse"`

	RequestTypeID *uint       `json:"requestTypeID"`
	RequestType   RequestType `json:"requestType"`

	RequestStatusID *uint         `json:"requestStatusID"`
	RequestStatus   RequestStatus `json:"requestStatus"`

	OwnerID *uint         `json:"ownerID"`
	Owner   StudentRecord `json:"owner"`
}
