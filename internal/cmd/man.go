package cmd

import (
	"fmt"
	"os"

	"github.com/fcjr/winburn/internal/validators"
	mcobra "github.com/muesli/mango-cobra"
	"github.com/muesli/roff"
	"github.com/spf13/cobra"
)

type manCmd struct {
	cmd *cobra.Command
}

func newManCmd() *manCmd {
	return &manCmd{
		cmd: &cobra.Command{
			Use:                   "man",
			Short:                 "Generates winburn's command line manpages",
			SilenceUsage:          true,
			DisableFlagsInUseLine: true,
			Hidden:                true,
			Args:                  validators.NoArgs(),
			Run: func(cmd *cobra.Command, args []string) {
				manPage, err := mcobra.NewManPage(1, cmd.Root())
				checkError(err)

				_, err = fmt.Fprint(os.Stdout, manPage.Build(roff.NewDocument()))
				checkError(err)
			},
		},
	}
}
