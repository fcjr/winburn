package cmd

import (
	"github.com/spf13/cobra"
)

func checkError(err error) {
	cobra.CheckErr(err)
}
