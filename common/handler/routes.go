package handler

import (
	"CRUD/common/db"
	middlewares "CRUD/common/middleware"
)

func Handler(rout Serv, conn db.DB) {
	rout.Srv.GET("/parents/get", middlewares.JwtAuthMiddleware(), GetALLParent(conn))
	rout.Srv.GET("/students/get", middlewares.JwtAuthMiddleware(), GetALLStudent(conn))
	rout.Srv.POST("/students/add", AddStudent(conn))
	rout.Srv.PUT("/students/update", EditStudents(conn))
	rout.Srv.PUT("/parents/update", EditParents(conn))
	rout.Srv.DELETE("/students/delete", RemoveStudent(conn))
	rout.Srv.POST("/user/create", AddUser(conn))
	rout.Srv.POST("/user/check", CheckUser(conn))
}
