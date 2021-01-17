package main

import "greenhill/backend/pkg/cmd"

const URL = "https://www.google.com"

func main() {
	cmd.Run("Google", URL)
}
