package handler

import (
	"CRUD/common/db"
	"CRUD/common/views"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddStudent(conn db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var StudDet struct {
			FirstName  string `json:"first_name"`
			LastName   string `json:"last_name"`
			FatherName string `json:"father_name"`
			MotherName string `json:"mother_name"`
			RollNo     int    `json:"roll_no"`
		}
		if err := c.ShouldBindJSON(&StudDet); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		}

		var stud views.Student
		stud.FirstName = StudDet.FirstName
		stud.LastName = StudDet.LastName
		stud.RollNo = StudDet.RollNo

		var parent views.Parents
		parent.FatherName = StudDet.FatherName
		parent.MotherName = StudDet.MotherName
		conn.InsertParent(&parent)
		stud.PId = parent.PId
		conn.InsertStudent(&stud)
	}
}
