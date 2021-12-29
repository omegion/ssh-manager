<h1 align="center">
SSH Key Manager for Bitwarden and 1Password
</h1>

<p align="center">
  <a href="https://ssh-manager.omegion.dev" target="_blank">
    <img width="180" src="https://ssh-manager.omegion.dev/img/logo.png" alt="logo">
  </a>
</p>

<p align="center">
    <img src="https://img.shields.io/github/workflow/status/omegion/ssh-manager/Code%20Check" alt="Check"></a>
    <img src="https://coveralls.io/repos/github/omegion/ssh-manager/badge.svg?branch=master" alt="Coverall"></a>
    <img src="https://goreportcard.com/badge/github.com/omegion/ssh-manager" alt="Report"></a>
    <a href="http://pkg.go.dev/github.com/omegion/ssh-manager"><img src="https://img.shields.io/badge/pkg.go.dev-doc-blue" alt="Doc"></a>
    <a href="https://github.com/omegion/ssh-manager/blob/master/LICENSE"><img src="https://img.shields.io/github/license/omegion/ssh-manager" alt="License"></a>
</p>

```shell
SSH Manager for your keys on 1Password and Bitwarden

Usage:
  ssh-manager [command]

Available Commands:
  add         Add Manager key to given provider.
  completion  Generate the autocompletion script for the specified shell
  get         Get Manager key from given provider.
  help        Help about any command
  list        List Manager keys from given provider.
  version     Print the version/build number

Flags:
  -h, --help               help for ssh-manager
      --logFormat string   Set the logging format. One of: text|json (default "text") (default "text")
      --logLevel string    Set the logging level. One of: debug|info|warn|error (default "info")

Use "ssh-manager [command] --help" for more information about a command.
```

## Installation

You can use `go` to build S3 Secrets Manager locally with:

```shell
go get -u github.com/omegion/ssh-manager
```

This will install `ssh-manager` binary to your `GOPATH`.

Or, you can use the usual commands to install or upgrade:

On OS X

```shell
VERSION=v0.3.0
$ curl -L https://github.com/omegion/ssh-manager/releases/download/$VERSION/ssh-manager-darwin-amd64 >/usr/local/bin/ssh-manager 
&& \
  chmod +x /usr/local/bin/ssh-manager
```

On Linux

```shell
VERSION=v0.3.0
$ curl -L https://github.com/omegion/ssh-manager/releases/download/$VERSION/ssh-manager-linux/amd64 >/usr/local/bin/ssh-manager 
&& \
    chmod +x /tmp/ssh-manager &&
    sudo cp /tmp/ssh-manager /usr/local/bin/ssh-manager
```

Otherwise, download one of the releases from the [release page](https://github.com/omegion/ssh-manager/releases/)
directly.

## Requirements

* Have the [Bitwarden CLI tool](https://github.com/bitwarden/cli) installed and available in the `$PATH` as `bw`.
* Or have the [1Password CLI tool](https://1password.com/downloads/command-line/) installed and available in the `$PATH`
  as `op`.
* Have the `ssh-agent` running in the current session.

## What does it do?

Injects SSL keys to `ssh-agent` stored in Bitwarden or 1Password.

## How to use it

1. Login to Bitwarden or 1Password with `bw` or `op`.
1. Add your key pairs to your password manager.

For Bitwarden
---

```shell
ssh-manager add --name my-server --private-key $PK_PATH --public-key $PUB_KEY_PATH --provider bw
```

For 1Password
---

```shell
ssh-manager add --name my-another-server --private-key $PK_PATH --public-key $PUB_KEY_PATH --provider op
```

## Improvements to be made

* 100% test coverage.
* Better covering for other features.

