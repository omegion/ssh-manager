# lpass-ssh

LastPass SSH Manager

```shell
printf "Private Key: %s\nPublic Key: %s\n" \
  "$(cat /Users/hakan/.ssh/hetzner)" "$(cat /Users/hakan/.ssh/hetzner.pub)" \
    | lpass add --non-interactive --sync=now --note-type=ssh-key "ssh-key-test-1"
```