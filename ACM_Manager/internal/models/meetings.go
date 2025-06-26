package models

import "time"

type Meeting struct {
	ID           int       `db:"id" json:"meeting_id"`
	Venue        string    `db:"venue" json:"venue" validate:"required"`
	Date         time.Time `db:"time" json:"date" validate:"required"`
	Repeated     bool      `db:"repeated" json:"repeated"`
	DepartmentID *string   `db:"department_id" json:"department_id,omitempty"`
}

//CREATE TABLE meetings (
//	id SERIAL PRIMARY KEY,
//	venue TEXT NOT NULL,
//	time TIMESTAMP NOT NULL,
//	repeated BOOLEAN DEFAULT FALSE,
//	department_id TEXT REFERENCES departments(name_of_dep) ON DELETE SET NULL
//);
