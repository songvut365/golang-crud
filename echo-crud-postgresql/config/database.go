package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func SetupDatabase() {
	var err error
	DB, err = sql.Open("postgres", "user=postgres password=1234 dbname=Company sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	} else {
		log.Println("DB Connected ...")
	}
}
