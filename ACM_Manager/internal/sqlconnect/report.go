package sqlconnect

import (
	"acmmanager/internal/models"
	"strconv"
)

//
//func GenerateReportForMemberDbHandler(memberId string) (map[string]interface{}, error) {
//	// connect db
//	db, err :=
//	// get member info
//
//	// get list of all tasks for a member
//
//	// get attendance report : meeting info : skipped --> then get the statistics (missed / total)
//
//}

func GetMemberDataForReport(memberId string) (models.Member, []models.Task, []models.Task, int, int, error) {
	id, err := strconv.Atoi(memberId)
	if err != nil {
		return models.Member{}, nil, nil, 0, 0, err
	}
	member, err := GetOneMemberByIdDbHandler(int64(id))
	if err != nil {
		return models.Member{}, nil, nil, 0, 0, err
	}
	tasksDone, err := GetTasksDbHandler(memberId, "true")
	if err != nil {
		return models.Member{}, nil, nil, 0, 0, err
	}
	tasksToDo, err := GetTasksDbHandler(memberId, "false")
	if err != nil {
		return models.Member{}, nil, nil, 0, 0, err
	}
	countAttended, countMissed, err := GetAttendanceCount(memberId)
	if err != nil {
		return models.Member{}, nil, nil, 0, 0, err
	}
	return member, tasksDone, tasksToDo, countAttended, countMissed, nil
}
