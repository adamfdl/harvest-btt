package harvest

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var (
	server  *httptest.Server
	mux     *http.ServeMux
	harvest harvestAPI
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	harvest = harvestAPI{client: server.Client(), baseURL: server.URL}
}

func teardown() {
	server.Close()
}

func TestGetTimeEntries(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/time_entries", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(timeEntriesResp))
	}))

	timeEntries, err := harvest.GetTimeEntriesBetween(time.Time{}, time.Time{})
	require.NoError(t, err)
	require.Equal(t, timeEntries.TimeEntries[0].ID, int64(1))
	require.Equal(t, timeEntries.TimeEntries[0].Hours, float64(2.0))
	require.Equal(t, timeEntries.TimeEntries[0].IsRunning, false)
	require.Equal(t, timeEntries.TimeEntries[0].Billable, true)
}

func TestRestartTimer(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/time_entries", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(timeEntriesResp))
	}))
	mux.HandleFunc("/time_entries/1/restart", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{
            "id": 1,
            "hours": 0,
            "is_running": true,
            "billable": true
        }`))
	}))

	timeEntries, err := harvest.GetTimeEntriesBetween(time.Time{}, time.Time{})
	require.NoError(t, err)
	require.Equal(t, timeEntries.TimeEntries[0].ID, int64(1))
	require.Equal(t, timeEntries.TimeEntries[0].Hours, float64(2.0))
	require.Equal(t, timeEntries.TimeEntries[0].IsRunning, false)
	require.Equal(t, timeEntries.TimeEntries[0].Billable, true)

	timeEntry, err := harvest.RestartTimeEntry(timeEntries.TimeEntries[0].ID)
	require.NoError(t, err)
	require.Equal(t, timeEntry.ID, int64(1))
	require.Equal(t, timeEntry.Hours, float64(0))
	require.Equal(t, timeEntry.IsRunning, true)
	require.Equal(t, timeEntry.Billable, true)
}

func TestStopTimer(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/time_entries", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(timeEntriesResp))
	}))
	mux.HandleFunc("/time_entries/1/stop", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{
            "id": 1,
            "hours": 1.2,
            "is_running": false,
            "billable": true
        }`))
	}))

	timeEntries, err := harvest.GetTimeEntriesBetween(time.Time{}, time.Time{})
	require.NoError(t, err)
	require.Equal(t, timeEntries.TimeEntries[0].ID, int64(1))
	require.Equal(t, timeEntries.TimeEntries[0].Hours, float64(2.0))
	require.Equal(t, timeEntries.TimeEntries[0].IsRunning, false)
	require.Equal(t, timeEntries.TimeEntries[0].Billable, true)

	timeEntry, err := harvest.StopTimeEntry(timeEntries.TimeEntries[0].ID)
	require.NoError(t, err)
	require.Equal(t, timeEntry.ID, int64(1))
	require.Equal(t, timeEntry.Hours, float64(1.2))
	require.Equal(t, timeEntry.IsRunning, false)
	require.Equal(t, timeEntry.Billable, true)
}

const timeEntriesResp = `
{
    "time_entries": [
        {
            "id": 1,
            "hours": 2.0,
            "is_running": false,
            "billable": true
        }
    ]
}`
