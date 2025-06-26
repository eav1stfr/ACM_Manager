package utils

import "time"

func IsInNextWeek(t time.Time) bool {
	now := time.Now().UTC()
	weekFromNow := now.AddDate(0, 0, 7)
	return !t.Before(now) && !t.After(weekFromNow)
}
