package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	userTable       = "users"
	todoListTable   = "lists"
	usersListsTable = "users_lists"
	todoItemsTable  = "items"
	listsItemsTable = "lists_items"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBname   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) *sql.DB {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s  password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.DBname, cfg.Password, cfg.SSLMode))
	if err != nil {
		log.Fatalf("error occured while connection to db: %s", err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("error occured while connection to db: %s", err.Error())
	}
	return db
}
