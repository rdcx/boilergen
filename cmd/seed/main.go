package main

import (
	"lightban/api/db"
	"lightban/api/db/seed"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	dsn := os.Getenv("DB_DSN")
	seed.Run(db.NewDB(dsn))
}
