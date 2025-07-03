package utils

import (
	"acmmanager/internal/models"
	"bytes"
	"fmt"
	"github.com/phpdave11/gofpdf"
	"strings"
	"time"
)

func GeneratePDF(member models.Member, tasksDone, tasksToDo []models.Task, countAttended, countMissed int, date *time.Time) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	marginLeft := 20.0
	pdf.SetLeftMargin(marginLeft)
	pdf.SetRightMargin(20)
	lineHeight := 8.0

	// Title
	pdf.SetFont("Arial", "B", 20)
	pdf.CellFormat(0, 12, "ACM@NU Member Report", "", 1, "C", false, 0, "")
	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(0, 10, fmt.Sprintf("%s %s", member.FirstName, member.LastName), "", 1, "C", false, 0, "")

	pdf.SetFont("Arial", "", 12)
	if date == nil {
		//pdf.CellFormat(0, 8, "10.08.2025 - 11.08.2025", "", 1, "C", false, 0, "")
		pdf.CellFormat(0, 8, fmt.Sprintf("01.09.2025 - %s", time.Now().Format("02.01.2006")), "", 1, "C", false, 0, "")

	} else {
		pdf.CellFormat(0, 8, fmt.Sprintf("%s - %s", date.Format("02.01.2006"), time.Now().Format("02.01.2006")), "", 1, "C", false, 0, "")
	}
	pdf.Ln(5)

	// Draw line
	pdf.SetDrawColor(100, 100, 100)
	pdf.Line(marginLeft, pdf.GetY(), 210-marginLeft, pdf.GetY())
	pdf.Ln(5)

	// Member section helper
	writeMemberSection := func(name string, tasksDone []string, tasksTodo []string, meetingSummary string) {
		pdf.SetFont("Arial", "B", 14)
		pdf.SetTextColor(0, 102, 204)
		pdf.CellFormat(0, 10, name, "", 1, "", false, 0, "")
		pdf.SetTextColor(0, 0, 0)

		// Tasks done
		pdf.SetFont("Arial", "B", 12)
		pdf.CellFormat(0, lineHeight, "Tasks completed:", "", 1, "", false, 0, "")
		pdf.SetFont("Arial", "", 12)
		if len(tasksDone) == 0 {
			pdf.CellFormat(0, lineHeight, "- None", "", 1, "", false, 0, "")
		} else {
			for i, task := range tasksDone {
				pdf.CellFormat(0, lineHeight, formatTask(i+1, task), "", 1, "", false, 0, "")
			}
		}

		// Tasks to do
		pdf.SetFont("Arial", "B", 12)
		pdf.CellFormat(0, lineHeight, "Tasks to do:", "", 1, "", false, 0, "")
		pdf.SetFont("Arial", "", 12)
		if len(tasksTodo) == 0 {
			pdf.CellFormat(0, lineHeight, "- None", "", 1, "", false, 0, "")
		} else {
			for i, task := range tasksTodo {
				pdf.CellFormat(0, lineHeight, formatTask(i+1, task), "", 1, "", false, 0, "")
			}
		}

		// Meetings
		pdf.SetFont("Arial", "B", 12)
		pdf.CellFormat(0, lineHeight, "Meeting Summary:", "", 1, "", false, 0, "")
		pdf.SetFont("Arial", "", 12)
		pdf.MultiCell(0, lineHeight, meetingSummary, "", "", false)

		pdf.Ln(4)
		pdf.SetDrawColor(200, 200, 200)
		pdf.Line(marginLeft, pdf.GetY(), 210-marginLeft, pdf.GetY())
		pdf.Ln(6)
	}

	tasksDoneList := make([]string, 0, len(tasksDone))
	for _, task := range tasksDone {
		tasksDoneList = append(tasksDoneList, task.Description)
	}
	tasksToDoList := make([]string, 0, len(tasksToDo))
	for _, task := range tasksToDo {
		taskInfo := fmt.Sprintf("%s with deadline %s", task.Description, task.Deadline)
		tasksToDoList = append(tasksToDoList, taskInfo)
	}
	// Add members
	writeMemberSection(
		fmt.Sprintf("%s %s", member.FirstName, member.LastName),
		tasksDoneList,
		tasksToDoList,
		fmt.Sprintf("%d meetings attended, %d meetings missed", countAttended, countMissed),
	)

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func formatTask(i int, text string) string {
	return string(rune('0'+i)) + ". " + text
}

func GeneratePDFForDepartment(members []models.MemberWithData, depId string, date *time.Time) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetAutoPageBreak(true, 15)
	pdf.AddPage()

	marginLeft := 20.0
	pdf.SetLeftMargin(marginLeft)
	pdf.SetRightMargin(20)
	lineHeight := 8.0

	// Department Title
	pdf.SetFont("Arial", "B", 20)
	pdf.CellFormat(0, 12, "ACM@NU Department Report", "", 1, "C", false, 0, "")

	pdf.SetFont("Arial", "", 20)
	pdf.CellFormat(0, 12, strings.ToUpper(depId), "", 1, "C", false, 0, "")

	pdf.SetFont("Arial", "", 12)
	if date == nil {
		pdf.CellFormat(0, 8, fmt.Sprintf("01.09.2025 - %s", time.Now().Format("02.01.2006")), "", 1, "C", false, 0, "")

	} else {
		pdf.CellFormat(0, 8, fmt.Sprintf("%s - %s", date.Format("02.01.2006"), time.Now().Format("02.01.2006")), "", 1, "C", false, 0, "")
	}
	pdf.Ln(5)

	pdf.SetDrawColor(100, 100, 100)
	pdf.Line(marginLeft, pdf.GetY(), 210-marginLeft, pdf.GetY())
	pdf.Ln(5)

	// Section helper
	writeMemberSection := func(m models.MemberWithData) {
		if pdf.GetY() > 260 {
			pdf.AddPage()
		}

		pdf.SetFont("Arial", "B", 14)
		pdf.SetTextColor(0, 102, 204)
		pdf.CellFormat(0, 10, fmt.Sprintf("%s %s", m.FirstName, m.LastName), "", 1, "", false, 0, "")
		pdf.SetTextColor(0, 0, 0)

		// Done
		pdf.SetFont("Arial", "B", 12)
		pdf.CellFormat(0, lineHeight, "Tasks completed:", "", 1, "", false, 0, "")
		pdf.SetFont("Arial", "", 12)
		if len(m.TasksDone) == 0 {
			pdf.CellFormat(0, lineHeight, "- None", "", 1, "", false, 0, "")
		} else {
			for i, t := range m.TasksDone {
				text := fmt.Sprintf("%d. %s", i+1, t.Description)
				pdf.CellFormat(0, lineHeight, text, "", 1, "", false, 0, "")
			}
		}

		// To do
		pdf.SetFont("Arial", "B", 12)
		pdf.CellFormat(0, lineHeight, "Tasks to do:", "", 1, "", false, 0, "")
		pdf.SetFont("Arial", "", 12)
		if len(m.TasksToDo) == 0 {
			pdf.CellFormat(0, lineHeight, "- None", "", 1, "", false, 0, "")
		} else {
			for i, t := range m.TasksToDo {
				text := fmt.Sprintf("%d. %s (deadline: %s)", i+1, t.Description, t.Deadline.Format("02.01.2006"))
				pdf.CellFormat(0, lineHeight, text, "", 1, "", false, 0, "")
			}
		}

		// Meeting summary
		pdf.SetFont("Arial", "B", 12)
		pdf.CellFormat(0, lineHeight, "Meeting Summary:", "", 1, "", false, 0, "")
		pdf.SetFont("Arial", "", 12)
		pdf.MultiCell(0, lineHeight,
			fmt.Sprintf("%d meetings attended, %d missed", m.CountAttended, m.CountMissed),
			"", "", false,
		)

		pdf.Ln(4)
		pdf.SetDrawColor(200, 200, 200)
		pdf.Line(marginLeft, pdf.GetY(), 210-marginLeft, pdf.GetY())
		pdf.Ln(6)
	}

	for _, member := range members {
		writeMemberSection(member)
	}

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
