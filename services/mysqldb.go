package services

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func ConnectDatabse() {
	var err error

	cfg := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASS"),
		Addr:                 os.Getenv("DB_URL") + ":" + os.Getenv("DB_PORT"),
		DBName:               os.Getenv("DB_NAME"),
		AllowNativePasswords: true,
	}

	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("MySQL DB connected!")
}
