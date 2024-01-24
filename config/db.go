package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func Connection() *sql.DB {
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "cart",
		AllowNativePasswords: true,
	}
	var err error
	Db, err = sql.Open("mysql", cfg.FormatDSN())
	// fmt.Println(cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := Db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	return Db
}
