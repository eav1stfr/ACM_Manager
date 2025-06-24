package models

import "time"

type Task struct {
	TaskID      int       `json:"task_id"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	Complexity  int       `json:"complexity"`
	Status      bool      `json:"status"`
	Assigned    bool      `json:"assigned"`
	AssignedTo  int       `json:"assigned_to,omitempty"`
}
