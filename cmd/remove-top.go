/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/ahmedsat/go-clip/client"
	"github.com/spf13/cobra"
)

// removeTopCmd represents the add command
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeTopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeTopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
