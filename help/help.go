package help

import (
	"fmt"
	"github.com/mcncl/pokedex/command"
)

func CommandHelp() command.CliCommand {
	command := command.CliCommand{
		Name:        "help",
		Description: "Display this help message",
		Callback: func() {
			fmt.Print("help called")
		},
	}
	return command
}
