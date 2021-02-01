package main

import (
	"github.com/briangreenhill/blackbox/pkg/cmd"
)

const URL = "https://www.google.com"
const ScheduleFormat = "* * * * *"

func main() {
	// cmd.Run("Google", URL)
	cmd.RunSchedule("Google", URL, ScheduleFormat)
}
