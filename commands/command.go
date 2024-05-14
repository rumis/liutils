package commands

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
)

var (
	exitStatus = 0
	exitMu     sync.Mutex
)

// SetExitStatus set exit status code
func SetExitStatus(n int) {
	exitMu.Lock()
	if exitStatus < n {
		exitStatus = n
	}
	exitMu.Unlock()
}

// Exit exits with code set with SetExitStatus()
func Exit() {
	os.Exit(exitStatus)
}

// Fatalf logs error and exit with code 1
func Fatalf(format string, args ...interface{}) {
	Errorf(format, args...)
	Exit()
}

// Errorf logs error and set exit status to 1, but not exit
func Errorf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
	SetExitStatus(1)
}

// Command base command line struct
type Command struct {
	// Run run the command
	Run func(cmd *Command, args []string)

	// UsageLine is the one-line usage message.
	// The first word in the line is taken to be the command's name.
	// UsageLine support go template syntax. It's recommended to use "{{.Exec}}" instead of hardcoding name
	UsageLine string

	// Short is the short description shown in the 'go help' output.
	// Short support go template syntax. It's recommended to use "{{.Exec}}" instead of hardcoding name
	Short string

	// Long is the long message shown in the 'go help <this-command>' output.
	// Long support go template syntax. It's recommended to use "{{.Exec}}" instead of hardcoding name
	Long string

	// Flag is a set of flags specific to this command.
	Flag flag.FlagSet

	// CustomFlags indicates that the command will do its own
	CustomFlags bool

	// Commands lists the available commands and help topics.
	Commands []*Command
}

// LongName returns the command's long name: all the words in the usage line between first word
func (c *Command) LongName() string {
	name := c.UsageLine
	if i := strings.Index(name, " ["); i >= 0 {
		name = strings.TrimSpace(name[:i])
	}
	if i := strings.Index(name, " "); i >= 0 {
		name = name[i+1:]
	} else {
		name = ""
	}
	return strings.TrimSpace(name)
}

// Name returns the command's short name: the last word in the usage line before a flag or argument
func (c *Command) Name() string {
	name := c.LongName()
	if i := strings.LastIndex(name, " "); i >= 0 {
		name = name[i+1:]
	}
	return strings.TrimSpace(name)
}

// Usage prints usage of the Command
// this will exit the application
func (c *Command) Usage() {
	buildCommandText(c)
	fmt.Fprintf(os.Stderr, "usage: %s\n", c.UsageLine)
	fmt.Fprintf(os.Stderr, "Run 'help %s' for details.\n", c.LongName())
	SetExitStatus(2)
	Exit()
}

// Runnable reports whether the command can be run; otherwise
// it is a documentation pseudo-command such as importpath.
func (c *Command) Runnable() bool {
	return c.Run != nil
}
