// Package logger creates structured and colored log messages
package logger

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"time"

	"github.com/bclicn/color"
	"github.com/vasanthdeveloper/samaya/constants"
)

// Arguments is the variable were the arguments parsed by cobra module
// will be passed on to the logger module
var Arguments = constants.ArgumentSkleton{}

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
			fmt.Println("samaya   v0.0.0 [Development]")
		} else {
			fmt.Println("samaya   v0.0.0 " + color.DarkGray("[Development]"))
		}
	}
}

// PrintVersion is the function that is going to output certain information
func PrintVersion() {
	if Arguments.Version == true {
		// [TODO]: Show the compiled kernel
		// [TODO]: Show the compiled machine name
		// [TODO]: Show the compiled time
		// [TODO]: Show the compiled license
		// [TODO]: Show the compiled Git branch
		// [TODO]: Show the compiled Git version
		// [TODO]: Show the compiled Git commit hash [short]
		os.Exit(0)
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
