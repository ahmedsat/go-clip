/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/ahmedsat/go-clip/client"
	"github.com/spf13/cobra"
)

// getTopCmd represents the add command
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getTopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getTopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
