# GPG

## Less passphrase prompts

Use `gpg` and `gpg-agent` to reduce the need to enter your passphrase securely.

### Setup

1. Create a file that contains your passphrase
1. Encrypt it with GPG  
  This will create an encrypted file with the same name as the original file but with `.gpg` extension. You can delete the original file after encryption.
  ```bash
  gpg --symmetric --cipher-algo AES256 <PATH_TO_PASSWORD_FILE>
  ```
1. Run gpg-agent for caching the passphrase (place it in your shell profile, e.g. `~/.bashrc` or `~/.zshrc`)
  ```bash
  gpg-agent --daemon > /dev/null 2>&1
  ```
  - Optionally, you can set the cache time (in seconds) for the passphrase
    - `--default-cache-ttl`
      - Time until a cached passphrase is expired (default: 600 seconds (10 minutes))
      - It gets reset every time the passphrase is used
    - `--max-cache-ttl`
      - The absolute time until a cached passphrase is expired (default: 7200 seconds (2 hours))
      - It will be expired once the specified seconds have passed since it was first cached regardless of usage

### Usage

`gpg --decrypt <path-to-encrypted-password-file>` will prompt you for the passphrase and output the file contents.  
With gpg-agent running, you will only need to enter the passphrase once within the cache time, and subsequent calls to `gpg --decrypt` will not prompt for the passphrase until the cache expires.

_e.g._
```bash
gpg --decrypt <PATH_TO_ENCRYPTED_PASSWORD_FILE> | xargs -I {} ssh user@host -p {}
```
