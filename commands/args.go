package commands

import (
	"flag"
	"os"
)

var defaultCommand = ""

func ArgsCompatible() []string {

	if len(os.Args) == 1 {
		return []string{os.Args[0], defaultCommand} // default command
	}
	if os.Args[1][0] != '-' {
		return os.Args
	}
	version := false
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.BoolVar(&version, "version", false, "show version")
	fs.Usage = func() {}
	fs.SetOutput(&null{})
	err := fs.Parse(os.Args[1:])
	if err == flag.ErrHelp {
		return []string{os.Args[0], "help"}
	}
	if version {
		return []string{os.Args[0], "version"}
	}
	return append([]string{os.Args[0], defaultCommand}, os.Args[1:]...)
}

type null struct{}

func (n *null) Write(p []byte) (int, error) {
	return len(p), nil
}
