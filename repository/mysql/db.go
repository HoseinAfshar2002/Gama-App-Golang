package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// ...

type MysqlDB struct {
	db *sql.DB
}

func New() *MysqlDB {
	db, err := sql.Open("mysql", "admin:123456@(localhost:3308)/myapp")
	if err != nil {
		panic(fmt.Errorf("can not open mysql database: %v", err))
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &MysqlDB{db}
}
