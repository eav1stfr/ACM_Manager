package sqlconnect

import (
	"acmmanager/internal/models"
	"acmmanager/utils"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func GetMembersDBHandler(dep string, idStr string) ([]models.Member, error) {
	db, err := ConnectDb()
	if err != nil {
		return nil, utils.ConnectingToDatabaseError
	}
	defer db.Close()
	var members []models.Member
	if dep != "" {
		query := "SELECT member_id FROM department_members WHERE department_name = $1"
		var ids []int64
		err = db.Select(&ids, query, dep)
		if err != nil {
			return nil, utils.DatabaseQueryError
		}
		for _, id := range ids {
			member, err := GetOneMemberByIdDbHandler(id)
			if err != nil {
				return nil, err
			}
			members = append(members, member)
		}
	} else if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return nil, utils.InvalidRequestPayloadError
		}
		member, err := GetOneMemberByIdDbHandler(int64(id))
		members = append(members, member)
	} else {
		query := "SELECT id, first_name, last_name, email, telegram, role, birthday FROM members WHERE role = $1 OR role = $2"

		err = db.Select(&members, query, "member", "head")
		if err != nil {
			return nil, utils.DatabaseQueryError
		}
	}
	return members, nil
}

func GetOneMemberByIdDbHandler(id int64) (models.Member, error) {
	query := "SELECT id, first_name, last_name, email, telegram, role, birthday FROM members WHERE id = $1"
	db, err := ConnectDb()
	if err != nil {
		return models.Member{}, utils.ConnectingToDatabaseError
	}
	var member models.Member
	err = db.Get(&member, query, id)
	if err == sql.ErrNoRows {
		return models.Member{}, utils.UnitNotFoundError
	} else if err != nil {
		return models.Member{}, utils.DatabaseQueryError
	}
	return member, nil
}

func PostMembersDBHandler(newMembers []models.Member) ([]models.Member, error) {
	db, err := ConnectDb()
	if err != nil {
		return nil, utils.ConnectingToDatabaseError
	}
	defer db.Close()

	tx, err := db.Beginx()
	if err != nil {
		return nil, utils.DatabaseQueryError
	}

	query := "INSERT INTO members (id, first_name, last_name, email, telegram, role, birthday) VALUES (:id, :first_name, :last_name, :email, :telegram, :role, :birthday)"
	addedMembers := make([]models.Member, 0, len(newMembers))

	for _, member := range newMembers {
		namedArgs := map[string]interface{}{
			"id":         member.ID,
			"first_name": member.FirstName,
			"last_name":  member.LastName,
			"email":      member.Email,
			"telegram":   member.Telegram,
			"role":       member.Role,
			"birthday":   member.Birthday,
		}
		stmt, err := tx.PrepareNamed(query)
		if err != nil {
			tx.Rollback()
			return nil, utils.DatabaseQueryError
		}
		defer stmt.Close()

		_, err = stmt.Exec(namedArgs)
		if err != nil {
			tx.Rollback()
			return nil, utils.DatabaseQueryError
		}
		addedMembers = append(addedMembers, member)
	}
	if err = tx.Commit(); err != nil {
		return nil, utils.DatabaseQueryError
	}
	return addedMembers, nil
}

func DeleteMembersHandler(ids []int64) ([]int64, error) {
	deletedIds := make([]int64, 0, len(ids))
	query := "DELETE FROM members WHERE id = :id"
	db, err := ConnectDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	tx, err := db.Beginx()
	if err != nil {
		return nil, utils.DatabaseQueryError
	}

	stmt, err := tx.PrepareNamed(query)
	if err != nil {
		tx.Rollback()
		return nil, utils.DatabaseQueryError
	}
	defer stmt.Close()

	for _, id := range ids {
		res, err := stmt.Exec(map[string]interface{}{"id": id})
		if err != nil {
			tx.Rollback()
			return nil, utils.DatabaseQueryError
		}
		rowsAffected, err := res.RowsAffected()

		if err != nil {
			tx.Rollback()
			return nil, utils.DatabaseQueryError
		}

		if rowsAffected == 0 {
			tx.Rollback()
			var noRowsErr = &utils.AppError{}
			noRowsErr.SetErrMessage(fmt.Sprintf("Unit with id %d is not found", id))
			noRowsErr.SetStatusCode(http.StatusNotFound)
			return nil, noRowsErr
		}

		deletedIds = append(deletedIds, id)
	}

	if err = tx.Commit(); err != nil {
		return nil, utils.DatabaseQueryError
	}

	return deletedIds, nil
}

func PatchMembersDbHandler(updates []map[string]interface{}) error {
	db, err := ConnectDb()
	if err != nil {
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return utils.StartingTransactionError
	}
	for _, update := range updates {
		idStr, ok := update["id"].(float64)
		log.Println("ID IS:", idStr)
		if !ok {
			log.Println("inside db handler err 1")
			tx.Rollback()
			return utils.InvalidRequestPayloadError
		}
		id := int64(idStr)
		if err != nil {
			log.Println("inside db handler err 1")
			return utils.InvalidRequestPayloadError
		}
		memberFromDb, err := GetOneMemberByIdDbHandler(id)
		if err != nil {
			tx.Rollback()
			return err
		}
		if v, ok := update["first_name"].(string); ok {
			memberFromDb.FirstName = v
		}
		if v, ok := update["last_name"].(string); ok {
			memberFromDb.LastName = v
		}
		if v, ok := update["email"].(string); ok {
			memberFromDb.Email = v
		}
		if v, ok := update["telegram"].(string); ok {
			memberFromDb.Telegram = v
		}
		if v, ok := update["role"].(string); ok {
			memberFromDb.Role = v
		}
		if v, ok := update["birthday"].(time.Time); ok {
			memberFromDb.Birthday = v
		}
		query := "UPDATE members SET first_name = $1, last_name = $2, email = $3, telegram = $4, role = $5, birthday = $6 WHERE id = $7"
		_, err = tx.Exec(query,
			memberFromDb.FirstName,
			memberFromDb.LastName,
			memberFromDb.Email,
			memberFromDb.Telegram,
			memberFromDb.Role,
			memberFromDb.Birthday,
			memberFromDb.ID)
		if err != nil {
			tx.Rollback()
			return utils.DatabaseQueryError
		}
	}
	if err := tx.Commit(); err != nil {
		return utils.CommitingTransactionError
	}
	return nil
}

//func getDepartmentsOfMemberDbHandler(memberId string) ([]string, error) {
//
//}
