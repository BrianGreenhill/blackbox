package main

import (
	"github.com/briangreenhill/blackbox/pkg/cmd"
	"github.com/robfig/cron/v3"
)

const URL = "https://www.google.com"
const ScheduleFormat = "* * * * *"

func main() {
	cron.New()
	cmd.Run("Google", URL)
}
