package models

import "time"

type Task struct {
	TaskID      int       `db:"id" json:"task_id"`
	Description string    `db:"description" json:"description"`
	Deadline    time.Time `db:"deadline" json:"deadline"`
	Complexity  int       `db:"complexity" json:"complexity"`
	Status      bool      `db:"status" json:"status"`
	Assigned    bool      `db:"assigned" json:"assigned"`
}

//CREATE TABLE tasks (
//	id SERIAL PRIMARY KEY,
//	description TEXT NOT NULL,
//	deadline TIMESTAMP NOT NULL,
//	complexity INT NOT NULL,
//	status BOOLEAN DEFAULT FALSE,
//	assigned BOOLEAN DEFAULT FALSE
//);
