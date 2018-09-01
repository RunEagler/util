package util

import "time"

func Date(t time.Time) time.Time {

	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
}

func Datetime(t time.Time) time.Time {

	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, time.UTC)
}

func DateString(t time.Time) string {
	return t.Format("2006-01-02")
}

func DatetimeString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
