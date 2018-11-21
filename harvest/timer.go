package harvest

import (
	"strconv"
	"time"

	"github.com/imroc/req"
)

// GetLatestTimer returns a latest timer given the
// time between the start of the day and now
func GetLatestTimer() (string, bool) {
	timeEntries, err := getTimeEntriesBetween(startOfDay(), time.Now())
	if err != nil {
		panic(err)
	}

	if timeEntries[0].IsRunning {
		return timeEntries[0].kitchenTimer(), true
	}

	return timeEntries[0].kitchenTimer(), false
}

// GetToggledTimer turns on and off a timer while also
// returning a timer it's latest timer. It returns a string
// for the timer, and a boolean for the it's latest state
func GetToggledTimer() (string, bool) {
	timeEntries, err := getTimeEntriesBetween(startOfDay(), time.Now())
	if err != nil {
		panic(err)
	}

	if timeEntries[0].IsRunning {
		return stopTimer(timeEntries[0]), false
	}

	return startTimer(timeEntries[0]), true
}

func stopTimer(latestEntry timeEntry) string {
	latestEntryID := strconv.Itoa(int(latestEntry.ID))
	if _, err := req.Patch(endpointStopTimer(latestEntryID), getAuthHeader()); err != nil {
		panic(err)
	}
	return latestEntry.kitchenTimer()
}

func startTimer(latestEntry timeEntry) string {
	latestEntryID := strconv.Itoa(int(latestEntry.ID))
	if _, err := req.Patch(endpointRestartTimer(latestEntryID), getAuthHeader()); err != nil {
		panic(err)
	}
	return "0:0"
}

func startOfDay() time.Time {
	y, m, d := time.Now().Date()
	jakarta, _ := time.LoadLocation("Asia/Jakarta")
	return time.Date(y, m, d, 0, 0, 0, 0, jakarta)
}
