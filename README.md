# goweekdate

`goweekdate` is a simple go package that allows you to easily get dates of the week days in range that you need.

## Installation

    $ go get github.com/vaberof/goweekdate

## Initialization

```go
weekDate := weekdate.New(time.Now(), "Asia/Novosibirsk")
```

With setted first parameter `startWeek` as time.Now() calculations will be making from current week. You can modify it
by using time.Now().Add() method to make calculations from previous/next week(s).

The second parameter `location` is need to make calculations relative to your location.

## Methods

- `WeekDays()` returns array of the names of the week days;
- `ShortDates(week int, include bool)` returns dates formatted as "02.01" (day, month).
  `week` is necessary to calculate dates starting from `startWeek` to given `week`. If `include` is true you will
  get dates in range of all weeks in a row. If it`s false, you will get dates of only the last given week;
- `FullDates(week int, include bool)`  returns dates formatted as "02.01.2006" (day, month, year). The same parameters
  as in `ShortDates()`.

## Example

```go 
package main

import (
    "time"
    "fmt"
    
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
    
    fmt.Println(weekDays)   // [Monday Tuesday Wednesday Thursday Friday Saturday Sunday]
    fmt.Println(shortDates) // [25.04 26.04 27.04 28.04 29.04 30.04 01.05 02.05 03.05 04.05 05.05 06.05 07.05 08.05]
    fmt.Println(fullDates)  // [02.05.2022 03.05.2022 04.05.2022 05.05.2022 06.05.2022 07.05.2022 08.05.2022]
}

func main() {
    Example()
}
```
