package main

import "github.com/exercism/cli/cmd"

var (
	Version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	cmd.Execute()
}
