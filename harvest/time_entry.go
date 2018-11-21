package harvest

import (
	"fmt"
	"math"
	"time"

	"github.com/imroc/req"
)

type timeEntries struct {
	TimeEntries []timeEntry `json:"time_entries"`
}

type timeEntry struct {
	ID        int64   `json:"id"`
	Hours     float64 `json:"hours"`
	IsRunning bool    `json:"is_running"`
	Billable  bool    `json:"billable"`
}

func (t *timeEntry) isBillable() bool {
	return t.Billable
}

// toKitchenTime formats a minute representation of an hour
// to standard kitchen time without its meridiem.
// As an example, setting 0.2 as its parameters would return 0:12
func (t *timeEntry) kitchenTimer() string {
	hours := math.Floor(t.Hours)
	minutes := math.Ceil((t.Hours - hours) * 60)
	return fmt.Sprintf("%d:%d", int(hours), int(minutes))
}

func getTimeEntriesBetween(from, to time.Time) ([]timeEntry, error) {
	var (
		qParam = req.QueryParam{
			"from": from.Format("20060102"),
			"to":   to.Format("20060102"),
		}
		harvest = &timeEntries{}
	)

	r, err := req.Get(endpointTimeEntries, getAuthHeader(), qParam)
	if err != nil {
		return nil, err
	}
	r.ToJSON(&harvest)
	return harvest.TimeEntries, nil
}
