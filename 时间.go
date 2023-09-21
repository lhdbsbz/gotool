package gotool

import "time"

func GetYear(someTime time.Time) string {
	return someTime.Format("2006")
}

func GetMonth(someTime time.Time) string {
	return someTime.Format("01")
}

func GetDay(someTime time.Time) string {
	return someTime.Format("02")
}

func GetHour(someTime time.Time) string {
	return someTime.Format("15")
}

func GetMin(someTime time.Time) string {
	return someTime.Format("04")
}

func GetSecond(someTime time.Time) string {
	return someTime.Format("05")
}
