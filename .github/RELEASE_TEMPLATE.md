## Installation

You can use `go` to build SSH Manager locally with:

```shell
go get -u github.com/omegion/ssh-manager
```

Or, you can use the usual commands to install or upgrade:

On OS X

```shell
$ curl -L https://github.com/omegion/ssh-manager/releases/download/{{.Env.VERSION}}/ssh-manager-darwin-amd64 >/usr/local/bin/ssh-manager 
&& chmod +x /usr/local/bin/ssh-manager
```

On Linux

```shell
$ curl -L https://github.com/omegion/ssh-manager/releases/download/{{.Env.VERSION}}/ssh-manager-linux-amd64 >/usr/local/bin/ssh-manager 
&& chmod +x /tmp/ssh-manager &&
    sudo cp /tmp/ssh-manager /usr/local/bin/ssh-manager
```

Otherwise, download one of the releases from the [release page](https://github.com/omegion/ssh-manager/releases/)
directly.

See the install [docs](https://ssh-manager.omegion.dev) for more install options and instructions.

## Changelog
