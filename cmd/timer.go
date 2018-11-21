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
	if r.isToggle {
		timer, isActive := harvest.GetToggledTimer()
		fmt.Println(bttData{
			Text:            timer,
			IconPath:        getIconPath(isActive),
			BackgroundColor: "0,0,0,0",
		})
	} else {
		timer, isActive := harvest.GetLatestTimer()
		fmt.Println(bttData{
			Text:            timer,
			IconPath:        getIconPath(isActive),
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
