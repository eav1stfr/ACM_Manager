package sqlconnect

import (
	"acmmanager/internal/models"
	"acmmanager/utils"
	"github.com/jmoiron/sqlx"
	"os"
)

func GetMembersDBHandler() ([]models.Member, error) {
	db, err := sqlx.Connect("mysql", os.Getenv("CONNECTION_STRING"))
	if err != nil {
		return nil, utils.ConnectingToDatabaseError
	}
	defer db.Close()
	//query := "SELECT id, first_name, last_name, email, telegram, role"
	//var
	return nil, nil
}
