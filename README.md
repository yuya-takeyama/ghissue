# ghissue

Command to create GitHub issues.

This aims to help automation tasks.

## Usage

### Set GitHub access token

You need to set GitHub's access token to environment variable `GITHUB_ACCESS_TOKEN`.

### Open a new issue

```
$ ghissue open user/repo

Available options are:
  -l, --labels=    Comma-separated list of labels
  -a, --assignees= Comma-separated list of usernames of assignees

You need to give title and comment from STDIN.
The first line will be the title and the rest will be the comment.

```

## Author

Yuya Takeyama

## License

The MIT License
