package models

type BoardMember struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Telegram  string `json:"telegram"`
	Role      string `json:"role"`
}
