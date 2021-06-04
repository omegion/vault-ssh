<h1 align="center">
Vault Signed SSH Certificate Manager
</h1>

<p align="center">
  <a href="https://vault-ssh.omegion.dev" target="_blank">
    <img width="180" src="https://vault-ssh.omegion.dev/img/logo.svg" alt="logo">
  </a>
</p>

<p align="center">
    <img src="https://img.shields.io/github/workflow/status/omegion/vault-ssh/Code%20Check" alt="Check"></a>
    <img src="https://coveralls.io/repos/github/omegion/vault-ssh/badge.svg?branch=master" alt="Coverall"></a>
    <img src="https://goreportcard.com/badge/github.com/omegion/vault-ssh" alt="Report"></a>
    <a href="http://pkg.go.dev/github.com/omegion/vault-ssh"><img src="https://img.shields.io/badge/pkg.go.dev-doc-blue" alt="Doc"></a>
    <a href="https://github.com/omegion/vault-ssh/blob/master/LICENSE"><img src="https://img.shields.io/github/license/omegion/vault-ssh" alt="License"></a>
</p>

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

## Requirements

* Vault Server

## What does it do?

It's a tool to create Signed SSH Certificates with Vault.

## How to use it

1. Enable a SSH engine in your Vault.

```shell
vault-ssh enable --path my-ssh-signer
```

2. Generate a Certificate CA for the engine.

```shell
vault-ssh certificate create --engine my-ssh-signer
```

3. Read created certificate to put on your server.

```shell
vault-ssh certificate get --engine my-ssh-signer
```

4. Create a role for the engine.

```shell
vault-ssh role create --name omegion --engine my-ssh-signer
```

5. Sign your public key with a role. The generated file will be written in `signed-key.pub` in this example.

```shell
vault-ssh sign \
  --role omegion \
  --engine my-ssh-signer \
  --public-key ~/.ssh/id_rsa.pub > signed-key.pub
```

6. SSH your server with signed key.

```shell
ssh -i signed-key.pub -i ~/.ssh/id_rsa root@1.1.1.1
```

## Improvements to be made

* 100% test coverage.
* Better covering for other features.