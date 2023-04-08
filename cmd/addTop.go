/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/ahmedsat/go-clip/client"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addTopCmd = &cobra.Command{
	Use:   "add-top",
	Short: "add to the top of clipboard",
	Long:  `add new value to clipboard`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			fmt.Println("no data to add")
			os.Exit(1)
		}

		client.AddTop("unix", sockAddr, args)
	},
}

func init() {
	rootCmd.AddCommand(addTopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
