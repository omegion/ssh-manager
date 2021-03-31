# Get Started

## Prerequisites

Bitwarden SSH Manager uses `bw` CLI command in background. So, you will need to install it and login to be able to use the SSH Manager.

- [Bitwarden CLI](https://bitwarden.com/help/article/cli/#quick-start)

## Install

```shell
go get -u github.com/omegion/bw-ssh
```

This will install `bw-ssh` binary to your `GOPATH`.

Let's verify that the binary has installed successfully.

```shell
❯ bw-ssh --help            
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
