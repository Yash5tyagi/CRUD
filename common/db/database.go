package db

import (
	"CRUD/common/views"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (db *DB) GetParent(c *gin.Context) {
	rows, err := db.Conn.Query(context.Background(), "SELECT father_name,mother_name FROM parents")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		log.Fatal(err)
		return
	}
	defer rows.Close()

	var parents []views.Parents
	for rows.Next() {
		var parent views.Parents

		if err := rows.Scan(&parent.FatherName, &parent.MotherName); err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		parents = append(parents, parent)
	}
	fmt.Println("Parents:")
	for _, parent := range parents {
		fmt.Printf("father name: %s, mother name: %s\n", parent.FatherName, parent.MotherName)
	}
	c.JSON(http.StatusOK, parents)
}

func (db *DB) GetStudent() {

}
