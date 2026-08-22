# Release

This document describes the process of releasing a crate.  
The process usually involves version management and optionally binary/package distribution.

## Version Management

This usually includes:
- Bumping the version in `Cargo.toml/lock`
- (Optionally) updating the `CHANGELOG.md` file
- Tag commit and push the changes to the remote repository

### cargo release

[cargo release](https://github.com/crate-ci/cargo-release) is a cargo subcommand that automates the version management process of releasing a crate.

> [!NOTE]
> `cargo release` publishes the crate to [crates.io](https://crates.io/) by default. Pass `--no-publish` or modify config to skip publishing.

#### Configuration

See [Configuration](https://github.com/crate-ci/cargo-release/blob/master/docs/reference.md#configuration).

#### Release steps

`cargo release` defaults to a dry-run mode, which will print the steps that would be taken to release the crate.
To make the actual changes, pass the `--execute` flag.

##### Simple

If you are releasing a crate on the current commit, just run the following command.

```sh
cargo release [level|version]
```

This will:
- Update the crate version (`Cargo.toml/lock`)
  - If the same version as crate is passed, it just skips the crate version bump and continues on the rest of steps.
- Create a git tag
- Push the changes to the remote repository
- Publish the crate to `crates.io`

When no argument is provided, the current version is used as the release version.

> [!NOTE]
> When `level` is either `rc/beta/alpha` and the current version does not have suffix, the version will be bumped to the next pre-release version in addition to adding version suffix.
For example, if the current version is `1.0.0`, running `cargo release rc` will bump the version to `1.0.1-rc.0`.
If the current version is `1.0.0-rc.0`, running `cargo release rc` will bump the version to `1.0.0-rc.1`.
This behavior is not customizable, so if you want to bump the version to `1.1.0-rc.1` from `1.0.0`, you will need to run `cargo release 1.1.0-rc.1` instead of `cargo release rc`.

##### With some chores

You might have changelog and other files the changes of which you want to include in the same commit as crate version bump.
`cargo release` command exits with an error if there are uncommitted changes, so you cannot just edit other files and run `cargo release`.
`cargo release` offers 2 functionalities for this purpose: [pre-release-replacements](https://github.com/crate-ci/cargo-release/blob/master/docs/reference.md#pre-release-replacements) and [pre-release-hooks](https://github.com/crate-ci/cargo-release/blob/master/docs/reference.md#pre-release-hook), and they can even automate the chores.

`pre-release-replacements` replaces texts (regexp) in the specified files when running `cargo release --execute` and the changes are included in the commit `cargo release` makes. This is especially useful to automate updating changelog. (This exact thing is documented [here](https://github.com/crate-ci/cargo-release/blob/master/docs/faq.md#maintaining-changelog).) Without `--execute`, it shows diff that would be made, so you can preview the changes before actually releasing.

If the text replacement is not enough, `pre-release-hooks` can run arbitrary commands. The working directory for the commands is the same directory as `Cargo.toml` (I only tested for simple crate that has single `Cargo.toml`).

## Binary/Package Distribution

If you are distributing built artifacts you need extra steps to publish them to online platforms like GitHub or GitLab.  
This usually involves:
- Building the artifacts
- Creating a release on the platform
- Uploading the artifacts to the release

### dist

[dist](https://axodotdev.github.io/cargo-dist/) generates CI/CD configuration for building and publishing binaries of a crate to GitHub/GitLab releases.

See the [documentation](https://axodotdev.github.io/cargo-dist/book/) for more details.
