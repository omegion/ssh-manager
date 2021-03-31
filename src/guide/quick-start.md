# Quick Start

Bitwarden SSH Manager is written in Golang. It is small and very fast tool. You can find detailed examples how to use
it.

## Generate Test SSH Key pair

Before using your SSH keys to store in Bitwarden, let's create dummy key and test with it.

1. Open Terminal.
1. Paste the command below

```shell
ssh-keygen -t ed25519 
```

1. When you're prompted to "Enter a file in which to save the key," enter `test`.
1. At the prompt, do not type a secure passphrase.

## Add SSH Key

Let's be sure that we have previously created keys:

```shell
❯ ls -l test*
-rw-------  1 X  staff  432 Mar 30 08:38 test
-rw-r--r--  1 X  staff  112 Mar 30 08:38 test.pub
```

Now we can add them to Bitwarden.

```shell
bw-ssh add --name test --private-key test --public-key test.pub
```

## Get SSH Key

Once we have SSH key pair on Bitwarden, let's get them to our local machine.

```shell
❯ bw-ssh get --name test
SSH Key test added.
```

Let's check `~/.ssh/keys` folder if our keys are added.

```shell
❯ ls -l ~/.ssh/keys/
-rw-------  1 X  staff   432 Mar 30 11:05 test
-rw-------  1 X  staff   112 Mar 30 11:05 test.pub
```

## Session Duration

After a login with Bitwarden CLI tool, it will return a `session key` that you will need to define it as environment
variable. Otherwise it will keep asking you to enter your credentials all the time. You can read for more info
at [Bitwarden documentation](https://bitwarden.com/help/article/cli/#environment-variable).