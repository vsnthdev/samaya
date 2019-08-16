package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var arguments = ArgumentSkleton{}

var rootCmd = &cobra.Command{
	Use:   "samaya",
	Short: "A short description",
	Long: `____ ____ _  _ ____ _   _ ____ 
[__  |__| |\/| |__|  \_/  |__| 
___] |  | |  | |  |   |   |  | `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(arguments.verbose)
	},
}

func main() {
	rootCmd.Flags().BoolVarP(&arguments.verbose, "verbose", "v", false, "Show extended output")
	rootCmd.Flags().BoolVarP(&arguments.verbose, "dry", "d", false, "Fetch the time, but don't update it")
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
