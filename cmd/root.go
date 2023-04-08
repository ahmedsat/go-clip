package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var sockAddr = "/tmp/rpc.sock"

var rootCmd = &cobra.Command{
	Use:   "go-clip",
	Short: "A simple unix clipboard display agnostic",
	Long:  `A simple unix clipboard display agnostic`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.PersistentFlags().StringVarP(&sockAddr, "address", "a", sockAddr, "path to network socket : default \"/tmp/rpc.sock\"")

}
