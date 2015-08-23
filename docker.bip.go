package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	cmds = append(cmds, &cobra.Command{
		Use:   "docker/bip [address]",
		Short: "Get or Set Docker Bridge IP",
		Long:  "This gets or sets the IP for the internal docker bridge.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println(APIGet("docker/bip"))
			} else if len(args) == 1 {
				APIPost("docker/bip", args[0])
			} else {
				cmd.Help()
			}
		},
	})
}
