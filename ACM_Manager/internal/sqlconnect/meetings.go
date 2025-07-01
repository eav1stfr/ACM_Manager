package sqlconnect

import (
	"acmmanager/internal/models"
	"acmmanager/utils"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"strconv"
	"time"
)

func CreateMeetingDbHandler(meeting models.Meeting) (models.Meeting, error) {
	db, err := ConnectDb()
	if err != nil {
		return models.Meeting{}, err
	}
	defer db.Close()
	var createdMeeting models.Meeting

	query := "INSERT INTO meetings (venue, time, repeated, department_id) VALUES ($1, $2, $3, $4) RETURNING *"

	err = db.Get(&createdMeeting, query, meeting.Venue, meeting.Date, meeting.Repeated, meeting.DepartmentID)
	if err != nil {
		log.Println(err)
		return createdMeeting, utils.DatabaseQueryError
	}

	if meeting.DepartmentID == nil {
		err = insertAttendanceForAllMembers(db, createdMeeting.ID)
		if err != nil {
			log.Println("Failed to populate attendance:", err)
			return createdMeeting, utils.DatabaseQueryError
		}
	} else {
		err = insertAttendanceForDepMembers(db, createdMeeting.ID, *meeting.DepartmentID)
		if err != nil {
			log.Println("Failed to populate attendance:", err)
			return createdMeeting, utils.DatabaseQueryError
		}
	}

	return createdMeeting, nil
}

func insertAttendanceForAllMembers(db *sqlx.DB, meetingID int) error {
	query := "SELECT id FROM members WHERE role = 'member' OR role = 'head'"
	var ids []int64
	err := db.Select(&ids, query)
	if err != nil {
		return err
	}
	fmt.Println(ids)
	fmt.Println(len(ids))
	query = "INSERT INTO meeting_attendance (meeting_id, member_id) VALUES ($1, $2)"
	for _, id := range ids {
		_, err = db.Exec(query, meetingID, id)
		if err != nil {
			return err
		}
	}
	return err
}

func insertAttendanceForDepMembers(db *sqlx.DB, meetingID int, depID string) error {
	query := "SELECT member_id FROM department_members WHERE department_name = $1"
	var ids []int64
	err := db.Select(&ids, query, depID)
	if err != nil {
		return utils.DatabaseQueryError
	}
	fmt.Println(ids)
	query = "INSERT INTO meeting_attendance (meeting_id, member_id) VALUES ($1, $2)"
	for _, id := range ids {
		_, err = db.Exec(query, meetingID, id)
		if err != nil {
			return err
		}
	}
	return nil
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

func GetAttendanceCount(id string, date *time.Time) (int, int, error) {
	memberId, err := strconv.Atoi(id)
	if err != nil {
		return 0, 0, utils.InvalidRequestPayloadError
	}
	db, err := ConnectDb()
	if err != nil {
		return 0, 0, err
	}
	defer db.Close()
	var query string
	args := []interface{}{memberId, "true"}
	if date != nil {
		query = "SELECT COUNT(*) FROM meeting_attendance ma JOIN meetings m ON ma.meeting_id = m.id WHERE ma.member_id = $1 AND ma.attended = $2 AND m.time >= $3"
		args = append(args, date)
	} else {
		query = "SELECT COUNT(*) FROM meeting_attendance WHERE member_id = $1 AND attended = $2"
	}

	var countAttended int
	err = db.Get(&countAttended, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, 0, utils.UnitNotFoundError
		}
		return 0, 0, utils.DatabaseQueryError
	}
	fmt.Println(countAttended)
	args[1] = "false"
	var countMissed int
	err = db.Get(&countMissed, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, 0, utils.UnitNotFoundError
		}
		return 0, 0, utils.DatabaseQueryError
	}
	return countAttended, countMissed, nil
}
