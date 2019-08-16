package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"runtime"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// Start is the function that is trigger by the command line parser cobra
func Start(cmd *cobra.Command, args []string) {
	// Send a GET HTTP request to our endpoint to get the current time
	response, err := http.Get(EndPoint)
	if err != nil {
		panic(err)
	}

	// Convert the response's Body into a string
	data, _ := ioutil.ReadAll(response.Body)
	strData := string(data)

	// Parse the string into a JSON
	var result map[string]interface{}
	json.Unmarshal([]byte(strData), &result)

	// Parse the datetime string into time object
	time, err := time.Parse(time.RFC3339, result["utc_datetime"].(string))
	if err != nil {
		panic(err)
	}

	if runtime.GOOS == "linux" {
		// As Linux time can be set using unix time, no need to parse and format the time
		SetLinuxTime(strconv.FormatInt(time.Unix(), 10), result["timezone"].(string))
	}
}
