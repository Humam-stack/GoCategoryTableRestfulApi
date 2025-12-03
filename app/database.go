package app

import (
	"belajar-golang-api/helpers"
	"database/sql"
	"time"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_restful_api?parseTime=true")
	helpers.PanicIfError(err)

	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetMaxIdleConns(5)
	return db
}
