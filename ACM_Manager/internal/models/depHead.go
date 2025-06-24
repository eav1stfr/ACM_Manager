package models

type DepHead struct {
	HeadID    int    `json:"head_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Telegram  string `json:"telegram"`
	Role      string `json:"role"`
	DepID     int    `json:"department_id"`
}
