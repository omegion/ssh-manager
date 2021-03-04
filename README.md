# Bitwarden SSH Manager

[![GithubBuild](https://img.shields.io/github/workflow/status/omegion/bw-ssh/Code%20Check)](http://pkg.go.dev/github.com/omegion/bw-ssh)
[![Coverage Status](https://coveralls.io/repos/github/omegion/bw-ssh/badge.svg?branch=master)](https://coveralls.io/github/omegion/bw-ssh?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/omegion/bw-ssh)](https://goreportcard.com/report/github.com/omegion/bw-ssh)
[![GoDoc](https://img.shields.io/badge/pkg.go.dev-doc-blue)](http://pkg.go.dev/github.com/omegion/bw-ssh)

```shell
CLI command to manage SSH keys stored on Bitwarden

Usage:
  bw-ssh [command]

Available Commands:
  add         Add SSH key to Bitwarden.
  get         Get SSH key from Bitwarden.
  help        Help about any command
  version     Print the version/build number

Flags:
  -h, --help   help for bw-ssh

Use "bw-ssh [command] --help" for more information about a command.

```

## Requirements

* Have the [Bitwarden CLI tool](https://github.com/bitwarden/cli) installed and available in the `$PATH` as `bw`.
* Have the `ssh-agent` running in the current session.

## What does it do?

Injects SSL keys to `ssh-agent` stored in Bitwarden.

## How to use it

1. Login to Bitwarden with `bw`.
1. Create a folder named `SSHKeys` folder in your Bitwarden.
1. Add your key pairs to Bitwarden

```shell
bw-ssh add --name my-server-1 --private-key $PK_PATH --public-key $PUB_KEY
```

## Improvements to be made

* 100% test coverage.
* Better covering for other features.

