package main

import (
	"database/sql"
	"fmt"
	"log"
	"rest/cmd/api"
	"rest/config"
	"rest/db"

	"github.com/go-sql-driver/mysql"
)

func main() {

	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Net:                  "tcp",
		Addr:                 config.Envs.DBHost + ":" + config.Envs.DBPort,
		DBName:               config.Envs.DBName,
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	initDBStorage(db)

	server := api.NewApiServer(config.Envs.Port, db)
	if err := api.Run(server); err != nil {
		fmt.Println("Error in running server")
		log.Fatal(err)
	}

}

func initDBStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("PING to DB success")
}
