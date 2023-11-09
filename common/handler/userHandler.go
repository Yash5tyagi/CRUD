package handler

import (
	"CRUD/common/db"
	"CRUD/common/views"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func AddUser(conn db.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var Usr struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := ctx.ShouldBindJSON(&Usr); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		}
		var usr views.User
		usr.UserName = Usr.Username
		bytes, err := bcrypt.GenerateFromPassword([]byte(Usr.Password), 14)
		if err != nil {
			log.Fatal(err)
		}
		usr.Password = string(bytes)
		conn.CreateUser(&usr)
	}
}

// func CheckUser(conn db.DB) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		type user struct{
// 			Username string `json:"username"`
// 			Password string `json:"password"`
// 		}
// 		rows := conn.SelectUser()

// 		defer rows.Close()
// 		var usrs []views.User
// 		for rows.Next() {
// 			var usr views.User

// 			if err := rows.Scan(&usr., &usr.MotherName); err != nil {

// 				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			}
// 			usrs = append(usrs, usr)
// 		}
// 		fmt.Println("Parents:")
// 		for _, parent := range usrs {
// 			fmt.Printf("father name: %s, mother name: %s\n", parent.FatherName, parent.MotherName)
// 		}
// 		ctx.JSON(http.StatusOK, usrs)
// 	}
// }
