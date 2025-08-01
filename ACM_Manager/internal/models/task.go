package models

import "time"

type Task struct {
	TaskID      int        `db:"id" json:"task_id"`
	Description string     `db:"description" json:"description" validate:"required"`
	Deadline    time.Time  `db:"deadline" json:"deadline" validate:"required"`
	Complexity  int        `db:"complexity" json:"complexity" validate:"required"`
	Status      bool       `db:"status" json:"status"`
	Assigned    bool       `db:"assigned" json:"assigned"`
	FinishedAt  *time.Time `db:"finished_at,omitempty" json:"finished_at"`
}
