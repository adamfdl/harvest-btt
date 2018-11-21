package harvest

import (
	"fmt"
	"os"

	"github.com/imroc/req"
)

var (
	accountID   = os.Getenv("HARVEST_ACCOUNT_ID")
	accessToken = os.Getenv("HARVEST_ACCESS_TOKEN")
)

var apiVersion = "2"

var (
	endpointHarvest      = "https://api.harvestapp.com/v" + apiVersion + "/"
	endpointTimeEntries  = endpointHarvest + "time_entries/"
	endpointStopTimer    = func(entryID string) string { return endpointTimeEntries + entryID + "/stop/" }
	endpointRestartTimer = func(entryID string) string { return endpointTimeEntries + entryID + "/restart/" }
)

func getAuthHeader() req.Header {
	return req.Header{
		"Authorization":      fmt.Sprintf("Bearer %s", accessToken),
		"Harvest-Account-Id": accountID,
	}
}
