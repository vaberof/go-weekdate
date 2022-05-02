package main

import (
	"fmt"
	"time"

	"github.com/vaberof/goweekdate"
)

func Example() {
	// Create new WeekDate object with specified startWeek and location.
	weekDate := weekdate.New(time.Now(), "Asia/Novosibirsk")

	// Get names of the week days.
	weekDays := weekDate.WeekDays()

	// Get dates formatted as ("02.01") on 2 weeks in a row.
	// 'week' is 2, 'include' is true.
	shortDates := weekDate.ShortDates(2, true)

	// Get dates formatted as ("02.01.2006") on the second week.
	// 'week' is 2, 'include' is false.
	fullDates := weekDate.FullDates(2, false)

	// Get days and dates of the type of map[Monday:02.05.2022, Tuesday:03.05.2022...]
	daysAndDates := weekDate.DaysAndDates()

	fmt.Println(weekDays)     // [Monday Tuesday Wednesday Thursday Friday Saturday Sunday]
	fmt.Println(shortDates)   // [02.05 03.05 04.05 05.05 06.05 07.05 08.05 09.05 10.05 11.05 12.05 13.05 14.05 15.05]
	fmt.Println(fullDates)    // [09.05.2022 10.05.2022 11.05.2022 12.05.2022 13.05.2022 14.05.2022 15.05.2022]
	fmt.Println(daysAndDates) // map[Friday:06.05.2022 Monday:02.05.2022 Saturday:07.05.2022 Sunday:08.05.2022 Thursday:05.05.2022 Tuesday:03.05.2022 Wednesday:04.05.2022]
}

func main() {
	Example()
}
