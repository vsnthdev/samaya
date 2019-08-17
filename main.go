package main

import (
	"github.com/vasanthdeveloper/samaya/constants"
	"github.com/vasanthdeveloper/samaya/logger"

	"github.com/spf13/cobra"
)

var arguments = constants.ArgumentSkleton{}

var rootCmd = &cobra.Command{
	Use:   "samaya",
	Short: "A short description",
	Long: `____ ____ _  _ ____ _   _ ____ 
[__  |__| |\/| |__|  \_/  |__| 
___] |  | |  | |  |   |   |  | `,
	Run: start,
}

func main() {
	rootCmd.Flags().BoolVarP(&arguments.Verbose, "verbose", "v", false, "Show extended output")
	rootCmd.Flags().BoolVarP(&arguments.DryRun, "dry", "d", false, "Fetch the time, but don't update it")
	rootCmd.Flags().BoolVarP(&arguments.Version, "version", "V", false, "Print the version number and exit")
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

func start(cmd *cobra.Command, args []string) {
	// This allows the logger to only log messages in verbose mode
	logger.Arguments = arguments

	// Print the app name & version
	logger.PrintApp()

	// If the user passed the version flag, then show the version information
	logger.PrintVersion()

	// Start the start function
	Start(arguments)
}
