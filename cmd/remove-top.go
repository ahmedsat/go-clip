package cmd

import (
	"github.com/ahmedsat/go-clip/client"
	"github.com/spf13/cobra"
)

var removeTopCmd = &cobra.Command{
	Use:   "remove-top",
	Short: "remove the top of clipboard",
	Long:  `remove the value on the top of clipboard`,
	Run: func(cmd *cobra.Command, args []string) {

		client.RemoveTop("unix", sockAddr, args)
	},
}

func init() {
	rootCmd.AddCommand(removeTopCmd)

}
