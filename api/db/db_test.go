package db_test

import (
	"lightban/api/db"
	"lightban/api/db/migration"
)

func testDB() *db.DB {
	dsn := "root:lban@tcp(127.0.0.1:3306)/lban_test?charset=utf8mb4&parseTime=True&loc=Local"
	db := db.NewDB(dsn)

	migration.Run(db.DB)

	return db
}
