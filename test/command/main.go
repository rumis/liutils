package main

import (
	"github.com/rumis/liutils/commands"
)

func main() {

	commands.SetRootLong(" test for command line")
	commands.RegisterDefaultCommand(cmdRun)
	commands.RegisterCommand(cmdVersion)

	commands.Execute()
}
