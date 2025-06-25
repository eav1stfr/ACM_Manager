package sqlconnect

import (
	"acmmanager/utils"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
)

func ConnectDb() (*sqlx.DB, error) {
	connectionString := os.Getenv("CONNECTION_STRING")
	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return nil, utils.ConnectingToDatabaseError
	}
	if err = db.Ping(); err != nil {
		return nil, utils.ConnectingToDatabaseError
	}
	return db, nil
}
