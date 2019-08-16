package main

import (
	"fmt"
	"os/exec"
)

// SetLinuxTime is the function that sets the time for a Linux operating system
func SetLinuxTime(unix string, timezone string) {
	fmt.Println("timedatectl set-timezone \"" + timezone + "\"")
	fmt.Println("date -s @" + unix)

	setDate := exec.Command("date", "-s @"+unix)
	setTimezone := exec.Command("timedatectl", "set-timezone", timezone)

	// Set the timezone first, and then set the date
	setTimezone.Run()
	setDate.Run()
}
