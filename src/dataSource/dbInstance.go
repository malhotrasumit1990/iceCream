package dataSource

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLDataSource struct {
	*sql.DB
}

func NewMySqlDataStore(connectionString string) (*MySQLDataSource, error) {
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
	}

	return &MySQLDataSource{DB: db}, nil
}

func (mds *MySQLDataSource) Close() {
	mds.Close()
}
