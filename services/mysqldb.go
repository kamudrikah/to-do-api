package services

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func ConnectDatabse() {
	var err error
	var (
		dbDriver = "mysql"
		dbSource = "root:password@tcp(mysql:3306)/todo"
	)

	db, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("MySQL DB connected!")
}
