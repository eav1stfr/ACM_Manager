package sqlconnect

import (
	"acmmanager/internal/models"
	"acmmanager/utils"
	"fmt"
	"strconv"
)

func GetTasksDbHandler(idStr string, status string) ([]models.Task, error) {
	db, err := ConnectDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var tasks []models.Task
	if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return nil, utils.InvalidRequestPayloadError
		}
		query := "SELECT * FROM tasks WHERE id IN (SELECT task_id FROM member_tasks WHERE member_id = $1)"
		if status != "" {
			query += " AND status = " + status
			fmt.Println(query)
		}
		err = db.Select(&tasks, query, id)
		if err != nil {
			return nil, utils.DatabaseQueryError
		}
	} else if status != "" {
		query := "SELECT * FROM tasks WHERE status = $1"
		err = db.Select(&tasks, query, status)
		if err != nil {
			return nil, utils.DatabaseQueryError
		}
	} else {
		query := "SELECT * FROM tasks"
		err = db.Select(&tasks, query)
		if err != nil {
			return nil, utils.DatabaseQueryError
		}
	}
	return tasks, nil
}
