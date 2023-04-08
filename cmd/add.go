package cmd

import (
	"fmt"
	"os"

	"github.com/ahmedsat/go-clip/client"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add to clipboard",
	Long: `add new value to clipboard
	works just as add-top`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			fmt.Println("no data to add")
			os.Exit(1)
		}

		client.AddTop("unix", sockAddr, args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
