# Get Started

## Prerequisites

- Vault Server

## Install

```shell
go get -u github.com/omegion/vault-ssh
```

This will install `vault-ssh` binary to your `GOPATH`.

Let's verify that the binary has installed successfully.

```shell
CLI command to manage SSH connections with Vault

Usage:
  vault-ssh [command]

Available Commands:
  certificate Manages certificates for SSH engine.
  enable      Enables SSH Engine.
  help        Help about any command
  role        Manages roles for SSH engine.
  sign        Signs given public key with SSH engine and role.
  version     Print the version/build number

Flags:
  -h, --help   help for vault-ssh

Use "vault-ssh [command] --help" for more information about a command.
```
