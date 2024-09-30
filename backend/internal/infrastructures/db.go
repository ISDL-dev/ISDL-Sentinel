package infrastructures

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

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

	DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", user, password, hostName, dbName))
	if err != nil {
		log.Fatalf("main sql.Open error err:%v", err)
	}

	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(100)
	DB.SetConnMaxLifetime(10 * time.Second)
}

func CloseDB() {
	log.Println("disconnect from database")
	DB.Close()
}
