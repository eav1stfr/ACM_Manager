package models

import "time"

type Meeting struct {
	ID           int       `db:"id" json:"meeting_id"`
	Venue        string    `db:"venue" json:"venue" validate:"required"`
	Date         time.Time `db:"time" json:"date" validate:"required"`
	Repeated     bool      `db:"repeated" json:"repeated"`
	DepartmentID *string   `db:"department_id" json:"department_id,omitempty"`
}
