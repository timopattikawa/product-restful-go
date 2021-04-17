package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewConnetions(connString string) (*sql.DB, error) {
	conn, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
