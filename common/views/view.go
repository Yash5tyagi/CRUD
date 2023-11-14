package views

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Student struct {
	Id        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	RollNo    int       `json:"roll_no"`
	PId       uuid.UUID `json:"pid"`
}

type Parents struct {
	PId        uuid.UUID `json:"pid"`
	FatherName string    `json:"father_name"`
	MotherName string    `json:"mother_name"`
}

type User struct {
	gorm.Model
	UserName string `json:"username"`
	Password string `json:"password"`
}
