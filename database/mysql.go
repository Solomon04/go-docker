package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"log"
	"os"
)

var Db *sql.DB

func InitDB() {
	// Use root:dbpass@tcp(172.17.0.2)/hoopspots, if you're using Windows.
	db, err := sql.Open(os.Getenv("DATABASE_CONNECTION"), os.Getenv("DATABASE_USER")+":"+os.Getenv("DATABASE_PASSWORD")+"@tcp("+os.Getenv("DATABASE_HOST")+")/"+os.Getenv("DATABASE_NAME")+"?parseTime=true")
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	Db = db
}
