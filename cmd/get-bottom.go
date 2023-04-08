/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/ahmedsat/go-clip/client"
	"github.com/spf13/cobra"
)

// getBottomCmd represents the add command
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getBottomCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getBottomCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
