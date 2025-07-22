# 2g

## Description

A custom git wrapper that extends Git to improve performance and provide custom
commands.

Any commands that are not a custom command is automatically passed to Git.
Git should have full functionality (including using vim to edit commit message).

## Cache Cloning

When cloning a repository, its contents is automatically cached, meaning that if
you clone the same repository again, we can simply update the locally cached
version instead of re-downloading the entire repository again.

## Lazy Commit History

In large projects, most of the time is spent downloading commit history rather
than files.

To fix this, 2g doesn't download commit history on the initial clone.

Instead, the repository will be bare cloned, and the git history will be synced
in the background by a daemon.
This means that you can clone a repository, make some changes, and (hopefully)
by the time you are ready to push, the 2g daemon will have finished generating
the commit history.

This is similar to how [sapling-scm](https://sapling-scm.com) lazily generates
its commit graph.

## Custom Commands

| Command              | Description                                                                                                                                                                              |
| -------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `install <repo.git>` | Clones a repository in ~/.local/bin and adds the repository to your PATH environment variable. Note: You will have to reload your shell to use any programs/scripts that were installed. |
