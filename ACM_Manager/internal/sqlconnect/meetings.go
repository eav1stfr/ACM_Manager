package sqlconnect

import (
	"acmmanager/internal/models"
	"acmmanager/utils"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"strconv"
)

func CreateMeetingDbHandler(meeting models.Meeting) (models.Meeting, error) {
	db, err := ConnectDb()
	if err != nil {
		return models.Meeting{}, err
	}
	defer db.Close()
	query := "INSERT INTO meetings (venue, time, repeated, department_id) VALUES ($1, $2, $3, $4) RETURNING *"
	var createdMeeting models.Meeting
	err = db.Get(&createdMeeting, query, meeting.Venue, meeting.Date, meeting.Repeated, meeting.DepartmentID)
	if err != nil {
		log.Println(err)
		return createdMeeting, utils.DatabaseQueryError
	}
	return createdMeeting, nil
}

func DeleteMeetingDbHandler(meetingId string) (int, error) {
	db, err := ConnectDb()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	id, err := strconv.Atoi(meetingId)
	if err != nil {
		return 0, utils.InvalidRequestPayloadError
	}

	query := "DELETE FROM meetings WHERE id = $1"
	_, err = db.Exec(query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, utils.UnitNotFoundError
		}
		return 0, utils.DatabaseQueryError
	}
	return id, nil
}

func GetMeetingsForWeekDbHandler(memberId string) ([]models.Meeting, error) {
	db, err := ConnectDb()
	if err != nil {
		return nil, utils.ConnectingToDatabaseError
	}
	defer db.Close()

	query := "SELECT * FROM meetings WHERE time >= NOW() AND time <= NOW() + INTERVAL '7 days' AND department_id IS NULL"
	var meetings []models.Meeting
	err = db.Select(&meetings, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return meetings, nil
		}
		return nil, utils.DatabaseQueryError
	}
	if memberId != "" {
		id, err := strconv.Atoi(memberId)
		if err != nil {
			return nil, utils.InvalidRequestPayloadError
		}
		deps, err := getDepsOfMember(id)
		if err != nil {
			return nil, err
		}

		if len(deps) > 0 {
			queryBase := "SELECT * FROM meetings WHERE time >= NOW() AND time <= NOW() + INTERVAL '7 days' AND department_id IN (?)"

			// expand slice into query placeholders
			queryWithIn, args, err := sqlx.In(queryBase, deps)
			if err != nil {
				return nil, err
			}

			queryWithIn = db.Rebind(queryWithIn)

			var meetingsForDep []models.Meeting
			err = db.Select(&meetingsForDep, queryWithIn, args...)
			if err != nil {
				return nil, utils.DatabaseQueryError
			}
			fmt.Println(queryWithIn)
			meetings = append(meetings, meetingsForDep...)
		}

	}
	return meetings, nil
}

func getDepsOfMember(memberId int) ([]string, error) {
	db, err := ConnectDb()
	if err != nil {
		return nil, utils.ConnectingToDatabaseError
	}
	defer db.Close()
	query := "SELECT department_name FROM department_members WHERE member_id = $1"
	var deps []string
	err = db.Select(&deps, query, memberId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.UnitNotFoundError
		}
		return nil, utils.DatabaseQueryError
	}
	return deps, nil
}
