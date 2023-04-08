/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/ahmedsat/go-clip/client"
	"github.com/spf13/cobra"
)

// removeBottomCmd represents the add command
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeBottomCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeBottomCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
