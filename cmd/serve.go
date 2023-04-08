package cmd

import (
	"github.com/ahmedsat/go-clip/server"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "starting the server",
	Long: `a background server which have to be running before 
	being able to execute any other command,
	notes that if server stop clip board data will be deleted`,
	Run: func(cmd *cobra.Command, args []string) {
		server.Serve(sockAddr)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

}
