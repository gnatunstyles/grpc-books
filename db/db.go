package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DBConn *sql.DB
)

type DBRepo struct {
	*sql.DB
}

//Creating new db
func NewDB(dsn string) (*DBRepo, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("creating DB error: %s", err)
	}
	fmt.Println("database connected successfully")
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("pinging DB error: %s", err)
	}
	return &DBRepo{
		DB: db,
	}, nil
}
