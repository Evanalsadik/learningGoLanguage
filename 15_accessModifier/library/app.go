package library

const secondInMinute int = 60
const minuteInHour int = 60
const HourInDay int = 24

func DayToHour(day int) int {
	return day * HourInDay
}

func hourToMinute(hour int) int {
	return hour * minuteInHour
}
