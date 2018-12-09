package cmd

import (
	"fmt"

	"github.com/adamfdl/harvest-btt/harvest"
	"github.com/spf13/cobra"
)

type timerCmd struct {
	*baseCmd
	isToggle bool
}

func newTimerCmd() *timerCmd {
	command := &timerCmd{}
	command.baseCmd = newBaseCmd(&cobra.Command{
		Use:   "timer",
		Short: "Timer functionality",
		Run:   command.timer,
	})

	command.cmd.Flags().BoolVar(&command.isToggle, "toggle", false, "")
	return command
}

func (r *timerCmd) timer(*cobra.Command, []string) {
	h := harvest.NewHarvestAPI()
	timeEntries, err := h.GetTimeEntriesBetween(startOfWeek(), endOfWeek())
	if err != nil {
		fmt.Println(bttData{
			Text: fmt.Sprintf("Error!"),
		})
		return
	}

	if r.isToggle {
		if timeEntries.TimeEntries[0].IsRunning {
			timeEntry, err := h.StopTimeEntry(timeEntries.TimeEntries[0].ID)
			if err != nil {
				fmt.Println(bttData{
					Text: fmt.Sprintf("Error!"),
				})
				return
			}
			if !timeEntry.IsRunning {
				fmt.Println(bttData{
					Text:            timeEntry.KitchenTimer(),
					IconPath:        getIconPath(timeEntries.TimeEntries[0].IsRunning),
					BackgroundColor: "0,0,0,0",
				})
			}
		} else {
			timeEntry, err := h.RestartTimeEntry(timeEntries.TimeEntries[0].ID)
			if err != nil {
				fmt.Println(bttData{
					Text: fmt.Sprintf("Error!"),
				})
				return
			}
			if timeEntry.IsRunning {
				fmt.Println(bttData{
					Text:            timeEntry.KitchenTimer(),
					IconPath:        getIconPath(timeEntries.TimeEntries[0].IsRunning),
					BackgroundColor: "0,0,0,0",
				})
			}
		}
	} else {
		fmt.Println(bttData{
			Text:            timeEntries.TimeEntries[0].KitchenTimer(),
			IconPath:        getIconPath(timeEntries.TimeEntries[0].IsRunning),
			BackgroundColor: "0,0,0,0",
		})
	}
}

func getIconPath(isActive bool) string {
	basePath := "/Users/adamfadhil/Projects/go/src/github.com/adamfdl/harvest-btt/resources/harvest-"
	if isActive {
		return fmt.Sprintf("%s%s", basePath, "active.png")
	}
	return fmt.Sprintf("%s%s", basePath, "inactive.png")
}
