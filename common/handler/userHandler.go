package handler

import (
	"CRUD/common/db"
	"CRUD/common/views"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func CheckUser(conn db.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var User struct {
			gorm.Model
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := ctx.ShouldBindJSON(&User); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		}

		rows := conn.SelectUser(User.Username)

		defer rows.Close()
		var usrs []views.User
		for rows.Next() {
			var usr views.User

			if err := rows.Scan(&usr.UserName, &usr.Password); err != nil {

				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			usrs = append(usrs, usr)
		}
		if len(usrs) == 0 {
			ctx.JSON(http.StatusBadRequest, "Wrong Username")
			return
		}
		for _, u := range usrs {
			if err := VerifyPassword(User.Password, u.Password); err != nil {
				ctx.JSON(http.StatusBadRequest, "Wrong Password")
				return
			}
			fmt.Printf("username:%s password:%s", u.UserName, u.Password)
		}
		token, err := GenerateToken(User.Username)

		if err != nil {
			log.Fatal(err)
		}
		ctx.JSON(http.StatusOK, gin.H{"token": token})
	}
}
