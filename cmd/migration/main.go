package main

import (
	"lightban/api/db"
	"lightban/api/db/migration"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	dsn := os.Getenv("DB_DSN")
	migration.Run(db.NewDB(dsn).DB)
}
