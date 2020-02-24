package cmd

import (
	"fmt"

	"github.com/hutsharing/krait/http"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve HTTP requests",
	Long:  `Start HTTP server and listen to incoming requests.`,
	Run:   serve,
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Add debugging mode flag
	serveCmd.Flags().Bool("debug", true, "Enable debugging mode")
	viper.GetViper().BindPFlag("debug", serveCmd.Flags().Lookup("debug"))

	// Add port flag
	serveCmd.Flags().UintP("port", "p", 9000, "Port to run server on.")
	viper.GetViper().BindPFlag("port", serveCmd.Flags().Lookup("port"))

	// add host name flag
	serveCmd.Flags().String("host", "127.0.0.1", "Listen on localhost")
	viper.GetViper().BindPFlag("host", serveCmd.Flags().Lookup("host"))

}

func serve(cmd *cobra.Command, args []string) {
	server := http.NewServer()
	server.Run()
	fmt.Println("Running server...")
}
