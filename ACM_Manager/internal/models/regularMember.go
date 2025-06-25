package models

import "time"

type Member struct {
	ID        int64     `db:"id" json:"id" validate:"required"`
	FirstName string    `db:"first_name" json:"first_name" validate:"required"`
	LastName  string    `db:"last_name" json:"last_name" validate:"required"`
	Email     string    `db:"email" json:"email" validate:"required,email"`
	Telegram  string    `db:"telegram" json:"telegram" validate:"required"`
	Role      string    `db:"role" json:"role" validate:"required"`
	Birthday  time.Time `db:"birthday" json:"birthday" validate:"required"`
}

//CREATE TABLE members (
//	id BIGINT PRIMARY KEY,
//	first_name TEXT NOT NULL,
//	last_name TEXT NOT NULL,
//	email TEXT UNIQUE NOT NULL,
//	telegram TEXT,
//	role TEXT NOT NULL,
//	birthday DATE NOT NULL
//);
