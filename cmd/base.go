package cmd

import (
	"github.com/spf13/cobra"
)

type baseCmd struct {
	cmd *cobra.Command
}

func newBaseCmd(cmd *cobra.Command) *baseCmd {
	return &baseCmd{cmd: cmd}
}

func (c *baseCmd) getCommand() *cobra.Command {
	return c.cmd
}

type commandsBuilder struct {
	commands []*cobra.Command
}

func newCommandsBuilder() *commandsBuilder {
	return &commandsBuilder{}
}

func (cb *commandsBuilder) appendCommands(commands ...*cobra.Command) *commandsBuilder {
	cb.commands = append(cb.commands, commands...)
	return cb
}

func (cb *commandsBuilder) addAllCommands() *commandsBuilder {
	cb.appendCommands(
		newReportCmd().getCommand(),
		newTimerCmd().getCommand(),
	)
	return cb
}

func (cb *commandsBuilder) build() {
	for _, command := range cb.commands {
		rootCmd.AddCommand(command)
	}
}
