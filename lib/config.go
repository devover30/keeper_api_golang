package lib

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDB() (*sql.DB, error) {
	mysqlURI := os.Getenv("MYSQL_URI")
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return db, nil
}
