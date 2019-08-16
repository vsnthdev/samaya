package main

// ArgumentSkleton is the structural object for creating CLI arguments
type ArgumentSkleton struct {
	verbose bool
	dryRun  bool
	version bool
}

// EndPoint is the string URL from where we can fetch the current time
var EndPoint = "http://worldtimeapi.org/api/ip"
