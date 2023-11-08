package handler

import (
	"CRUD/common/db"
	"CRUD/common/views"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetALLParent(conn db.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		rows := conn.SelectParent()

		defer rows.Close()
		var parents []views.Parents
		for rows.Next() {
			var parent views.Parents

			if err := rows.Scan(&parent.FatherName, &parent.MotherName); err != nil {

				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			parents = append(parents, parent)
		}
		fmt.Println("Parents:")
		for _, parent := range parents {
			fmt.Printf("father name: %s, mother name: %s\n", parent.FatherName, parent.MotherName)
		}
		c.JSON(http.StatusOK, parents)
	}

}

func GetALLStudent(conn db.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		rows := conn.SelectStudent()

		defer rows.Close()
		var students []views.Student
		for rows.Next() {
			var student views.Student

			if err := rows.Scan(&student.FirstName, &student.LastName, &student.RollNo); err != nil {

				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			students = append(students, student)
		}
		fmt.Println("Students:")
		for _, student := range students {
			fmt.Printf("first name: %s, last name: %s, roll no.: %d\n", student.FirstName, student.LastName, student.RollNo)
		}
		c.JSON(http.StatusOK, students)
	}

}
