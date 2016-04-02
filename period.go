package taskboard

import "time"
import "fmt"

/*
 * A task is assigned to a period by specifying
 * the first day of the period on the task.
 * A period lasts for two weeks, and the first
 * period starts on February 1st 2016.
 */

var defaultLocationString = "Europe/Copenhagen"
var firstPeriodStart = time.Date(2016, 02, 1, 0, 0, 0, 0, GetDefaultLocation()).UTC()
var periodLength = 14

func GetDefaultLocation() *time.Location {
	loc, err := time.LoadLocation(defaultLocationString)
	if err != nil {
		panic(fmt.Sprintf("Configuration error, cannot load configured locatoin %v", err))
	}
	return loc
}

func GetFirstPeriodStartUTCHour() int {
    return firstPeriodStart.Hour()
}

//Must be called with a time in UTC
func PeriodStartByTime(t time.Time) time.Time { //TODO Test this
	periodStart := time.Date(t.Year(), t.Month(), t.Day(), GetFirstPeriodStartUTCHour(), 0, 0, 0, time.UTC)

	delta := periodStart.Sub(firstPeriodStart)

	deltaHours := int(delta / time.Hour)

	//We will run into a problem when passing over to summer time, as that
	//will remove an hour from one of the days, so we artificially add it here.
	//It does not matter that we also add this hour for all other days as it will be
	//truncated away when converting hours into whole days, but we need to be aware of
	//the parity of the delta
	if deltaHours < 0 {
		deltaHours--
	} else {
		deltaHours++
	}
	deltaDays := int(deltaHours / 24)

	daysOffset := deltaDays % periodLength

	//Problem with a negative offset is that it is the number of days up to
	//the next period starts and not down to the current period start.
	if daysOffset < 0 {
		daysOffset = -daysOffset
		daysOffset = periodLength - daysOffset
	}

	periodStart = periodStart.AddDate(0, 0, -daysOffset)

    convertedToDefaultLocation := periodStart.In(GetDefaultLocation())

    return time.Date(convertedToDefaultLocation.Year(), convertedToDefaultLocation.Month(), convertedToDefaultLocation.Day(), 0, 0, 0, 0, GetDefaultLocation()).UTC()
}

//CurrentPeriodStart uses local time which may not 
//be the same timezone as GetDefaultLocation
func CurrentPeriodStart() time.Time { 
	return PeriodStartByTime(time.Now().UTC())
}
