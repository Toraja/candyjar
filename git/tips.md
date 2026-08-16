# Tips

## Reduce the number of password prompts with http protocol

If your environment does not allow SSH authentication and you must use http, use one of the below methods to reduce the number of password prompts.

### netrc

This is most simple and straitforward, but leaves your credential as plaintext.

### Credentil helper

See [credential-helper.md](./credential-helper.md) for the detail.

```gitconfig
[credential "https://github.com/Toraja"]
username = Toraja
helper = cache --timeout=604800 # 1 week
```

### GIT_ASKPASS

TBA (This is just like `SSH_ASKPASS`)

