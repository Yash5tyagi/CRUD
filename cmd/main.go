package main

import (
	"CRUD/common/db"
	"CRUD/common/handler"
	"context"
)

func main() {
	con := db.New(db.Init())
	rt := handler.New()
	handler.Handle(*rt, *con)
	rt.Srv.Run(":3000")
	defer db.Init().Close(context.Background())
}
