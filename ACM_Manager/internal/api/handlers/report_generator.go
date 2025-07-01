package handlers

import (
	"acmmanager/internal/models"
	"acmmanager/internal/sqlconnect"
	"acmmanager/utils"
	"github.com/phpdave11/gofpdf"
	"net/http"
	"strconv"
)

func GenerateReportForMember(w http.ResponseWriter, r *http.Request) {
	memberId := r.URL.Query().Get("member_id")
	if memberId == "" {
		http.Error(w, utils.InvalidRequestPayloadError.Error(), utils.InvalidRequestPayloadError.GetStatusCode())
		return
	}

	member, tasksDone, tasksToDo, countAttended, countMissed, err := sqlconnect.GetMemberDataForReport(memberId)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			http.Error(w, appErr.Error(), appErr.GetStatusCode())
			return
		}
		http.Error(w, utils.InvalidRequestPayloadError.Error(), utils.InvalidRequestPayloadError.GetStatusCode())
		return
	}

	pdfBytes, err := utils.GeneratePDF(member, tasksDone, tasksToDo, countAttended, countMissed)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			http.Error(w, appErr.Error(), appErr.GetStatusCode())
			return
		}
		http.Error(w, utils.InvalidRequestPayloadError.Error(), utils.InvalidRequestPayloadError.GetStatusCode())
		return
	}
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "inline; filename=from-function.pdf")
	_, err = w.Write(pdfBytes)
	if err != nil {
		http.Error(w, utils.UnknownInternalServerError.Error(), utils.UnknownInternalServerError.GetStatusCode())
	}
}

func GenerateReportForDepartment(w http.ResponseWriter, r *http.Request) {
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
		_, tasksDone, tasksToDo, countAttended, countMissed, err := sqlconnect.GetMemberDataForReport(memberId)
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
	pdfBytes, err := utils.GeneratePDFForDepartment(membersInfoForReport, depId)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			http.Error(w, appErr.Error(), appErr.GetStatusCode())
			return
		}
		http.Error(w, utils.InvalidRequestPayloadError.Error(), utils.InvalidRequestPayloadError.GetStatusCode())
		return
	}
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "inline; filename=from-function.pdf")
	_, err = w.Write(pdfBytes)
	if err != nil {
		http.Error(w, utils.UnknownInternalServerError.Error(), utils.UnknownInternalServerError.GetStatusCode())
	}
}

func PDFHandler(w http.ResponseWriter, r *http.Request) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, PDF from Go HTTP server!")

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "inline; filename=example.pdf")
	err := pdf.Output(w)
	if err != nil {
		http.Error(w, "failed to generate PDF", http.StatusInternalServerError)
		return
	}
}
