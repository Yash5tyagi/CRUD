package handler

import (
	"CRUD/common/db"
	"CRUD/common/views"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func EditStudents(conn db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var StudDet struct {
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			RollNo    int    `json:"roll_no"`
		}
		if err := c.ShouldBindJSON(&StudDet); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		}

		var stud views.Student
		stud.FirstName = StudDet.FirstName
		stud.LastName = StudDet.LastName
		stud.RollNo = StudDet.RollNo
		conn.UpdateStudent(&stud)
	}
}

func EditParents(conn db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var ParentDet struct {
			FatherName string    `json:"father_name"`
			MotherName string    `json:"mother_name"`
			SId        uuid.UUID `json:"sid"`
		}
		if err := c.ShouldBindJSON(&ParentDet); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		}

		var parent views.Parents
		parent.FatherName = ParentDet.FatherName
		parent.MotherName = ParentDet.MotherName
		parent.SId = ParentDet.SId
		conn.UpdateParent(&parent)
	}
}
