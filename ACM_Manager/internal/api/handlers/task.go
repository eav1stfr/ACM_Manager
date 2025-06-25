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
	fmt.Println("ID IS", id, "AND STATUS IS", status)
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
		Status string
		Count  int
		Data   []models.Task
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

}

func UpdateTask(w http.ResponseWriter, r *http.Request) {

}

func DeleteTasks(w http.ResponseWriter, r *http.Request) {

}
