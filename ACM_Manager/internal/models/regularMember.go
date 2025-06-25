package models

import "time"

type Member struct {
	ID        int64     `db:"id" json:"id"`
	FirstName string    `db:"first_name" json:"first_name"`
	LastName  string    `db:"last_name" json:"last_name"`
	Email     string    `db:"email" json:"email"`
	Telegram  string    `db:"telegram" json:"telegram"`
	Role      string    `db:"role" json:"role"`
	Birthday  time.Time `db:"birthday" json:"birthday"`
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
