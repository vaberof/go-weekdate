# goweekdate

`goweekdate` is a simple go package that allows you to easily get dates of the week days in range that you need.

## Installation

    go get github.com/vaberof/goweekdate

## Initialization

```go
weekDate := weekdate.New(time.Now(), "Asia/Novosibirsk")
```

With setting first parameter `weekStart` as _time.Now()_ calculations will be making from monday of current week. You can
modify it by using _time.Now().Add()_ method to make calculations from previous/next week(s).

The second parameter `location` is need to make calculations relative to your location.

## Methods

- `WeekDays()` returns array of the names of the week days;
- `ShortDates(week int, include bool)` returns array of the dates formatted as "02.01" (day, month).
  `week` is necessary to calculate dates starting from `weekStart` to given `week`. If `include` is true you will get
  dates in range of all weeks in a row. If it`s false, you will get dates of only the last given week;
- `FullDates(week int, include bool)`  returns array of the dates formatted as "02.01.2006" (day, month, year). The same
  parameters as in `ShortDates` method;
- `DaysAndDates()` returns map of days and dates of the `weekStart`.

## Example

```go 
package main

import (
	"fmt"
	"time"

	"github.com/vaberof/goweekdate"
)

func Example() {
	// Create new WeekDate object with specified startWeek and location.
	weekDate := weekdate.New(time.Now(), "Asia/Novosibirsk")

	// Get names of the week days of type time.Weekday.
	weekDays := weekDate.WeekDays()

	// Get names of the week days of type string.
	stringWeekDays := weekdate.StringWeekDays(weekDays)

	// Get dates of the week of type time.Time.
	// 'week' is 1, 'include' is true.
	dates := weekDate.Dates(1, true)

	// Get dates formatted as ("02") on 1 week.
	day := weekdate.GetFormattedDates(dates, "02")

	// Get dates formatted as ("02.01") on 1 week.
	dayMonth := weekdate.GetFormattedDates(dates, "02.01")

	// Get dates formatted as ("02.01.2006") on 1 week.
	dayMonthYear := weekdate.GetFormattedDates(dates, "02.01.2006")

	// Get days and dates of the type of map[Monday:09.05.2022, Tuesday:10.05.2022...]
	daysAndDates := weekDate.DaysAndDates()

	fmt.Println(stringWeekDays) // [Monday Tuesday Wednesday Thursday Friday Saturday Sunday]
	fmt.Println(day)            // [09 10 11 12 13 14 15]
	fmt.Println(dayMonth)       // [09.05 10.05 11.05 12.05 13.05 14.05 15.05]
	fmt.Println(dayMonthYear)   // [09.05.2022 10.05.2022 11.05.2022 12.05.2022 13.05.2022 14.05.2022 15.05.2022]
	fmt.Println(daysAndDates)   // map[Friday:13.05.2022 Monday:09.05.2022 Saturday:14.05.2022 Sunday:15.05.2022 Thursday:12.05.2022 Tuesday:10.05.2022 Wednesday:11.05.2022]
}

func main() {
	Example()
}
```
