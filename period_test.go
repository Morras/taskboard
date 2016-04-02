package taskboard

import "testing"
import "time"
//These test asume a 14 day period
//With the first period on February 1st 2016

func TestPeriodStartByTime(t *testing.T) {
	expectedPeriodStart := time.Date(2016, 3, 14, 0, 0, 0, 0, GetDefaultLocation()).UTC()

	testDate := time.Date(2016, 3, 23, 10, 14, 0, 0, GetDefaultLocation()).UTC()

	actualPeriodStart := PeriodStartByTime(testDate)

	if actualPeriodStart != expectedPeriodStart {
		t.Errorf("Got wrong period start, expected %v got %v", expectedPeriodStart, actualPeriodStart)
	}
    
	expectedPeriodStart = time.Date(2016, 7, 4, 0, 0, 0, 0, GetDefaultLocation()).UTC()

	testDate = time.Date(2016, 7, 14, 13, 14, 0, 0, GetDefaultLocation()).UTC()

	actualPeriodStart = PeriodStartByTime(testDate)

	if actualPeriodStart != expectedPeriodStart {
		t.Errorf("Got wrong period start, expected %v got %v", expectedPeriodStart, actualPeriodStart)
	}
}

func TestPeriodAtMidnightLocalTime(t *testing.T){
	expectedPeriodStart := time.Date(2016, 2, 1, 0, 0, 0, 0, GetDefaultLocation()).UTC()

	testDate := time.Date(2016, 2, 1, 0, 0, 0, 0, GetDefaultLocation()).UTC()

	actualPeriodStart := PeriodStartByTime(testDate)

	if actualPeriodStart != expectedPeriodStart {
		t.Errorf("Got wrong period start, expected %v got %v", expectedPeriodStart, actualPeriodStart)
	}
}

func TestPeriodGoingToSummerTime(t *testing.T) {
	expectedPeriodStart := time.Date(2016, 3, 28, 0, 0, 0, 0, GetDefaultLocation()).UTC()

	testDate := time.Date(2016, 3, 31, 0, 0, 0, 0, GetDefaultLocation()).UTC()

	actualPeriodStart := PeriodStartByTime(testDate)

	if actualPeriodStart != expectedPeriodStart {
		t.Errorf("Got wrong period start, expected %v got %v", expectedPeriodStart, actualPeriodStart)
	}
}

func TestPeriodGoingFromSummerTime(t *testing.T) {
   	expectedPeriodStart := time.Date(2016, 10, 24, 0, 0, 0, 0, GetDefaultLocation()).UTC()

	testDate := time.Date(2016, 11, 2, 0, 0, 0, 0, GetDefaultLocation()).UTC()

	actualPeriodStart := PeriodStartByTime(testDate)

	if actualPeriodStart != expectedPeriodStart {
		t.Errorf("Got wrong period start, expected %v got %v", expectedPeriodStart, actualPeriodStart)
	}
}

func TestPeriodBeforeBeginning(t *testing.T) {
	expectedPeriodStart := time.Date(2015, 12, 21, 0, 0, 0, 0, GetDefaultLocation()).UTC()

	testDate := time.Date(2016, 1, 2, 12, 0, 0, 0, GetDefaultLocation()).UTC()

	actualPeriodStart := PeriodStartByTime(testDate)

	if actualPeriodStart != expectedPeriodStart {
		t.Errorf("Got wrong time for period before canonical start period, expected %v got %v", expectedPeriodStart, actualPeriodStart)
	}
}

func TestPeriodInSummertimeBeforeBeginning(t *testing.T) {
    expectedPeriodStart := time.Date(2015, 8, 3, 0, 0, 0, 0, GetDefaultLocation()).UTC()

	testDate := time.Date(2015, 8, 10, 12, 0, 0, 0, GetDefaultLocation()).UTC()

	actualPeriodStart := PeriodStartByTime(testDate)

	if actualPeriodStart != expectedPeriodStart {
		t.Errorf("Got wrong time for period before canonical start period, expected %v got %v", expectedPeriodStart, actualPeriodStart)
	}
}

func TestPeriodSameDay(t *testing.T) {
	expectedPeriodStart := time.Date(2016, 02, 15, 0, 0, 0, 0, GetDefaultLocation()).UTC()

	testDate := time.Date(2016, 2, 15, 7, 0, 0, 0, GetDefaultLocation()).UTC()

	actualPeriodStart := PeriodStartByTime(testDate)

	if actualPeriodStart != expectedPeriodStart {
		t.Errorf("Got wrong time for date that is the same as a period start, expected %v got %v", expectedPeriodStart, actualPeriodStart)
	}
}

func TestCurrentPeriodStart(t *testing.T) {
	currentPeriodStart := CurrentPeriodStart()
	expectedPeriodStart := PeriodStartByTime(time.Now().UTC())

	if currentPeriodStart != expectedPeriodStart {
		t.Error("CurrentPeriodStart does not give the same result as calling PeriodStart with now. Got ", currentPeriodStart, " Expected ", expectedPeriodStart)
	}
}
