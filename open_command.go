package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/google/go-github/github"
	flags "github.com/jessevdk/go-flags"
)

type OpenCommand struct {
	stdin  io.Reader
	stdout io.Writer
	stderr io.Writer
	github *github.Client
}

type options struct {
	Labels    string `short:"l" long:"labels" description:"Labels separated by ',' (comma)"`
	Assignees string `short:"a" long:"assignees" description:"Usernames of assignees separated by ',' (comma)"`
}

func (c *OpenCommand) Help() string {
	return "usage: " + AppName + " open user/repo\n\n" +
		"Available options are:\n" +
		"  -l, --labels=    Comma-separated list of labels\n" +
		"  -a, --assignees= Comma-separated list of usernames of assignees\n\n" +
		"You need to give title and comment from STDIN.\n" +
		"The first line will be the title and the rest will be the comment."
}

func (c *OpenCommand) Synopsis() string {
	return "Open a new issue"
}

func (c *OpenCommand) Run(args []string) int {
	apiToken := os.Getenv("GITHUB_API_TOKEN")
	if apiToken == "" {
		fmt.Fprintln(c.stderr, "error: you need to set GitHub's access toekn to environment variable GITHUB_API_TOKEN")
		return 1
	}

	var opts options
	optParser := flags.NewParser(&opts, flags.Default^flags.PrintErrors)
	args, err := optParser.ParseArgs(args)
	if err != nil {
		fmt.Fprint(c.stderr, "error: ")
		fmt.Fprintln(c.stderr, err)
		return 1
	}

	if len(args) < 1 {
		fmt.Fprintln(c.stderr, "error: repository is not specified")
		fmt.Fprintln(c.stderr, c.Help())
		return 1
	}

	repoInfos := strings.Split(args[0], "/")
	user := repoInfos[0]
	repo := repoInfos[1]

	if len(repoInfos) < 2 {
		fmt.Fprintln(c.stderr, "error: specified repository format is invalid")
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(c.stdin)
	s := buf.String()
	texts := strings.SplitN(s, "\n", 2)

	labels := strings.Split(opts.Labels, ",")
	assignees := strings.Split(opts.Assignees, ",")

	issueRequest := &github.IssueRequest{
		Title:     &texts[0],
		Body:      &texts[1],
		Labels:    &labels,
		Assignees: &assignees,
	}

	issue, _, err := c.github.Issues.Create(user, repo, issueRequest)

	if err != nil {
		fmt.Fprintln(c.stderr, "error: failed to create an issue")
		fmt.Fprintln(c.stderr, err)
		return 1
	}

	fmt.Fprintln(c.stderr, "succeeded to create an issue!")
	fmt.Fprintln(c.stdout, *issue.HTMLURL)
	return 0
}
