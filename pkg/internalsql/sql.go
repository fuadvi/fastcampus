package internalsql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Connect(dataSource string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return nil, err
	}

	return db, nil
}
