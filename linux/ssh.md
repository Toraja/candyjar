# SSH

## Client

### Connect without password prompt
#### keychain
1. Create a pair of RSA keys on the local machine.
    ```sh
    ssh-keygen -t <algorithm> -C <description> -f [<output dir>/]<filename>
    # e.g.
    ssh-keygen -t ed25519 -C mypc -f ~/.ssh/foobar
    ```
1. Store the private key in somewhere secure on the local machine.
1. Add the content of the public key to the file `~/.ssh/authorized_key` of the user
on the remote machine you want to connect to.  
    Or you can simply run the command below on the local machine.
    ```sh
    ssh-copy-id <user>@<host>
    ```
1. On the local machine, add the code below to bashrc (or something alike),
where `<pass-to-key>` is the absolute path or the relative path from
`~/.ssh` to the private key.
    ```sh
    eval $(keychain --agents ssh --eval <path-to-key>)
    ```
1. SSH to the server.

#### Custom ssh-askpass
1. Create a custom ssh-askpass file which has the below contents.
  ```sh
  echo <password of your key>
  ```
1. Make it executable
1. Run ssh command with the custom ssh-askpass
  ```sh
  SSH_ASKPASS_REQUIRE=force SSH_ASKPASS=<path to ssh-askpass> ssh ...
  ```

If keychain outputs warning message like below, update keychain as it is old and does not support
OpenSSH 6.8+ format.  
`Can't determine fingerprint from the following line, falling back to filename`

### Options 
Add as option with `-o` or edit `/etc/ssh/ssh_config`.

#### Connect without storing finger print  
`StrictHostKeyChecking=no`

#### Not storing host key to known_hosts file
`UserKnownHostsFile=/dev/null`

## Troubleshoot
- `WARNING: REMOTE HOST IDENTIFICATION HAS CHANGED!` is displayed and cannot proceed.
    - Remove the entry for the host from `.ssh/known_hosts` file either by:
        - running `ssh-keygen -R '<host>'`.
            - If the host's port differs from 22 `<host>` should include port number as in `[111.222.111.222]:10022`.
        - simply edit `.ssh/known_hosts` file and remove the line.
            - some hosts are encoded and hard to tell which line is for which host.

## Server (daemon)

### Setup

#### Install
```sh
## alpine linux
apk update
apk --no-cache add openssh augeas
```
`augeas` is for `augtool` which is to edit config file with ease.

#### Congifuration

##### Hostkeys
Hostkey is required to start sshd.  
(Hostkey will be stored in `known_hosts` on client side when a client accesses).  
To create one, run the following command as it exactly is.  
```sh
ssh-keygen -t rsa -f /etc/ssh/ssh_host_rsa_key -N ''
```
Password for the key must be empty or sshd complains as below:
```
sshd: no hostkeys available -- exiting.
```

##### sshd_config
By default, logging in as non-root user is possible (either with key or password).  
Logging in as root with key is allowed but not with password. To enable it, change  
`PermitRootLogin` to `yes`.  
Although login as root with key is allowed (if `PubkeyAuthentication` is `yes`),  
client's public key must be passed to the server somehow without using `ssh-copy-id`  
or `sftp` as those requires password authentication.  

Use `augtool` to add or change option.  
```sh
augtool set /files/etc/ssh/sshd_config/PasswordAuthentication no
```
(Un)commenting is not supported, so even if commented option with the same name exists,
it is ignored and new line is added to the bottom.

**PermitRootLogin**  
Default: _prohibit-password_  
Must be yes to login as root with password.  
Logging in with password when this is `prohibit-password` result in `Permission denied, please try again.`

**PasswordAuthentication**  
Default: _yes_  
Enable login with password.  

**PubkeyAuthentication**  
Default: _yes_  
Setting this to `yes` allows logging in with key.  
If a client accesses using key when this is set to `no`, the client will be prompted to  
enter password.

**AuthorizedKeysFile**  
Add clients' public key to any file specified here.  
Centralised location such as `/etc/ssh/authorized_keys/%u` (`%u` means login user name).

#### ssh login user
Add a user and group for the user.  
User added by `adduser --disabled-password` will be locked so you need to unlock it  
by running `password -u ${user}`.  
Or, add a user with `useradd` command with option `-p '*'`.  
In this way, the user will not be locked.  
System user cannot be logged in, so do not add `--system` option.

#### authorized_keys
Add clients' public key to one of files specified at `AuthorizedKeysFile` in `sshd_config`.  
The contents of authorized_keys is simply copies of clients public key.

## Port Forwarding
- `AllowTcpForwarding` in `sshd_config` must be set to `yes`.
- The client has to have private key of both jump server and destination server  
  if you are using key authentication.
- Add option `-J <user>@<jump host>` to command or the snippet below to `.ssh/config` in order to  
  access to destination without manually accessing to jump server.
  ```
  Host <destination host>
    ProxyJump <jump server host>
  ```

## X11 Forwarding

### Setup

#### Client

##### Packages
- xauth

##### SSH Configuration
**ssh_config**  
- ForwardX11 yes  
  (or add `-X` option to ssh command)

##### Others
- `DISPLAY` environment variable is set

#### Server

##### Packages
- xauth
- any GUI app (e.g. `xeyes`)

##### SSH Configuration
**sshd_config**  
- X11Forwarding yes
- X11UseLocalhost no

### Via jump server
- Origin and destination server must be set up as illustrated above.
- Connect to destination server via jump server with port forwarding.
- `X11Forwarding` does not have to be enabled on jump server.
