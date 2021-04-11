# Quick Start

Vault is popular secret manager from Hashicorp.

## Create SSH engine and role.

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
vault-ssh certificate read --engine my-ssh-signer
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