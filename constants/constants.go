// Package constants holds some constant variables that are used across the project
package constants

// ArgumentSkleton is the structure where all of the relavent command line arguments are stored by the cobra package
type ArgumentSkleton struct {
	Verbose bool
	DryRun  bool
	Version bool
}

// EndPoint is the variable that will contain the URL at which we send a HTTP
// GET request to get the time
var EndPoint = "https://worldtimeapi.org/api/ip"
