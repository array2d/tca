package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "tca",
	Short: "table control all",
}

func init() {
	RootCmd.AddCommand(
		apiserver,
	)
}
