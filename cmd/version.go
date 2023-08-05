package cmd

import (
	"fmt"

	"github.com/fcjr/winburn/validators"
	"github.com/fcjr/winburn/version"
	"github.com/spf13/cobra"
)

type versionCmd struct {
	cmd *cobra.Command
}

func newVersionCmd() *versionCmd {
	return &versionCmd{
		cmd: &cobra.Command{
			Use:   "version",
			Args:  validators.NoArgs(),
			Short: "Get the version of the winburn CLI",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println(version.String())
			},
		},
	}
}
