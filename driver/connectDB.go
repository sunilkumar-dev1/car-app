package driver

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func DbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "Sunil"
	dbPass := "ZopSmart@123"
	dbName := "Car"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Printf("db connection done!\n")
	}
	return db
}
