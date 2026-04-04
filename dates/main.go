package main

func DaysBetweenDates(date1 string, date2 string) int {
	// time: O(n), space: O(1)
	// n - year2-year1

	if date1 > date2 {
		date1, date2 = date2, date1
	}
	year1, month1, day1 := StrToDate(date1)
	year2, month2, day2 := StrToDate(date2)
	if year1 == year2 {
		return YearDay(year2, month2, day2) - YearDay(year1, month1, day1)
	} else {
		dayCount := YearDays(year1) - YearDay(year1, month1, day1)
		for y := year1 + 1; y < year2; y++ {
			dayCount += YearDays(y)
		}
		return dayCount + YearDay(year2, month2, day2)
	}
}

func YearDay(year int, month int, day int) int {
	// time: O(n), space: O(1)
	// n - month

	dayCount := 0
	for m := 1; m < month; m++ {
		dayCount += MonthDays(m, year)
	}
	return dayCount + day
}

func MonthDays(month int, year int) int {
	// time: O(1), space: O(1)

	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2:
		if IsLeap(year) {
			return 29
		}
		return 28
	}
	return 0
}

func YearDays(year int) int {
	// time: O(1), space: O(1)

	if IsLeap(year) {
		return 366
	}
	return 365
}

func IsLeap(year int) bool {
	// time: O(1), space: O(1)

	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func StrToDate(s string) (year int, month int, day int) {
	// format: "YYYY*MM*DD" (* - any letter)

	return strToInt(s[:4]), strToInt(s[5:7]), strToInt(s[8:])
}

func strToInt(s string) int {
	// time: O(n), space: O(1)

	res := 0
	for i := range s {
		res = res*10 + int(s[i]-'0')
	}
	return res
}

func WeekDay(year int, month int, day int) int {
	// time: O(n), space: O(1)
	// n - year-1971
	// from 0 to 6

	const startYear = 1971
	const startWeekDay = 4 // friday
	dayCount := 0
	for y := startYear; y < year; y++ {
		dayCount += YearDays(y)
	}
	dayCount += YearDay(year, month, day)
	return (dayCount - startWeekDay) % 7
}

func WeekDayToStr(day int) string {
	// time: O(1), space: O(1)

	switch day {
	case 0:
		return "Monday"
	case 1:
		return "Tuesday"
	case 2:
		return "Wednesday"
	case 3:
		return "Thursday"
	case 4:
		return "Friday"
	case 5:
		return "Saturday"
	case 6:
		return "Sunday"
	}
	return ""
}
