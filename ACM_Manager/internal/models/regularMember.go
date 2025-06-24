package models

type RegularMember struct {
	ID            int    `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	Telegram      string `json:"telegram"`
	Role          string `json:"role"`
	DepartmentIDs []int  `json:"department_ids"`
	Tasks         []int  `json:"tasks"`
}
