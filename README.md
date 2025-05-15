# `git blame-ignore`

Manages your blame ignore revs file
([typically `.git-blame-ignore-revs`][github-docs]).

This will use the filename defined in the `blame.ignoreRevsFile` git configuration key,
falling back to `.git-blame-ignore-revs` as a default.

## Why?

I keep calling this over and over:

```shell
git log -1 --format="%H" >> .git-blame-ignore-revs
```

So I figured I'd write this to do that for me. Feel free to just copy and paste that
command if that works better than installing this executable for you ;)

## Basic usage

```shell
# Add the latest commit hash
git blame-ignore add

# Add a commit hash for a specific ref
git blame-ignore add REF
```

[github-docs]: https://docs.github.com/en/repositories/working-with-files/using-files/viewing-and-understanding-files#ignore-commits-in-the-blame-view
