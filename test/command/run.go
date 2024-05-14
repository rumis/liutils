package main

import (
	"fmt"

	"github.com/rumis/liutils/commands"
)

var cmdRun = &commands.Command{
	UsageLine: "{{.Exec}} run [-config config.json] [-confdir dir]",
	Short:     "run {{.Exec}} with the config, the default command",
	Long: `
	Run {{.Exec}} with config, the default command.
	
	The -config=file flags set the config files for 
	{{.Exec}}. Multiple assign is accepted.
	
	The -confdir=dir flag sets a dir with multiple json config
	
	The -format=json flag sets the format of config files. 
	Default "auto".
	
	The -test flag tells {{.Exec}} to test config files only, 
	without launching the server.
	
	The -dump flag tells {{.Exec}} to print the merged config.
		`,
}

var configFiles = cmdRun.Flag.String("config", "", "Config path for Xray.")
var configDir = cmdRun.Flag.String("confdir", "", "A dir with multiple json config")
var dump = cmdRun.Flag.Bool("dump", false, "Dump merged config only, without launching Xray server.")
var test = cmdRun.Flag.Bool("test", false, "Test config file only, without launching Xray server.")
var format = cmdRun.Flag.String("format", "auto", "Format of input file.")

func init() {
	cmdRun.Run = func(cmd *commands.Command, args []string) {
		fmt.Println(append([]string{"start run..."}, args...))
		fmt.Println("config:", *configFiles)
	}
}
