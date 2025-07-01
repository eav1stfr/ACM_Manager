package models

type MemberWithData struct {
	FirstName     string
	LastName      string
	TasksDone     []Task
	TasksToDo     []Task
	CountAttended int
	CountMissed   int
}
