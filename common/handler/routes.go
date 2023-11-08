package handler

import (
	"CRUD/common/db"
)

func Handler(rout Serv, conn db.DB) {
	rout.Srv.GET("/parents", GetALLParent(conn))
	rout.Srv.GET("/students", GetALLStudent(conn))
	rout.Srv.POST("/students", AddStudent(conn))
	rout.Srv.PUT("/students/update", EditStudents(conn))
	rout.Srv.PUT("/parents/update", EditParents(conn))
}
