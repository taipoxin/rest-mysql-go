package models

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Datastore interface {
	AllPosts() ([]*Post, error)
}

type DB struct {
	*sql.DB
}

func EstablishConnection() *DB {

	var db *sql.DB
	dbType := os.Getenv("DATABASE_TYPE")
	switch dbType {
	case "mysql":
		db = connectMysql()
	default:
		log.Fatal("invalid .env:DATABASE_TYPE, available: mysql")
	}

	return &DB{db}
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
