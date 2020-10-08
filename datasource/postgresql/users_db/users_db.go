package users_db

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
	"os"
)

const (
	db_host      = "db_host"
	db_port      = "db_port"
	db_user_name = "db_user_name"
	db_user_pass = "db_user_pass"
	db_name      = "db_name"
)
var (
	Client *sql.DB
	dbHost = os.Getenv(db_host)
	dbPort = os.Getenv(db_port)
	dbUser = os.Getenv(db_user_name)
	dbPass = os.Getenv(db_user_pass)
	dbName = os.Getenv(db_name)
)

func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)
	var err error
	Client, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}
