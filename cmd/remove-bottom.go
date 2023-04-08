package cmd

import (
	"github.com/ahmedsat/go-clip/client"
	"github.com/spf13/cobra"
)

var removeBottomCmd = &cobra.Command{
	Use:   "remove-bottom",
	Short: "remove the bottom of clipboard",
	Long:  `remove the value on the bottom of clipboard`,
	Run: func(cmd *cobra.Command, args []string) {

		client.RemoveBottom("unix", sockAddr, args)
	},
}

func init() {
	rootCmd.AddCommand(removeBottomCmd)

}
