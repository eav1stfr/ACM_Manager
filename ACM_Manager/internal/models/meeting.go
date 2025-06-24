package models

import "time"

type Meeting struct {
	ID            int          `json:"meeting_id"`
	Venue         string       `json:"venue"`
	Time          time.Time    `json:"time"`
	Day           time.Weekday `json:"day"`
	Repeated      bool         `json:"repeated"`
	ForDepartment bool         `json:"for_dep"`
	DepartmentID  *int         `json:"department_id,omitempty"`
}
