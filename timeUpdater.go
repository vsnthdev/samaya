package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"

	"vasanthdeveloper.com/samaya/logger"

	"vasanthdeveloper.com/samaya/constants"
	"vasanthdeveloper.com/samaya/specific"
)

// Start is the function that is trigger by the command line parser cobra
func Start(arguments constants.ArgumentSkleton) {

	var resp *http.Response

	// Check if a timezone was provided
	if arguments.Timezone != "auto" {
		endpointString := constants.TimezoneEndPoint + arguments.Timezone
		resp = sendHTTPRequest(endpointString)
	} else {
		resp = sendHTTPRequest(constants.AutoEndPoint)
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
func sendHTTPRequest(url string) *http.Response {
	// Tell the user that we started sending HTTP request
	logger.Info("Sending HTTP Request to: " + url)

	// Send a GET HTTP request to our endpoint to get the current time
	response, err := http.Get(url)
	if err != nil {
		// Tell the user that we were unable to send an HTTP request to the API
		logger.Fatal("Failed to send HTTP GET request to: " + url)

		os.Exit(1)
	}

	return response
}
