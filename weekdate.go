package goweekdate

import (
	"log"
	"time"
)

// A WeekDate represents attributes which necessary
// to make calculations of dates.
type WeekDate struct {
	curDay    time.Time
	weekStart time.Time
	weekDay   time.Weekday
	weekDays  []time.Weekday

	location  string
	shortDate string
	fullDate  string

	shortDates []string
	fullDates  []string
}

// New creates a new WeekDate object from the specified weekStart and location.
//
// weekStart represents given local time of type time.Time.
// Can be modified with time.Now.Add to start calculations from previous/next week(s).
//
// location represents given location of type string which will be used as time.Location to set it.
func New(weekStart time.Time, location string) *WeekDate {
	return &WeekDate{
		weekStart: weekStart,
		location:  location,
	}
}

// WeekDays returns array of the names of the week days.
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

// ShortDates returns array of dates formatted as "02.01" (day, month).
//
// week is necessary to get dates from weekStart to given week number.
//
// include is option that allows to get dates from weekStart to given week in a row if it`s true,
// or get dates only of the one given week if it`s false.
func (w *WeekDate) ShortDates(week int, include bool) []string {
	w.shortDates = []string{}
	w.monday()

	if !include {
		w.skipShortWeek(week)
		return w.shortDates
	}

	for weekDay := 1; weekDay <= 7*week; weekDay++ {
		w.shortDates = append(w.shortDates, w.shortDate)
		w.curDay = w.curDay.Add(time.Hour * 24)
		w.shortDate = w.curDay.Format("02.01")
	}

	return w.shortDates
}

// FullDates returns array of full dates formatted as "02.01.2006" (day, month, year).
//
// week is necessary to get dates from weekStart to given week number.
//
// include is option that allows to get dates from weekStart to given week in a row if it`s true,
// or get dates only of the one given week if it`s false.
func (w *WeekDate) FullDates(week int, include bool) []string {
	w.fullDates = []string{}
	w.monday()

	if !include {
		w.skipFullWeek(week)
		return w.fullDates
	}

	for weekDay := 1; weekDay <= 7*week; weekDay++ {
		w.fullDates = append(w.fullDates, w.fullDate)
		w.curDay = w.curDay.Add(time.Hour * 24)
		w.fullDate = w.curDay.Format("02.01.2006")
	}

	return w.fullDates
}

// currentDay sets day and weekday from specified weekStart
// for curDay.
func (w *WeekDate) currentDay() {
	w.curDay = w.weekStart.In(location(w.location))
	w.weekDay = w.curDay.Weekday()
}

// monday sets curDay of the specified weekStart to the Monday.
func (w *WeekDate) monday() {
	w.currentDay()

	for w.weekDay.String() != "Monday" {
		w.curDay = w.curDay.Add(-time.Hour * 24)
		w.weekDay = w.curDay.Weekday()
	}

	w.shortDate = w.curDay.Format("02.01")
	w.fullDate = w.curDay.Format("02.01.2006")
}

// skipShortWeek skips weeks to get dates formatted as ("02.01") of the one given week if include in ShortDates is false.
func (w *WeekDate) skipShortWeek(week int) {
	skipWeek := 7*week - 7

	for weekDay := 1; weekDay <= skipWeek; weekDay++ {
		w.curDay = w.curDay.Add(time.Hour * 24)
		w.shortDate = w.curDay.Format("02.01.2006")
	}

	for weekDay := 1; weekDay <= 7; weekDay++ {
		w.shortDates = append(w.shortDates, w.shortDate)
		w.curDay = w.curDay.Add(time.Hour * 24)
		w.shortDate = w.curDay.Format("02.01")
	}
}

// skipFullWeek skips weeks to get dates formatted as ("02.01.2006") of the one given week if include in FullDates is false.
func (w *WeekDate) skipFullWeek(week int) {
	skipWeek := 7*week - 7

	for weekDay := 1; weekDay <= skipWeek; weekDay++ {
		w.curDay = w.curDay.Add(time.Hour * 24)
		w.fullDate = w.curDay.Format("02.01.2006")
	}

	for weekDay := 1; weekDay <= 7; weekDay++ {
		w.fullDates = append(w.fullDates, w.fullDate)
		w.curDay = w.curDay.Add(time.Hour * 24)
		w.fullDate = w.curDay.Format("02.01.2006")
	}
}

// location sets given location with time.LoadLocation.
// Fatales if cannot load given location.
func location(location string) *time.Location {
	loc, err := time.LoadLocation(location)
	if err != nil {
		log.Fatalf("failed to load a location: %v", err)
	}

	return loc
}
