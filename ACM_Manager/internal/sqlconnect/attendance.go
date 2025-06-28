package sqlconnect

import (
	"acmmanager/utils"
	"database/sql"
	"strconv"
)

func MarkAttendanceDbHandler(memberIdStr, meetingIdStr string) error {
	db, err := ConnectDb()
	if err != nil {
		return utils.ConnectingToDatabaseError
	}
	defer db.Close()

	memberId, err := strconv.Atoi(memberIdStr)
	if err != nil {
		return utils.InvalidRequestPayloadError
	}
	meetingId, err := strconv.Atoi(meetingIdStr)
	if err != nil {
		return utils.InvalidRequestPayloadError
	}
	query := "UPDATE meeting_attendance SET attended = true WHERE meeting_id = $1 AND member_id = $2"
	_, err = db.Exec(query, meetingId, memberId)
	if err != nil {
		if err == sql.ErrNoRows {
			return utils.UnitNotFoundError
		}
		return utils.DatabaseQueryError
	}
	return nil
}

func GetReasonOfAbsence(memberIdStr, meetingIdStr, comment string) error {
	db, err := ConnectDb()
	if err != nil {
		return utils.ConnectingToDatabaseError
	}
	defer db.Close()
	memberId, err := strconv.Atoi(memberIdStr)
	if err != nil {
		return utils.InvalidRequestPayloadError
	}
	meetingId, err := strconv.Atoi(meetingIdStr)
	if err != nil {
		return utils.InvalidRequestPayloadError
	}
	query := "UPDATE meeting_attendance SET comment = $1 WHERE meeting_id = $2 AND member_id = $3"
	_, err = db.Exec(query, comment, meetingId, memberId)
	if err != nil {
		if err == sql.ErrNoRows {
			return utils.UnitNotFoundError
		}
		return utils.DatabaseQueryError
	}
	return nil
}
