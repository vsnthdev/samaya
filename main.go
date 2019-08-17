package main

import (
	"fmt"
	"os"

	"github.com/vasanthdeveloper/samaya/constants"

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
	if arguments.DryRun == true || arguments.Verbose == true {
		fmt.Println("samaya, v0.0.0")
	}

	// If, version was true, the next thing todo is exit
	if arguments.Version == true {
		os.Exit(0)
	}

	// Start the start function
	Start(arguments)
}
