package models

import (
	"database/sql"
	"log"
	"os"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// Datastore presenting interface with methods for handlers
type Datastore interface {
	AllPosts() ([]*Post, error)
	GetPost(id int64) (*Post, error)
	AddPost(title string) error
	UpdatePost(id int64, title string) (bool, error)
	DeletePost(id int64) (bool, error)
}

// DbHelper presenting helper for sql.DB, implement Datastore
type DbHelper struct {
	*sql.DB
}

// EstablishConnection use DATABASE_TYPE and return DB with sql.DB inside
func EstablishConnection() *DbHelper {
	var db *sql.DB
	dbType := os.Getenv("DATABASE_TYPE")
	switch dbType {
	case "mysql":
		db = connectMysql()
	default:
		log.Fatal("invalid .env:DATABASE_TYPE, available: mysql")
	}
	return &DbHelper{db}
}

func connectMysql() *sql.DB {
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")

	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASS")

	dbname := os.Getenv("MYSQL_DB")

	db, err := sql.Open("mysql", user+":"+pass+"@tcp("+host+":"+port+")/"+dbname)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("connected to mysql db on %s:%s", host, port)

	return db
}
