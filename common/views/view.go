package views

import "github.com/google/uuid"

type Student struct {
	Id        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Pid       uuid.UUID `json:"pid"`
}

type Parents struct {
	Id         uuid.UUID `json:"id"`
	FatherName string    `json:"father_name"`
	MotherName string    `json:"mother_name"`
}
