package cmd

import (
	"fmt"
	"time"

	"github.com/adamfdl/harvest-btt/harvest"
	"github.com/grsmv/goweek"
	"github.com/spf13/cobra"
)

type reportCmd struct {
	*baseCmd
	isNonBillables bool
}

func newReportCmd() *reportCmd {
	command := &reportCmd{}
	command.baseCmd = newBaseCmd(&cobra.Command{
		Use:   "reports",
		Short: "Get a report of this week's timesheet, billables or non billables",
		Run:   command.getReports,
	})

	command.cmd.Flags().BoolVar(&command.isNonBillables, "non-billables", false, "")
	return command

}

func (r *reportCmd) getReports(*cobra.Command, []string) {
	h := harvest.NewHarvestAPI()
	timeEntries, err := h.GetTimeEntriesBetween(startOfWeek(), endOfWeek())
	if err != nil {
		fmt.Println(bttData{
			Text: fmt.Sprintf("Error!"),
		})
		return
	}

	var hours float64
	if r.isNonBillables {
		hours = getNonBillablesFromTimeEntries(timeEntries.TimeEntries)
	} else {
		hours = getBillablesFromTimeEntries(timeEntries.TimeEntries)
	}

	fmt.Println(bttData{
		Text: fmt.Sprintf("Billable: %.2f", hours),
	})
}

func getBillablesFromTimeEntries(timeEntries []harvest.TimeEntry) float64 {
	var billables float64
	for _, timeEntry := range timeEntries {
		if timeEntry.Billable {
			billables += timeEntry.Hours
		}
	}
	return billables
}

func getNonBillablesFromTimeEntries(timeEntries []harvest.TimeEntry) float64 {
	var nonBillables float64
	for _, timeEntry := range timeEntries {
		if !timeEntry.Billable {
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
