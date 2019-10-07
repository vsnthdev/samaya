// Package constants holds some constant variables that are used across the project
package constants

// ArgumentSkeleton is the structure where all of the relevant command line arguments are stored by the cobra package
type ArgumentSkeleton struct {
	Verbose         bool
	DryRun          bool
	Version         bool
	WaitForInternet bool
	Timezone        string
	Delay           int
}

// AutoEndPoint is the variable that will contain the URL at which we send a HTTP
// GET request to get the time
var AutoEndPoint = "https://worldtimeapi.org/api/ip"

// TimezoneEndPoint is the variable that will contain the URL to which we append the timezone
// And send a GET HTTP request to get time
var TimezoneEndPoint = "https://worldtimeapi.org/api/timezone/"
