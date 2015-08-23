package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	cmds = append(cmds, &cobra.Command{
		Use:   "hostname [hostname]",
		Short: "Get or Set Hostname",
		Long:  "This gets or sets the hostname for the system.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Printf("%s", APIGet("hostname"))
			} else if len(args) == 1 {
				APIPost("hostname", args[0])
			} else {
				cmd.Help()
			}
		},
	})
}
