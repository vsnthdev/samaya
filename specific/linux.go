package specific

import (
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/vasanthdeveloper/samaya/constants"
	"github.com/vasanthdeveloper/samaya/logger"
)

// SetLinuxTime is the function that sets the time for a Linux operating system
func SetLinuxTime(utcDatetime string, timezone string, arguments constants.ArgumentSkleton) {
	// Parse the datetime string into time object
	time, err := time.Parse(time.RFC3339, utcDatetime)
	if err != nil {
		// Tell the user we were unable to parse the time stamp
		logger.Fatal("Failed to parse time string from the server.")
	}

	setDate := exec.Command("date", "-s @"+strconv.Itoa(int(time.Unix())))
	setTimezone := exec.Command("timedatectl", "set-timezone", timezone)

	// Set the timezone first, and then set the date
	logger.Info("Setting timezone")
	logger.Command("timedatectl set-timezone \"" + timezone + "\"")
	err = setTimezone.Run()
	if err != nil {
		// Tell the user we failed at setting the timezone
		logger.Fatal("Failed to set system timezone")

		os.Exit(7)
	}

	logger.Info("Setting system time")
	logger.Command("date -s @" + strconv.Itoa(int(time.Unix())))
	err = setDate.Run()
	if err != nil {
		// Tell the user we failed at setting the timezone
		logger.Fatal("Failed to set system time")

		os.Exit(5)
	}

	// Print how long it took to execute the commands
	logger.PrintElapsedTime()
}
