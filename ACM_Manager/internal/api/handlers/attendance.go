package handlers

import (
	"acmmanager/internal/sqlconnect"
	"acmmanager/utils"
	"encoding/json"
	"net/http"
)

func MarkAttendance(w http.ResponseWriter, r *http.Request) {
	meetingId := r.URL.Query().Get("meeting_id")
	memberId := r.URL.Query().Get("member_id")
	if meetingId == "" || memberId == "" {
		http.Error(w, utils.MissingRequiredParametersError.Error(), utils.MissingRequiredParametersError.GetStatusCode())
		return
	}
	err := sqlconnect.MarkAttendanceDbHandler(memberId, meetingId)
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

func GetReasonOfAbsence(w http.ResponseWriter, r *http.Request) {
	meetingId := r.URL.Query().Get("meeting_id")
	memberId := r.URL.Query().Get("member_id")
	if meetingId == "" || memberId == "" {
		http.Error(w, utils.MissingRequiredParametersError.Error(), utils.MissingRequiredParametersError.GetStatusCode())
		return
	}
	var comment map[string]string
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, utils.InvalidRequestPayloadError.Error(), utils.InvalidRequestPayloadError.GetStatusCode())
		return
	}
	comm, exists := comment["comment"]
	if !exists {
		http.Error(w, utils.InvalidRequestPayloadError.Error(), utils.InvalidRequestPayloadError.GetStatusCode())
		return
	}
	err = sqlconnect.GetReasonOfAbsence(memberId, meetingId, comm)
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
