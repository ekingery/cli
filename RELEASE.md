# Cutting a CLI Release

The Exercism CLI uses [GoReleaser](https://goreleaser.com) to automate the
release process. 

## Requirements

1. [Install GoReleaser](https://goreleaser.com/install/)
1. [Install snapcraft](https://snapcraft.io/docs/snapcraft-overview)
1. [Setup GitHub token](https://goreleaser.com/environment/#github-token)
1. Have a gpg key installed on your machine - it is [used for signing the artifacts](https://goreleaser.com/sign/)

## Cut a release

```bash

# Test run
goreleaser --skip-publish --snapshot --rm-dist

# Commit any changes, then create a new tag and push it
git tag -a v3.0.16 -m "Trying out GoReleaser"
git push origin v3.0.16

# Build and release
goreleaser --rm-dist

# If that errors out with a missing github token, there may be an issue where 
# goreleaser does not read the file in ~/.config (on mac only?). 
# In that case, try an `export GITHUB_TOKEN=xxx` and then re-run goreleaser

# Remember to update cmd/version.go in the code
# (until we use: https://goreleaser.com/environment/#using-the-main-version)

# You must be logged into snapcraft to publish a new snap
snapcraft login

# Push to snapcraft
for f in `ls dist/*.snap`; do snapcraft push --release=stable $f; done

# [TODO] Push to homebrew
```

Lastly, head to [the release page](https://github.com/exercism/cli/releases) to test and publish the draft.


# **** [TODO] Former content below - comb through and update ****

## Update the Changelog

Make sure all the recent changes are reflected in the "next release" section
of the Changelog. Make this a separate commit from bumping the version.

You can view changes using the /compare/ view:
https://github.com/exercism/cli/compare/$PREVIOUS_RELEASE...master


## Bump the version

Edit the `Version` constant in `cmd/version.go`, and edit the Changelog.

All the changes in the "next release" section should be moved to a new section
that describes the version number, and gives it a date.

The "next release" section should contain only "Your contribution here".

_Note: It's useful to add the version to the commit message when you bump it: e.g. `Bump version to v2.3.4`._

## Generate the Binaries

```plain
$ rm release/*
$ CGO_ENABLED=0 bin/build-all
```

## Cut Release on GitHub

Go to [the exercism/cli "new release" page](https://github.com/exercism/cli/releases/new).

Describe the release, select a specific commit to target, name the version `v{VERSION}`, where
VERSION matches the value of the `Version` constant.

Upload all the binaries from `release/*`.

Paste the release text and describe the new changes.

```
To install, follow the interactive installation instructions at https://exercism.io/cli-walkthrough

---

[describe changes in this release]

```

## Update Homebrew

This is helpful for the (many) Mac OS X users.

First, get a copy of the latest tarball of the source code:

```
cd ~/tmp && wget https://github.com/exercism/cli/archive/vX.Y.Z.tar.gz
```

Get the SHA256 of the tarball:

```
shasum -a 256 vX.Y.Z.tar.gz
```

Update the homebrew formula:

```
cd $(brew --repository)
git checkout master
brew update
brew bump-formula-pr --strict exercism --url=https://github.com/exercism/cli/archive/vX.Y.Z.tar.gz --sha256=$SHA
```

For more information see their [contribution guidelines](https://github.com/Homebrew/homebrew/blob/master/share/doc/homebrew/How-To-Open-a-Homebrew-Pull-Request-(and-get-it-merged).md#how-to-open-a-homebrew-pull-request-and-get-it-merged).

## Update the docs site

If there are any significant changes, we should describe them on
[exercism.io/cli]([https://exercism.io/cli).

The codebase lives at [exercism/website-copy](https://github.com/exercism/website-copy) in `pages/cli.md`.
