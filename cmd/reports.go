package cmd

import (
	"fmt"

	"github.com/adamfdl/harvest-tb/harvest"
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
	if r.isNonBillables {
		fmt.Println(bttData{
			Text: fmt.Sprintf("Non Billable: %.2f", harvest.GetNonBillables()),
		})
	} else {
		fmt.Println(bttData{
			Text: fmt.Sprintf("Billable: %.2f", harvest.GetBillables()),
		})
	}
}
