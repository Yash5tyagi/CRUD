package main

import (
	"CRUD/common/db"
	"CRUD/common/handler"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	con := db.New(db.Init())
	rt := handler.New()

	handler.Handler(*rt, *con)
	ApiHost := os.Getenv("API_HOST")
	ApiPort := os.Getenv("API_PORT")
	rt.Srv.Run(fmt.Sprintf("%s:%s", ApiHost, ApiPort))
	defer db.Init().Close(context.Background())
}
