# Credential Helper

## Type of Helpers

### Built-in
Built-in credential managers come with git by default. These do not need to be installed separately.

```sh
$ ll /usr/lib/git-core/git-credential*
lrwxrwxrwx 1 root root 3 Apr 21 08:19 /usr/lib/git-core/git-credential -> git
lrwxrwxrwx 1 root root 3 Apr 21 08:19 /usr/lib/git-core/git-credential-cache -> git
lrwxrwxrwx 1 root root 3 Apr 21 08:19 /usr/lib/git-core/git-credential-cache--daemon -> git
lrwxrwxrwx 1 root root 3 Apr 21 08:19 /usr/lib/git-core/git-credential-store -> git
```

#### Store
Credential is saved on disk in a plaintext file permanently (until you change/remove them).

#### Cache
Credenatial is kept in memory for a limited time and not written to disk.

### External Program

TBA

## Config

Add the below snippet to gitconfig.  
With `username` set, the authenticate user is automatically set.  
The value for `helper` must not be quoted.  
The URL after `credential` is optional and if specified, the cache is used only for the repositories that matches the URL.

```gitconfig
[credential "https://github.com/Alice"]
username = Alice
helper = cache --timeout=604800 # 1 week
```

Multiple settings can be added for different paths.

```gitconfig
; `Alice` will be the authenticate user for repositories under `foo`, and `Bob` for bar
[credential "https://github.com/foo"]
username = Alice
helper = cache --timeout=604800

[credential "https://github.com/bar"]
username = Bob
helper = cache --timeout=604800
```

By default, credential manager is given only remote's host name and user (e.g. `https://Alice@github.com`) and uses it as key for the cache[^1], so if you need to use different account for different repositories with the same host (e.g. you have personal GitHub repo and enterprise GitHub repo for your work), credential manager cannot distinguish them and saves/returns a credential for only one of those account.  
There are 2 ways to resolve this.

[^1]: https://github.com/git-ecosystem/git-credential-manager/blob/main/docs/configuration.md#credentialusehttppath

### Set different git user for different repoistory

Credential helper include user in the cache keys, so if the appropriate user is set for each repositories, credential manager does not mistake credential for the repository.

Here is an example.
Suppose you have personal and work repositories stored on your machine as below.

```
~/workspace/personal/project-a
~/workspace/personal/project-b
~/workspace/work/project-x
~/workspace/work/project-y
```

Now add the below to global `gitconfig`.

```gitconfig
[includeIf "gitdir:~/workspace/personal"]
    path = ~/personal.gitconfig

[includeIf "gitdir:~/workspace/work"]
    path = ~/work.gitconfig
```

Then add below files.

_~/personal.gitconfig_
```gitconfig
[user]
    name = Alice personal
    email = alice-personal@users.noreply.github.com
```

_~/work.gitconfig_
```gitconfig
[user]
    name = Alice at work
    email = alice-at-work@users.noreply.github.com
```

With this configuration, when you authenticate in the repositories under `~/workspace/personal`, the user will be automatically `Alice personal` and under `~/workspace/work`, it will be `Alice at work`.  
Credential manager will use the user name to retrieve appropriate credential.

### Use credential.useHttpPath config

By setting `credential.useHttpPath` to `true`, credential manager is now provided the entire URL of the repo (e.g. `https://Alice@github.com/Alice/candyjar.git`) and uses it the key for cache, so it can distinguish the different repositories of the same host.

The below examples should give you some idea about how `useHttpPath` changes the cache key.

With `credential.useHttpPath = false`, the entire URL with user is displayed.
```
$ git push
Password for 'https://Alice@github.com/Alice/candyjar.git':
```

With `credential.useHttpPath = true`, only `user@hostname` is displayed.
```
$ git push
Password for 'https://Alice@github.com':
```

One downside of this is that now each cache is effective only within the repository that is stored, so if you have multiple repositories that you want to authenticate with the same credential (e.g. repositories that you own), you need to enter credential for each repository.
