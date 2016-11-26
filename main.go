package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

const AppName = "ghissue"

func main() {
	c := cli.NewCLI(AppName, Version)
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"open": func() (cli.Command, error) {
			return &OpenCommand{
				stdin:  os.Stdin,
				stdout: os.Stdout,
				stderr: os.Stderr,
				github: getGitHubClient(),
			}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	os.Exit(exitStatus)
}
