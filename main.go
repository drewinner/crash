package main

import (
	"fmt"
	"github.com/drewinner/crash/cmd"
	"github.com/spf13/cobra"
	"os"
)
var versionFlag bool
var RootCmd = &cobra.Command{
	Use: "crash",
	RunE: func(c *cobra.Command, args []string) error {
		return c.Usage()
	},
}

func init() {
	RootCmd.AddCommand(cmd.Start)
	RootCmd.Flags().BoolVarP(&versionFlag, "version", "v", false, "show version")
	cmd.Start.Flags().BoolVar(&cmd.PreqOrderFlag, "preq-order", false, "start modules in the order of prerequisites")
	cmd.Start.Flags().BoolVar(&cmd.ConsoleOutputFlag, "console-output", false, "print the module's output to the console")
}

func main() {
	if err := RootCmd.Execute();err != nil {
		fmt.Printf("%+v\n",err)
		os.Exit(1)
	}
}


