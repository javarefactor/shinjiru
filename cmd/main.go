package main

import (
	"os"

	"github.com/kenjitheman/shinjiru/cli"
)

func main() {
	defer os.Exit(0)

	cmd := cli.CommandLine{}
	cmd.Run()
}
