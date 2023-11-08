package views

import "github.com/google/uuid"

type Student struct {
	Id        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	RollNo    int       `json:"roll_no"`
}

type Parents struct {
	SId        uuid.UUID `json:"sid"`
	FatherName string    `json:"father_name"`
	MotherName string    `json:"mother_name"`
}
