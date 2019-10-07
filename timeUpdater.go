package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"time"

	"vasanthdeveloper.com/samaya/logger"

	"vasanthdeveloper.com/samaya/constants"
	"vasanthdeveloper.com/samaya/specific"
)

// Start is the function that is trigger by the command line parser cobra
func Start(arguments constants.ArgumentSkeleton) {

	var resp *http.Response

	// Check if a delay is provided, if yes then wait the the number of seconds
	if arguments.Delay > 0 {
		time.Sleep(time.Duration(arguments.Delay) * time.Second)
	}

	// Check if a timezone was provided
	if arguments.Timezone != "auto" {
		endpointString := constants.TimezoneEndPoint + arguments.Timezone
		resp = sendHTTPRequest(endpointString, arguments, false)
	} else {
		resp = sendHTTPRequest(constants.AutoEndPoint, arguments, false)
	}

	// Convert the response's Body into a string
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// Tell the user we were unable to get the response from server
		logger.Fatal("Failed to get response from the server.")

		os.Exit(2)
	}
	strData := string(data)

	// Tell the user we have started to parse the JSON string
	logger.Info("Parsing response body")

	// Parse the string into a JSON
	var result map[string]interface{}
	err = json.Unmarshal([]byte(strData), &result)
	if err != nil {
		// Tell the user that we were unable to parse server's response
		logger.Fatal("Failed parsing the server response.")

		os.Exit(3)
	}

	// Check if there is an error
	if result["error"] != nil {
		switch errorType := result["error"]; errorType {
		case "unknown location":
			logger.Fatal("Invalid timezone provided")
			os.Exit(9)
		default:
			logger.Fatal("The server responded with an error")
			os.Exit(8)
		}
	}

	if runtime.GOOS == "linux" {
		// As Linux time can be set using unix time, no need to parse and format the time
		specific.SetLinuxTime(result["utc_datetime"].(string), result["timezone"].(string), arguments)
	} else if runtime.GOOS == "windows" {
		// Windows requires datetime and not utc_datetime
		specific.SetWindowsTime(result["datetime"].(string), arguments)
	}
}

// sendHTTPRequest is the function that sends a HTTP GET request and returns the body
func sendHTTPRequest(url string, arguments constants.ArgumentSkeleton, recursing bool) *http.Response {
	// Tell the user that we started sending HTTP request
	// only once, the first time
	if recursing == false {
		logger.Info("Sending HTTP Request to: " + url)
	}

	// Send a GET HTTP request to our endpoint to get the current time
	response, err := http.Get(url)
	if err != nil {
		// Check if we have to wait until internet comes
		if arguments.WaitForInternet == false {
			// Tell the user that we were unable to send an HTTP request to the API
			logger.Fatal("Failed to send HTTP GET request to: " + url)
			os.Exit(1)
		}

		// Loop this function until internet comes
		return sendHTTPRequest(url, arguments, true)
	}

	return response
}
