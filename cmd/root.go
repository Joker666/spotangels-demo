package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd is the root of all sub commands in the binary
// it doesn't have a Run method as it executes other sub commands
var rootCmd = &cobra.Command{
	Use:   "spotangels-demo",
	Short: "spotangels-demo is a http server to serve api",
}

func init() {
	rootCmd.AddCommand(srvCmd)
}

// Execute executes the command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
