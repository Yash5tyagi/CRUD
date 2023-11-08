package handler

import (
	"CRUD/common/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RemoveStudent(conn db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var StudDet struct {
			Id string `json:"id"`
		}
		if err := c.ShouldBindJSON(&StudDet); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		}
	}
}
