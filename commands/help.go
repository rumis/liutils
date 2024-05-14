package commands

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var usageTemplate = `{{.Long | trim}}

Usage:

	{{.UsageLine}} <command> [arguments]

The commands are:
{{range .Commands}}{{if and (ne .Short "") (or (.Runnable) .Commands)}}
	{{.Name | width $.CommandsWidth}} {{.Short}}{{end}}{{end}}

Use "{{.Exec}} help{{with .LongName}} {{.}}{{end}} <command>" for more information about a command.
`

var helpTemplate = `{{if .Runnable}}usage: {{.UsageLine}}

{{end}}{{.Long | trim}}
`

// Help implements the 'help' command.
func Help(w io.Writer, args []string) {
	cmd := RootCommand
Args:
	for i, arg := range args {
		for _, sub := range cmd.Commands {
			if sub.Name() == arg {
				cmd = sub
				continue Args
			}
		}

		// helpSuccess is the help command using as many args as possible that would succeed.
		helpSuccess := CommandEnv.Exec + " help"
		if i > 0 {
			helpSuccess += " " + strings.Join(args[:i], " ")
		}
		fmt.Fprintf(os.Stderr, "%s help %s: unknown help topic. Run '%s'.\n", CommandEnv.Exec, strings.Join(args, " "), helpSuccess)
		SetExitStatus(2) // failed at 'help cmd'
		Exit()
	}
	if len(cmd.Commands) > 0 {
		PrintUsage(os.Stdout, cmd)
	} else {
		buildCommandText(cmd)
		tmpl(os.Stdout, helpTemplate, makeTmplData(cmd))
	}
}

// PrintUsage prints usage of cmd to w
func PrintUsage(w io.Writer, cmd *Command) {
	buildCommandText(cmd)
	bw := bufio.NewWriter(w)
	tmpl(bw, usageTemplate, makeTmplData(cmd))
	bw.Flush()
}
