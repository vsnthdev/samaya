package specific

import (
	"os/exec"
	"strconv"
	"time"

	"github.com/vasanthdeveloper/samaya/constants"
)

// SetLinuxTime is the function that sets the time for a Linux operating system
func SetLinuxTime(utcDatetime string, timezone string, arguments constants.ArgumentSkleton) {
	// Parse the datetime string into time object
	time, err := time.Parse(time.RFC3339, utcDatetime)
	if err != nil {
		panic(err)
	}

	setDate := exec.Command("date", "-s @"+strconv.Itoa(int(time.Unix())))
	setTimezone := exec.Command("timedatectl", "set-timezone", timezone)

	// Set the timezone first, and then set the date
	err = setTimezone.Run()
	if err != nil {
		panic(err)
	}
	err = setDate.Run()
	if err != nil {
		panic(err)
	}
}
