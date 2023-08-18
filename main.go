package main

import (
	"os"

	"main.go/cli"
)

func main() {
	defer os.Exit(0)

	cmd := cli.CommandLine{}
	cmd.Run()
}
