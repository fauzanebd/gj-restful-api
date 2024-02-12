package app

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

func GetDBConnection(additionalParameters ...string) (*sql.DB, error) {
	additionalParams := strings.Join(additionalParameters, "&")
	var db *sql.DB
	var err error
	if len(additionalParameters) > 0 {
		db, err = sql.Open("mysql", fmt.Sprintf("root:makan@tcp(localhost:3306)/categoryapi?%s", additionalParams))
	} else {
		db, err = sql.Open("mysql", "root:makan@tcp(localhost:3306)/categoryapi")
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db, err
}
