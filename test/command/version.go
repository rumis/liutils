package main

import (
	"fmt"

	"github.com/rumis/liutils/commands"
)

var cmdVersion = &commands.Command{
	UsageLine: "{{.Exec}} version",
	Short:     "show current version of {{.Exec}}",
	Long:      "version prinets the build information of the {{.Exec}} executables",
	Run: func(cmd *commands.Command, args []string) {
		v1 := fmt.Sprintf("%v.%v.%v", 1, 1, 1)
		fmt.Println(v1)
	},
}
