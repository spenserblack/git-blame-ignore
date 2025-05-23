# `git blame-ignore`

Manages your blame ignore revs file
([typically `.git-blame-ignore-revs`][github-docs]).

This will use the filename defined in the `blame.ignoreRevsFile` git configuration key,
falling back to `.git-blame-ignore-revs` as a default.

## Why?

I kept calling the same commands over and over to add a new rev to
`.git-blame-ignore-revs`, and it was getting repetitive. This is a slight shortcut.

## Basic usage

```shell
# Add the latest commit hash
git blame-ignore add

# Add a commit hash for a specific ref
git blame-ignore add REF
```

## Installation

See [`INSTALLATION.md`][installation]

[github-docs]: https://docs.github.com/en/repositories/working-with-files/using-files/viewing-and-understanding-files#ignore-commits-in-the-blame-view
[installation]: ./INSTALLATION.md
