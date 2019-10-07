package main

import (
	"fmt"
	"os"

	"github.com/gosuri/uitable"

	"vasanthdeveloper.com/samaya/constants"
	"vasanthdeveloper.com/samaya/logger"

	"github.com/spf13/cobra"
)

var arguments = constants.ArgumentSkeleton{}

var rootCmd = &cobra.Command{
	Use:   "samaya",
	Short: "samaya",
	Long: `samaya
A time synchronization program that uses HTTP protocol.`,
	Run: start,
}

func main() {
	rootCmd.Flags().BoolVarP(&arguments.Verbose, "verbose", "v", false, "Show extended output")
	rootCmd.Flags().BoolVarP(&arguments.DryRun, "dry", "d", false, "Fetch the time, but don't update it")
	rootCmd.Flags().BoolVarP(&arguments.Version, "version", "V", false, "Print the version number and exit")
	rootCmd.Flags().IntVarP(&arguments.Delay, "delay", "D", 0, "Delay the number of seconds before requesting the time")
	rootCmd.Flags().BoolVarP(&arguments.WaitForInternet, "wait for internet", "w", false, "Wait until an active internet connection is established.")

	rootCmd.Flags().StringVarP(&arguments.Timezone, "timezone", "t", "auto", "Set time of that timezone")
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

func start(cmd *cobra.Command, args []string) {
	// This allows the logger to only log messages in verbose mode
	logger.Arguments = arguments
	logger.Build = Build
	logger.Version = Version

	// If the user passed the version flag, then show the version information
	printVersion()

	// Print the app name & version
	logger.PrintApp()

	// Start the start function
	Start(arguments)
}

func printVersion() {
	if arguments.Version == true {
		table := uitable.New()
		table.AddRow("Name:", "samaya")
		table.AddRow("Version:", Version)
		table.AddRow("License:", "GPL-v2.0")
		table.AddRow("Build:", Build)
		table.AddRow("Branch:", GitBranch)
		table.AddRow("Commit:", CommitHash)
		table.AddRow("Build By:", Username+"@"+Hostname)
		table.AddRow("Build On:", BuildTime)
		table.AddRow("Build Kernel:", BuildOS+" "+Kernel)

		fmt.Println(table)
		os.Exit(0)
	}
}
