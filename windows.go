package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/vasanthdeveloper/samaya/constants"
)

// SetWindowsTime is the function that sets the time for a Linux operating system
func SetWindowsTime(dateTime string, arguments constants.ArgumentSkleton) {
	// Tell the user that we are parsing the time
	if arguments.Verbose == true {
		log.Println("Parsing the time string from server.")
	}

	// Parse the datetime string into time object
	time, err := time.Parse(time.RFC3339, dateTime)
	if err != nil {
		// Tell the user that we were unable to parse Windows time
		if arguments.Verbose == true {
			log.Fatalln("Failed to parse time string from the server.")
		}

		os.Exit(4)
	}

	// Format the time according to Windows Powershell
	formattedTime := time.Format("02 January 2006 03:04:05 PM")

	// Set the time and date on Windows
	var outBytes, errBytes bytes.Buffer
	setDate := exec.Command("powershell", "Set-Date", "\""+formattedTime+"\"")
	setDate.Stdout = &outBytes
	setDate.Stderr = &errBytes

	// Tell the user that we have started to set the time
	if arguments.Verbose == true {
		log.Println("Setting time.")
	}
	// Run the above created command
	err = setDate.Run()
	if err != nil {
		// Tell the user that we were unable to set the time
		if arguments.Verbose == true {
			log.Fatalln("Failed to set time.")
		}

		os.Exit(5)
	}
}
