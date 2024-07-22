# 2g

## Description

A custom git wrapper that extends Git with some custom commands that I find
useful.

Any commands that are not a custom command is automatically passed to Git.
Git should have full functionality (including using vim to edit commit message).

## Custom Commands

- `2g explore <url>`
    Allows you to explore a Git repo without polluting your directories.

- `2g install <url>`
    Clones a repository in `~/.local/bin` and adds the repository to your `PATH`
    environment variable.
    Note: You will have to reload your shell to use any programs/scripts that were
    installed.
