package weekdate

import (
	"log"
	"time"
)

// A WeekDate represents attributes which necessary
// to make calculations of dates.
type WeekDate struct {
	curDay    time.Time
	weekStart time.Time

	date  time.Time
	dates []time.Time

	weekDay  time.Weekday
	weekDays []time.Weekday

	daysAndDates map[string]string

	location string
}

// New creates a new WeekDate object from the specified weekStart and location.
//
// weekStart represents given local time of type time.Time.
// Can be modified with time.Now.Add to start calculations from previous/next week(s).
//
// location represents given location of type string which will be used as time.Location to set it.
func New(weekStart time.Time, location string) *WeekDate {
	return &WeekDate{
		weekStart:    weekStart,
		location:     location,
		daysAndDates: map[string]string{},
	}
}

// WeekDays returns array of type time.Weekday of the names of the week days.
func (w *WeekDate) WeekDays() []time.Weekday {
	w.weekDays = []time.Weekday{}
	w.monday()

	for weekDay := 1; weekDay <= 7; weekDay++ {
		w.weekDays = append(w.weekDays, w.weekDay)
		w.curDay = w.curDay.Add(time.Hour * 24)
		w.weekDay = w.curDay.Weekday()
	}

	return w.weekDays
}

// StringWeekDays returns array of type string of the names of the week days.
func StringWeekDays(weekDays []time.Weekday) []string {
	var stringWeekDays []string

	for i := 0; i < len(weekDays); i++ {
		stringWeekDays = append(stringWeekDays, weekDays[i].String())
	}

	return stringWeekDays
}

// Dates returns array of type time.Time of dates.
//
// week is necessary to get dates from weekStart to given week number.
//
// include is option that allows to get dates from weekStart to given week in a row if it`s true,
// or get dates only of the one given week if it`s false.
func (w *WeekDate) Dates(week int, include bool) []time.Time {
	w.dates = []time.Time{}
	w.monday()

	if !include {
		w.skipWeek(week)
		return w.dates
	}

	for weekDay := 1; weekDay <= 7*week; weekDay++ {
		w.dates = append(w.dates, w.date)
		w.curDay = w.curDay.Add(time.Hour * 24)
		w.date = w.curDay
	}

	return w.dates
}

// GetFormattedDates returns array of type string of formatted dates.
// dates is given array from Dates method.
// format is necessary to get dates in format that you need.
func GetFormattedDates(dates []time.Time, format string) []string {
	var formattedDates []string

	for i := 0; i < len(dates); i++ {
		formattedDates = append(formattedDates, dates[i].Format(format))
	}

	return formattedDates
}

// DaysAndDates returns map of days and dates of the weekStart, where day is a key and date is a value.
func (w *WeekDate) DaysAndDates() map[string]string {
	w.monday()

	for weekDay := 1; weekDay <= 7; weekDay++ {
		w.daysAndDates[w.curDay.Weekday().String()] = w.curDay.Format("02.01.2006")
		w.curDay = w.curDay.Add(time.Hour * 24)
	}

	return w.daysAndDates
}

// skipWeek skips weeks starting from weekStart
// to get dates of the one given week if include in Dates is false.
func (w *WeekDate) skipWeek(week int) {
	skipWeek := 7*week - 7

	for weekDay := 1; weekDay <= skipWeek; weekDay++ {
		w.curDay = w.curDay.Add(time.Hour * 24)
	}
	w.date = w.curDay

	for weekDay := 1; weekDay <= 7; weekDay++ {
		w.dates = append(w.dates, w.date)
		w.curDay = w.curDay.Add(time.Hour * 24)
		w.date = w.curDay
	}
}

// currentDay sets day and weekday from specified weekStart
// for curDay.
func (w *WeekDate) currentDay() {
	w.curDay = w.weekStart.In(setLocation(w.location))
	w.weekDay = w.curDay.Weekday()
}

// monday sets curDay of the specified weekStart to the Monday.
func (w *WeekDate) monday() {
	w.currentDay()

	for w.weekDay.String() != "Monday" {
		w.curDay = w.curDay.Add(-time.Hour * 24)
		w.weekDay = w.curDay.Weekday()
	}

	w.date = w.curDay
}

// setLocation sets given location with time.LoadLocation.
// Fatales if cannot load given location.
func setLocation(location string) *time.Location {
	loc, err := time.LoadLocation(location)
	if err != nil {
		log.Fatalf("failed to load a setLocation: %v", err)
	}

	return loc
}
