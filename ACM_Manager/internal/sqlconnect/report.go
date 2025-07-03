package sqlconnect

import (
	"acmmanager/internal/models"
	"acmmanager/utils"
	"database/sql"
	"strconv"
	"time"
)

func GetMemberDataForReport(memberId string, startDate *time.Time) (models.Member, []models.Task, []models.Task, int, int, error) {
	id, err := strconv.Atoi(memberId)
	if err != nil {
		return models.Member{}, nil, nil, 0, 0, err
	}
	member, err := GetOneMemberByIdDbHandler(int64(id))
	if err != nil {
		return models.Member{}, nil, nil, 0, 0, err
	}
	tasksDone, tasksToDo, err := getTasksInfo(memberId, startDate)
	if err != nil {
		return models.Member{}, nil, nil, 0, 0, err
	}
	countAttended, countMissed, err := GetAttendanceCount(memberId, startDate)
	if err != nil {
		return models.Member{}, nil, nil, 0, 0, err
	}
	return member, tasksDone, tasksToDo, countAttended, countMissed, nil
}

func getTasksInfo(memberId string, date *time.Time) ([]models.Task, []models.Task, error) {
	db, err := ConnectDb()
	if err != nil {
		return nil, nil, err
	}
	defer db.Close()
	var query string
	id, err := strconv.Atoi(memberId)
	if err != nil {
		return nil, nil, utils.InvalidRequestPayloadError
	}
	args := []interface{}{id, "true"}
	if date != nil {
		query = "SELECT * FROM tasks WHERE id IN (SELECT task_id FROM member_tasks WHERE member_id = $1) AND status = $2 AND finished_at >= $3"
		args = append(args, date)
	} else {
		query = "SELECT * FROM tasks WHERE id IN (SELECT task_id FROM member_tasks WHERE member_id = $1) AND status = $2"
	}
	var tasksDone []models.Task
	err = db.Select(&tasksDone, query, args...)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, nil, utils.DatabaseQueryError
		}
	}
	args[1] = "false"
	var tasksToDo []models.Task
	err = db.Select(&tasksToDo, query, args...)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, nil, utils.DatabaseQueryError
		}
	}
	return tasksDone, tasksToDo, nil
}
