package handlers

import (
	"acmmanager/internal/models"
	"acmmanager/internal/sqlconnect"
	"acmmanager/utils"
	"encoding/json"
	"net/http"
	"time"
)

func CreateMeeting(w http.ResponseWriter, r *http.Request) {
	var meeting models.Meeting
	err := json.NewDecoder(r.Body).Decode(&meeting)
	if err != nil {
		http.Error(w, utils.InvalidRequestPayloadError.Error(), utils.InvalidRequestPayloadError.GetStatusCode())
		return
	}
	if meeting.Date.UTC().Before(time.Now().UTC()) {
		http.Error(w, "that time is passed!!", http.StatusBadRequest)
		return
	}
	err = utils.ValidateMeetingPost(meeting)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			http.Error(w, appErr.Error(), appErr.GetStatusCode())
			return
		}
		http.Error(w, utils.UnknownInternalServerError.Error(), utils.UnknownInternalServerError.GetStatusCode())
		return
	}
	createdMeeting, err := sqlconnect.CreateMeetingDbHandler(meeting)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			http.Error(w, appErr.Error(), appErr.GetStatusCode())
			return
		}
		http.Error(w, utils.UnknownInternalServerError.Error(), utils.UnknownInternalServerError.GetStatusCode())
		return
	}
	response := struct {
		Status string         `json:"status"`
		Data   models.Meeting `json:"data"`
	}{
		Status: "success",
		Data:   createdMeeting,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, utils.EncodingResponseError.Error(), utils.EncodingResponseError.GetStatusCode())
	}
}

func DeleteMeeting(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("meeting_id")
	if id == "" {
		http.Error(w, utils.InvalidRequestPayloadError.Error(), utils.InvalidRequestPayloadError.GetStatusCode())
		return
	}
	_, err := sqlconnect.DeleteMeetingDbHandler(id)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			http.Error(w, appErr.Error(), appErr.GetStatusCode())
			return
		}
		http.Error(w, utils.UnknownInternalServerError.Error(), utils.UnknownInternalServerError.GetStatusCode())
		return
	}
}

func GetMeetingsForWeek(w http.ResponseWriter, r *http.Request) {
	memberId := r.URL.Query().Get("member_id")
	meetings, err := sqlconnect.GetMeetingsForWeekDbHandler(memberId)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			http.Error(w, appErr.Error(), appErr.GetStatusCode())
			return
		}
		http.Error(w, utils.UnknownInternalServerError.Error(), utils.UnknownInternalServerError.GetStatusCode())
		return
	}
	response := struct {
		Status string           `json:"status"`
		Count  int              `json:"count"`
		Data   []models.Meeting `json:"data"`
	}{
		Status: "success",
		Count:  len(meetings),
		Data:   meetings,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, utils.EncodingResponseError.Error(), utils.EncodingResponseError.GetStatusCode())
	}
}

// get meetings that are for everyone - easy
// how to add meetings that are for certain deps
