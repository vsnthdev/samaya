package main

import (
	"os/exec"
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

	setDate := exec.Command("date", "-s @"+string(time.Unix()))
	setTimezone := exec.Command("timedatectl", "set-timezone", timezone)

	// Set the timezone first, and then set the date
	setTimezone.Run()
	setDate.Run()
}
