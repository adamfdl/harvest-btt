package harvest

import (
	"time"

	"github.com/grsmv/goweek"
)

// GetBillables is a wrapper for retrieving billable hours
// between the start and the end of the week given current time
func GetBillables() float64 {
	timeEntries, err := getTimeEntriesBetween(startOfWeek(), time.Now())
	if err != nil {
		panic(err)
	}
	return getBillablesFromTimeEntries(timeEntries)
}

// GetNonBillables is a wrapper for retrieving non billable
// hours between the start and the end of the week given current time
func GetNonBillables() float64 {
	timeEntries, err := getTimeEntriesBetween(startOfWeek(), endOfWeek())
	if err != nil {
		panic(err)
	}
	return getNonBillablesFromTimeEntries(timeEntries)
}

func getBillablesFromTimeEntries(timeEntries []timeEntry) float64 {
	var billables float64
	for _, timeEntry := range timeEntries {
		if timeEntry.isBillable() {
			billables += timeEntry.Hours
		}
	}
	return billables
}

func getNonBillablesFromTimeEntries(timeEntries []timeEntry) float64 {
	var nonBillables float64
	for _, timeEntry := range timeEntries {
		if !timeEntry.isBillable() {
			nonBillables += timeEntry.Hours
		}
	}
	return nonBillables
}

func startOfWeek() time.Time {
	week, _ := goweek.NewWeek(time.Now().ISOWeek())
	return week.Days[0]
}

func endOfWeek() time.Time {
	week, _ := goweek.NewWeek(time.Now().AddDate(0, 0, 7).ISOWeek())
	return week.Days[6]
}
