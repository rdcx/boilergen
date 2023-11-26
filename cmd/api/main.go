package main

import (
	"lightban/api/db"
	"lightban/api/handler"
	"lightban/api/router"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	var port string

	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}

	godotenv.Load()

	dsn := os.Getenv("DB_DSN")
	db := db.NewDB(dsn)

	h := handler.NewHandler(db)

	r := router.SetUp(h)

	if port == "443" {
		r.RunTLS(":443", "server.crt", "server.key")
		return
	}
	r.Run(":8080")
}
