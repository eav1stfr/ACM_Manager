package sqlconnect

import (
	"acmmanager/internal/models"
	"acmmanager/utils"
	"database/sql"
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

func CreateTaskDbHandler(task models.Task) (models.Task, error) {
	db, err := ConnectDb()
	if err != nil {
		return models.Task{}, err
	}
	defer db.Close()

	query := "INSERT INTO tasks (description, deadline, complexity) VALUES ($1, $2, $3)"
	_, err = db.Exec(query, task.Description, task.Deadline, task.Complexity)
	if err != nil {
		return models.Task{}, utils.DatabaseQueryError
	}
	return task, nil
}

func DeleteTaskDbHandler(idStr string) (int, error) {
	db, err := ConnectDb()
	if err != nil {
		return 0, utils.ConnectingToDatabaseError
	}
	defer db.Close()

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, utils.InvalidRequestPayloadError
	}

	query := "DELETE FROM tasks WHERE id = $1"
	res, err := db.Exec(query, id)
	if err != nil {
		return 0, utils.DatabaseQueryError
	}
	rowsAffected, err := res.RowsAffected()
	if rowsAffected == 0 {
		return 0, utils.UnitNotFoundError
	}
	return id, nil
}

func AssignTaskDbHandler(taskIdStr, memberIdStr string) ([]interface{}, error) {
	taskID, err := strconv.Atoi(taskIdStr)
	if err != nil {
		return nil, utils.InvalidRequestPayloadError
	}
	memberID, err := strconv.Atoi(memberIdStr)
	if err != nil {
		return nil, utils.InvalidRequestPayloadError
	}
	res := make([]interface{}, 0, 2)
	db, err := ConnectDb()
	if err != nil {
		return nil, err
	}
	tx, err := db.Beginx()
	if err != nil {
		return nil, utils.StartingTransactionError
	}
	member, err := GetOneMemberByIdDbHandler(int64(memberID))
	if err != nil {
		return nil, err
	}
	queryUpdateAssignedStatus := "UPDATE tasks SET assigned = true WHERE id = $1 RETURNING *"
	queryCreateRelation := "INSERT INTO member_tasks (member_id, task_id) VALUES ($1, $2)"

	var updatedTask models.Task
	err = tx.Get(&updatedTask, queryUpdateAssignedStatus, taskID)
	if err != nil {
		tx.Rollback()
		if err == sql.ErrNoRows {
			return nil, utils.UnitNotFoundError
		}
		return nil, utils.DatabaseQueryError
	}
	_, err = tx.Exec(queryCreateRelation, memberID, taskID)
	if err != nil {
		tx.Rollback()
		return nil, utils.DatabaseQueryError
	}
	err = tx.Commit()
	if err != nil {
		return nil, utils.CommitingTransactionError
	}
	res = append(res, updatedTask)
	res = append(res, member)
	return res, nil
}

func MarkTaskAsDoneDbHandler(taskIdStr string) (models.Task, error) {
	taskID, err := strconv.Atoi(taskIdStr)
	if err != nil {
		return models.Task{}, utils.InvalidRequestPayloadError
	}
	db, err := ConnectDb()
	if err != nil {
		return models.Task{}, utils.ConnectingToDatabaseError
	}
	defer db.Close()
	var updatedTask models.Task
	query := "UPDATE tasks SET status = true, finished_at = NOW() WHERE id = $1 RETURNING *"
	err = db.Get(&updatedTask, query, taskID)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Task{}, utils.UnitNotFoundError
		}
		return models.Task{}, utils.DatabaseQueryError
	}
	return updatedTask, nil
}
