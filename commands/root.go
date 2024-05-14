package commands

// RootCommand is the root command of all commands
var RootCommand *Command

func init() {
	RootCommand = &Command{
		UsageLine: CommandEnv.Exec,
		Long:      "The root command",
	}
}

// RegisterDefaultCommand register a command to RootCommand and set this command as default
func RegisterDefaultCommand(cmd *Command) {
	RootCommand.Commands = append(RootCommand.Commands, cmd)
	defaultCommand = cmd.Name()
}

// RegisterCommand register a command to RootCommand
func RegisterCommand(cmd ...*Command) {
	RootCommand.Commands = append(RootCommand.Commands, cmd...)
}

// SetRootLong set root command long description
func SetRootLong(l string) {
	RootCommand.Long = l
}
