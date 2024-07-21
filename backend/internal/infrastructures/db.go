package infrastructures

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	log.Println("connect to database")
	openDB()
}

func openDB() {
	var err error
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	hostName := os.Getenv("MYSQL_HOSTNAME")
	dbName := os.Getenv("MYSQL_DATABASE")

	DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?", user, password, hostName, dbName))
	if err != nil {
		log.Fatalf("main sql.Open error err:%v", err)
	}
}

func CloseDB() {
	log.Println("disconnect from database")
	DB.Close()
}
