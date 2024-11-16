package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func NewMySQLStorage(config mysql.Config) (*sql.DB, error) {

	db, err := sql.Open("mysql", config.FormatDSN())

	if err != nil {
		log.Fatalln("Error in connecting to DB")
		log.Fatal(err)
		return nil, err
	}
	log.Println("Connected to DB")
	return db, nil
}
