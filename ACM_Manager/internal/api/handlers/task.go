package handlers

import (
	"acmmanager/internal/models"
	"acmmanager/internal/sqlconnect"
	"acmmanager/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	status := r.URL.Query().Get("status")
	tasks, err := sqlconnect.GetTasksDbHandler(id, status)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			http.Error(w, appErr.Error(), appErr.GetStatusCode())
			return
		}
		http.Error(w, utils.UnknownInternalServerError.Error(), utils.UnknownInternalServerError.GetStatusCode())
		return
	}
	response := struct {
		Status string        `json:"status"`
		Count  int           `json:"count"`
		Data   []models.Task `json:"data"`
	}{
		Status: "success",
		Count:  len(tasks),
		Data:   tasks,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, utils.EncodingResponseError.Error(), utils.EncodingResponseError.GetStatusCode())
	}
}

func CreateTasks(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, utils.InvalidRequestPayloadError.Error(), utils.InvalidRequestPayloadError.GetStatusCode())
		return
	}
	err = utils.ValidateTaskPost(task)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			http.Error(w, appErr.Error(), appErr.GetStatusCode())
			return
		}
		http.Error(w, utils.UnknownInternalServerError.Error(), utils.UnknownInternalServerError.GetStatusCode())
		return
	}
	_, err = sqlconnect.CreateTaskDbHandler(task)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			http.Error(w, appErr.Error(), appErr.GetStatusCode())
			return
		}
		http.Error(w, utils.UnknownInternalServerError.Error(), utils.UnknownInternalServerError.GetStatusCode())
		return
	}
	response := struct {
		Status string      `json:"status"`
		Data   models.Task `json:"data"`
	}{
		Status: "success",
		Data:   task,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
}

func AssignTask(w http.ResponseWriter, r *http.Request) {
	taskIdStr := r.URL.Query().Get("task_id")
	memberIdStr := r.URL.Query().Get("member_id")
	if taskIdStr == "" || memberIdStr == "" {
		http.Error(w, utils.InvalidRequestPayloadError.Error(), utils.InvalidRequestPayloadError.GetStatusCode())
		return
	}
	res, err := sqlconnect.AssignTaskDbHandler(taskIdStr, memberIdStr)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			http.Error(w, appErr.Error(), appErr.GetStatusCode())
			return
		}
		http.Error(w, utils.UnknownInternalServerError.Error(), utils.UnknownInternalServerError.GetStatusCode())
		return
	}
	fmt.Println(res)
}

func MarkTaskAsDone(w http.ResponseWriter, r *http.Request) {
	taskId := r.URL.Query().Get("task_id")
	if taskId == "" {
		http.Error(w, utils.InvalidRequestPayloadError.Error(), utils.InvalidRequestPayloadError.GetStatusCode())
		return
	}
	task, err := sqlconnect.MarkTaskAsDoneDbHandler(taskId)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			http.Error(w, appErr.Error(), appErr.GetStatusCode())
			return
		}
		http.Error(w, utils.UnknownInternalServerError.Error(), utils.UnknownInternalServerError.GetStatusCode())
		return
	}
	response := struct {
		Status string      `json:"status"`
		Data   models.Task `json:"data"`
	}{
		Status: "successfully updated",
		Data:   task,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, utils.EncodingResponseError.Error(), utils.EncodingResponseError.GetStatusCode())
	}
}

func DeleteTasks(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	_, err := sqlconnect.DeleteTaskDbHandler(idStr)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			http.Error(w, appErr.Error(), appErr.GetStatusCode())
			return
		}
		http.Error(w, utils.UnknownInternalServerError.Error(), utils.UnknownInternalServerError.GetStatusCode())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
