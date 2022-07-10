package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DBConn *sql.DB
)
