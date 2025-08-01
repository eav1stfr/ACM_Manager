package handlers

import (
	"acmmanager/internal/models"
	"acmmanager/internal/sqlconnect"
	"acmmanager/utils"
	"encoding/json"
	"net/http"
)

func GetMembersHandler(w http.ResponseWriter, r *http.Request) {
	dep := r.URL.Query().Get("department")
	id := r.URL.Query().Get("id")
	members, err := sqlconnect.GetMembersDBHandler(dep, id)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			http.Error(w, appErr.Error(), appErr.GetStatusCode())
			return
		}
		http.Error(w, utils.UnknownInternalServerError.Error(), utils.UnknownInternalServerError.GetStatusCode())
	}
	w.Header().Set("Content-Type", "application/json")
	response := struct {
		Status string          `json:"status"`
		Count  int             `json:"count"`
		Data   []models.Member `json:"data"`
	}{
		Status: "success",
		Count:  len(members),
		Data:   members,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, utils.EncodingResponseError.Error(), utils.EncodingResponseError.GetStatusCode())
		return
	}
}

func CreateMemberHandler(w http.ResponseWriter, r *http.Request) {
	var newMember models.Member
	err := json.NewDecoder(r.Body).Decode(&newMember)
	if err != nil {
		http.Error(w, utils.InvalidRequestPayloadError.Error(), utils.InvalidRequestPayloadError.GetStatusCode())
		return
	}
	if err = utils.ValidatePostMember(newMember); err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			http.Error(w, appErr.Error(), appErr.GetStatusCode())
			return
		}
	}

	err = sqlconnect.PostMemberDbHandler(newMember)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			http.Error(w, appErr.Error(), appErr.GetStatusCode())
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := struct {
		Status string        `json:"status"`
		Data   models.Member `json:"data"`
	}{
		Status: "success",
		Data:   newMember,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, utils.EncodingResponseError.Error(), utils.EncodingResponseError.GetStatusCode())
		return
	}
}

func CreateMembersHandler(w http.ResponseWriter, r *http.Request) {
	var newMembers []models.Member
	err := json.NewDecoder(r.Body).Decode(&newMembers)
	if err != nil {
		http.Error(w, utils.InvalidRequestPayloadError.Error(), utils.InvalidRequestPayloadError.GetStatusCode())
		return
	}

	if err = utils.ValidateMemberPost(newMembers); err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			http.Error(w, appErr.Error(), appErr.GetStatusCode())
			return
		}
		http.Error(w, utils.UnknownInternalServerError.Error(), utils.UnknownInternalServerError.GetStatusCode())
		return
	}

	addedMembers, err := sqlconnect.PostMembersDBHandler(newMembers)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			http.Error(w, appErr.Error(), appErr.GetStatusCode())
			return
		}
		http.Error(w, utils.UnknownInternalServerError.Error(), utils.UnknownInternalServerError.GetStatusCode())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := struct {
		Status string          `json:"status"`
		Count  int             `json:"count"`
		Data   []models.Member `json:"data"`
	}{
		Status: "success",
		Count:  len(addedMembers),
		Data:   addedMembers,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, utils.EncodingResponseError.Error(), utils.EncodingResponseError.GetStatusCode())
		return
	}
}

func DeleteMembersHandler(w http.ResponseWriter, r *http.Request) {
	var ids []int64
	err := json.NewDecoder(r.Body).Decode(&ids)
	if err != nil {
		http.Error(w, utils.InvalidRequestPayloadError.Error(), utils.InvalidRequestPayloadError.GetStatusCode())
		return
	}
	deletedIds, err := sqlconnect.DeleteMembersHandler(ids)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			http.Error(w, appErr.Error(), appErr.GetStatusCode())
			return
		}
		http.Error(w, utils.UnknownInternalServerError.Error(), utils.UnknownInternalServerError.GetStatusCode())
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response := struct {
		Status     string  `json:"status"`
		DeletedIDs []int64 `json:"deleted_ids"`
	}{
		Status:     "Student(s) successfully deleted",
		DeletedIDs: deletedIds,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, utils.EncodingResponseError.Error(), utils.EncodingResponseError.GetStatusCode())
	}
}

func DeleteAllMembers(w http.ResponseWriter, r *http.Request) {
	err := sqlconnect.DeleteAllMembersDbHandler()
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			http.Error(w, appErr.Error(), appErr.GetStatusCode())
			return
		}
	}

}

func PatchMembersHandler(w http.ResponseWriter, r *http.Request) {
	var updates []map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		http.Error(w, utils.InvalidRequestPayloadError.Error(), utils.InvalidRequestPayloadError.GetStatusCode())
		return
	}
	err = sqlconnect.PatchMembersDbHandler(updates)
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
