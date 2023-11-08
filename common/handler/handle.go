package handler

import "CRUD/common/db"

func Handle(rout Serv, conn db.DB) {
	rout.Srv.GET("/parents", conn.GetParent)

}
