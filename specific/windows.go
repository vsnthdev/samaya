package specific

import (
	"os"
	"os/exec"
	"time"

	"github.com/vasanthdeveloper/samaya/constants"
	"github.com/vasanthdeveloper/samaya/logger"
)

// SetWindowsTime is the function that sets the time for a Linux operating system
func SetWindowsTime(dateTime string, arguments constants.ArgumentSkleton) {
	// Tell the user that we are parsing the time
	logger.Info("Parsing the time string from server.")

	// Parse the datetime string into time object
	time, err := time.Parse(time.RFC3339, dateTime)
	if err != nil {
		// Tell the user that we were unable to parse Windows time
		logger.Fatal("Failed to parse time string from the server.")

		os.Exit(4)
	}

	// Format the time according to Windows Powershell
	formattedTime := time.Format("02 January 2006 03:04:05 PM")

	// Set the time and date on Windows
	setDate := exec.Command("powershell", "Set-Date", "\""+formattedTime+"\"")

	// Tell the user that we have started to set the time
	logger.Info("Setting system time")
	logger.Command("powershell Set-Date \"" + formattedTime + "\"")

	// Run the above created command
	err = setDate.Run()
	if err != nil {
		// Tell the user that we were unable to set the time
		logger.Fatal("Failed to set system time")

		os.Exit(5)
	}

	// Print how much time it took to execute this command
	logger.PrintElapsedTime()
}
