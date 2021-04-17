package mysql

import (
	"database/sql"

	"github.com/timopattikawa/mtp-restful-product/product/repository"
)

type mysqlProduct struct {
	dbCon *sql.DB
}

func New(db *sql.DB) repository.RepositoryProduct {
	return &mysqlProduct{
		dbCon: db,
	}
}
