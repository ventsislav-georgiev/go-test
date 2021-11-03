package main

import (
	"strconv"
)

var dateTime = "26 10 2021"

func dateWithSlice() (int, int, int) {
	day := dateTime[0:2]
	month := dateTime[3:5]
	year := dateTime[6:10]

	dayInt, _ := strconv.Atoi(day)
	monthInt, _ := strconv.Atoi(month)
	yearInt, _ := strconv.Atoi(year)

	return dayInt, monthInt, yearInt
}
