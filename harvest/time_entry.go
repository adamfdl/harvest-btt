package harvest

import (
	"fmt"
	"math"
	"time"
)

type timeEntries struct {
	TimeEntries []TimeEntry `json:"time_entries"`
}

type TimeEntry struct {
	ID        int64   `json:"id"`
	Hours     float64 `json:"hours"`
	IsRunning bool    `json:"is_running"`
	Billable  bool    `json:"billable"`
}

// toKitchenTime formats a minute representation of an hour
// to standard kitchen time without its meridiem.
// As an example, setting 0.2 as its parameters would return 0:12
func (t *TimeEntry) KitchenTimer() string {
	hours := math.Floor(t.Hours)
	minutes := math.Ceil((t.Hours - hours) * 60)
	return fmt.Sprintf("%d:%d", int(hours), int(minutes))
}

func (h *harvestAPI) GetTimeEntriesBetween(from, to time.Time) (*timeEntries, error) {
	timeEntries := &timeEntries{}

	params := params{"from": from.Format("20060102"), "to": to.Format("20060102")}
	resp, err := h.sendRequest(_get, "/time_entries", params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := decodeResponse(resp, timeEntries); err != nil {
		return nil, err
	}

	return timeEntries, nil
}

func (h *harvestAPI) RestartTimeEntry(entryID int64) (*TimeEntry, error) {
	TimeEntry := &TimeEntry{}

	resp, err := h.sendRequest(_patch, fmt.Sprintf("/time_entries/%d/restart", entryID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := decodeResponse(resp, TimeEntry); err != nil {
		return nil, err
	}

	return TimeEntry, nil
}

func (h *harvestAPI) StopTimeEntry(entryID int64) (*TimeEntry, error) {
	TimeEntry := &TimeEntry{}

	resp, err := h.sendRequest(_patch, fmt.Sprintf("/time_entries/%d/stop", entryID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := decodeResponse(resp, TimeEntry); err != nil {
		return nil, err
	}

	return TimeEntry, nil
}
