package main

import "github.com/briangreenhill/blackbox/pkg/cmd"

const URL = "https://www.google.com"

func main() {
	cmd.Run("Google", URL)
}
