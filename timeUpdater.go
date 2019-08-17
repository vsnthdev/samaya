package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/vasanthdeveloper/samaya/constants"
)

// Start is the function that is trigger by the command line parser cobra
func Start(arguments constants.ArgumentSkleton) {

	// Tell the user that we started sending HTTP request
	if arguments.Verbose == true {
		log.Println("Sending HTTP Request to: " + constants.EndPoint + ".")
	}

	// Send a GET HTTP request to our endpoint to get the current time
	response, err := http.Get(constants.EndPoint)
	if err != nil {
		// Tell the user that we were unable to send an HTTP request to the API
		if arguments.Verbose == true {
			log.Fatalln("Failed to send HTTP GET request to: " + constants.EndPoint + ".")
		}

		os.Exit(1)
	}

	// Convert the response's Body into a string
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		// Tell the user we were unable to get the response from server
		if arguments.Verbose == true {
			log.Fatalln("Failed to get response from the server.")
		}

		os.Exit(2)
	}
	strData := string(data)

	// Tell the user we have started to parse the JSON string
	if arguments.Verbose == true {
		log.Println("Parsing response body.")
	}

	// Parse the string into a JSON
	var result map[string]interface{}
	err = json.Unmarshal([]byte(strData), &result)
	if err != nil {
		// Tell the user that we were unable to parse server's response
		if arguments.Verbose == true {
			log.Fatalln("Failed parsing the server response.")
		}

		os.Exit(3)
	}

	if runtime.GOOS == "linux" {
		// As Linux time can be set using unix time, no need to parse and format the time
		SetLinuxTime(result["utc_datetime"].(string), result["timezone"].(string), arguments)
	} else if runtime.GOOS == "windows" {
		// Windows requires datetime and not utc_datetime
		SetWindowsTime(result["datetime"].(string), arguments)
	}
}
