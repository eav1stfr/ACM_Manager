package handlers

import (
	"acmmanager/internal/models"
	"acmmanager/internal/sqlconnect"
	"acmmanager/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func GenerateReportForMember(w http.ResponseWriter, r *http.Request) {
	var startDate *time.Time
	json.NewDecoder(r.Body).Decode(&startDate)

	memberId := r.URL.Query().Get("member_id")
	if memberId == "" {
		http.Error(w, utils.InvalidRequestPayloadError.Error(), utils.InvalidRequestPayloadError.GetStatusCode())
		return
	}

	member, tasksDone, tasksToDo, countAttended, countMissed, err := sqlconnect.GetMemberDataForReport(memberId, startDate)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			http.Error(w, appErr.Error(), appErr.GetStatusCode())
			return
		}
		http.Error(w, utils.InvalidRequestPayloadError.Error(), utils.InvalidRequestPayloadError.GetStatusCode())
		return
	}

	pdfBytes, err := utils.GeneratePDF(member, tasksDone, tasksToDo, countAttended, countMissed, startDate)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			http.Error(w, appErr.Error(), appErr.GetStatusCode())
			return
		}
		http.Error(w, utils.InvalidRequestPayloadError.Error(), utils.InvalidRequestPayloadError.GetStatusCode())
		return
	}
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=%s-%s.pdf", member.FirstName, member.LastName))
	_, err = w.Write(pdfBytes)
	if err != nil {
		http.Error(w, utils.UnknownInternalServerError.Error(), utils.UnknownInternalServerError.GetStatusCode())
	}
}

func GenerateReportForDepartment(w http.ResponseWriter, r *http.Request) {
	var startDate *time.Time
	json.NewDecoder(r.Body).Decode(&startDate)

	depId := r.URL.Query().Get("dep_id")
	if depId == "" {
		http.Error(w, utils.InvalidRequestPayloadError.Error(), utils.InvalidRequestPayloadError.GetStatusCode())
		return
	}

	members, err := sqlconnect.GetMembersDBHandler(depId, "")
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			http.Error(w, appErr.Error(), appErr.GetStatusCode())
			return
		}
		http.Error(w, utils.InvalidRequestPayloadError.Error(), utils.InvalidRequestPayloadError.GetStatusCode())
		return
	}
	var membersInfoForReport []models.MemberWithData
	for _, member := range members {
		memberId := strconv.Itoa(int(member.ID))
		_, tasksDone, tasksToDo, countAttended, countMissed, err := sqlconnect.GetMemberDataForReport(memberId, startDate)
		if err != nil {
			if appErr, ok := err.(*utils.AppError); ok {
				http.Error(w, appErr.Error(), appErr.GetStatusCode())
				return
			}
			http.Error(w, utils.InvalidRequestPayloadError.Error(), utils.InvalidRequestPayloadError.GetStatusCode())
			return
		}
		memberInfo := models.MemberWithData{
			FirstName:     member.FirstName,
			LastName:      member.LastName,
			TasksDone:     tasksDone,
			TasksToDo:     tasksToDo,
			CountAttended: countAttended,
			CountMissed:   countMissed,
		}
		membersInfoForReport = append(membersInfoForReport, memberInfo)
	}
	pdfBytes, err := utils.GeneratePDFForDepartment(membersInfoForReport, depId, startDate)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			http.Error(w, appErr.Error(), appErr.GetStatusCode())
			return
		}
		http.Error(w, utils.InvalidRequestPayloadError.Error(), utils.InvalidRequestPayloadError.GetStatusCode())
		return
	}
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=%s-report", depId))
	_, err = w.Write(pdfBytes)
	if err != nil {
		http.Error(w, utils.UnknownInternalServerError.Error(), utils.UnknownInternalServerError.GetStatusCode())
	}
}
