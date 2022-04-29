# goweekdate

`goweekdate` is a simple go package that allows you to easily get dates of the week days in range that you need.

## Installation

    $ go get github.com/vaberof/goweekdate

## Example

``` 
package main

import (
    "time"
    "fmt"
    
    github.com/vaberof/goweekdate
)

func ExampleDate() {
    // Create new WeekDate object with specified startWeek and location.
    // With settled parameter startWeek as time.Now() calculations 
    // will be making from current week and 
    // as time.Now().Add(-time.Hour * 24 * 7) from previous week.
    weekDate := goweekdate.New(time.Now(), "Asia/Novosibirsk")

    // Get names of the week days.
    weekDays := weekDate.WeekDays()
    
    // Get dates formatted as ("02.01") on 2 weeks in a row.
    // 'week' is 2, 'include' is true.
    shortDates := weekDate.ShortDates(2, true)
    
    // Get dates formatted as ("02.01.2006") on the second week.
    // 'week' is 2, 'include' is false.
    fullDates := weekDate.FullDates(2, false)
    
    fmt.Println(weekDays)
    fmt.Println(shortDates)
    fmt.Println(fullDates)
}

func main() {
    ExampleDate()
}
```
