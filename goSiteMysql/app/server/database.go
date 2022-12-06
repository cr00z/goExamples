package server

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func InitDb() (err error) {
	//login:password@tcp(host:port)/dbname
	dataSourceName := "user:password@tcp(localhost:3306)/mydb"
	Db, err = sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		return
	}
	err = Db.Ping()
	return
}
