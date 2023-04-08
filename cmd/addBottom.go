package cmd

import (
	"fmt"
	"os"

	"github.com/ahmedsat/go-clip/client"
	"github.com/spf13/cobra"
)

var addBottomCmd = &cobra.Command{
	Use:   "add-bottom",
	Short: "add to the bottom of clipboard",
	Long:  `add new value to clipboard`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			fmt.Println("no data to add")
			os.Exit(1)
		}

		client.AddBottom("unix", sockAddr, args)
	},
}

func init() {
	rootCmd.AddCommand(addBottomCmd)

}
