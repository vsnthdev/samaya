// Package logger creates structured and colored log messages
package logger

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"time"

	"github.com/bclicn/color"
	"vasanthdeveloper.com/samaya/constants"
)

// Arguments is the variable were the arguments parsed by cobra module
// will be passed on to the logger module
var Arguments = constants.ArgumentSkeleton{}

// Build is where the program's build type is passed from main package
var Build string

// Version is where the program's version is passed from the main package
var Version string

// ExecutionStartTime is the time variable that holds the time at which we finished parsing
// the command line arguments
var ExecutionStartTime = time.Now()

// Info function will log an info message
func Info(message string) {
	if Arguments.Verbose == true {
		if runtime.GOOS == "windows" {
			fmt.Println("info     " + message + ".")
		} else {
			fmt.Printf(color.LightBlue("info")+"     %s.\n", message)
		}
	}
}

// Command message will log a message that shows a command
func Command(command string) {
	if Arguments.Verbose == true {
		if runtime.GOOS == "windows" {
			fmt.Println("command  " + command)
		} else {
			fmt.Printf(color.LightPurple("command")+"  %s\n", command)
		}
	}
}

// Fatal logs an error message
func Fatal(message string) {
	if runtime.GOOS == "windows" {
		os.Stderr.WriteString("fatal    " + message + "\n")
	} else {
		os.Stderr.WriteString(color.LightRed("fatal") + "    " + message + ".\n")
	}
}

// PrintApp is the function that prints the app name and it's version
func PrintApp() {
	if Arguments.Verbose == true || Arguments.Version {
		if runtime.GOOS == "windows" {
			fmt.Println("samaya   v" + Version + " [" + Build + "]")
		} else {
			fmt.Println("samaya   v" + Version + " " + color.DarkGray("[") + color.DarkGray(Build) + color.DarkGray("]"))
		}
	}
}

// PrintElapsedTime is the function that prints how much time was taken
func PrintElapsedTime() {
	if Arguments.Verbose == true {
		elapsedTime := time.Now()
		elapsed := elapsedTime.Sub(ExecutionStartTime)

		if runtime.GOOS == "windows" {
			fmt.Printf("took     %0.2fs \n", roundUp(elapsed.Seconds(), 2))
		} else {
			fmt.Printf(color.Green("took")+"     %0.2fs \n", roundUp(elapsed.Seconds(), 2))
		}
	}
}

func roundUp(input float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * input
	round = math.Ceil(digit)
	newVal = round / pow
	return
}
