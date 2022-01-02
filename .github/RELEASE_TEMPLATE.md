## Installation

You can use `go` to build SSH Manager locally with:

```shell
go get -u github.com/omegion/ssh-manager
```

Or, you can use the usual commands to install or upgrade:

On OS X

```shell
sudo curl -fL https://github.com/omegion/ssh-manager/releases/download/{{.Env.VERSION}}/ssh-manager-darwin-amd64 -o /usr/local/bin/ssh-manager \
&& sudo chmod +x /usr/local/bin/ssh-manager
```

On Linux

```shell
sudo curl -fL https://github.com/omegion/ssh-manager/releases/download/{{.Env.VERSION}}/ssh-manager-linux-amd64 -o /usr/local/bin/ssh-manager \
&& sudo chmod +x /usr/local/bin/ssh-manager
```

On Windows (Powershell)

```powershell
Invoke-WebRequest -Uri https://github.com/omegion/ssh-manager/releases/download/{{.Env.VERSION}}/ssh-manager-windows-amd64 -OutFile $home\AppData\Local\Microsoft\WindowsApps\ssh-manager.exe
```

Otherwise, download one of the releases from the [release page](https://github.com/omegion/ssh-manager/releases/)
directly.

See the install [docs](https://ssh-manager.omegion.dev) for more install options and instructions.

## Changelog
