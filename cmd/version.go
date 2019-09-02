package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "1.1.0"

// NewCmdVersion returns the "ssh-config-gen version" command
func NewCmdVersion() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "print cli version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("version %s\n", Version)
		},
	}
}
