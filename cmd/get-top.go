package cmd

import (
	"github.com/ahmedsat/go-clip/client"
	"github.com/spf13/cobra"
)

var getTopCmd = &cobra.Command{
	Use:   "get-top",
	Short: "get from clipboard",
	Long:  `get the value from the top of clipboard`,
	Run: func(cmd *cobra.Command, args []string) {

		client.GetTop("unix", sockAddr, args)
	},
}

func init() {
	rootCmd.AddCommand(getTopCmd)

}
