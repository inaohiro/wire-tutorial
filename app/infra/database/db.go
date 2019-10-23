package database

import (
	"database/sql"
	"log"
)

type DataSource struct {
	DriverName     string
	DataSourceName string
}

func Open(src DataSource) (*sql.DB, func(), error) {
	db, err := sql.Open(src.DriverName, src.DataSourceName)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}
	return db, cleanup, nil
}
