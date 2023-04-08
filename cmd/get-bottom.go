package cmd

import (
	"github.com/ahmedsat/go-clip/client"
	"github.com/spf13/cobra"
)

var getBottomCmd = &cobra.Command{
	Use:   "get-bottom",
	Short: "get from clipboard",
	Long: `get the value from the top of clipboard
	works just as get-top`,
	Run: func(cmd *cobra.Command, args []string) {

		client.GetBottom("unix", sockAddr, args)
	},
}

func init() {
	rootCmd.AddCommand(getBottomCmd)

}
