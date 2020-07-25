package main

import (
	"time"
)

// Friday 13th or Black Friday is considered as unlucky day.
// Calculate how many unlucky days are in the given year.
// Find the number of Friday 13th in the given year.
// Input: Year as an integer.

func unluckyDays(year int) int {
	counter := 0
	for m := 1; m <= 12; m++ {
		if time.Date(year, time.Month(m), 13, 0, 0, 0, 0, time.UTC).Weekday() == 5 {
			counter++
		}
	}
	return counter
}
