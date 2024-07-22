# 2g

## Description

A custom git wrapper that extends Git with some custom commands that I find
useful.

Any commands that are not a custom command is automatically passed to Git.
Git should have full functionality (including using vim to edit commit message).

## Custom Commands

- `2g explore <repo.git>`
    Allows you to explore a Git repo without polluting your directories.

- `2g install <repo.git>`
    Clones a repository in `~/.local/bin` and adds the repository to your `PATH`
    environment variable.
    Note: You will have to reload your shell to use any programs/scripts that
    were installed.

- `2g clone-file <repo.git>/<file_path>`
    Similar to the patched clone command, using `clone-file` allows you to clone
    a single file from a git repository. Without any Git history attached.

- `2g cache-clone <repo.git>`
    A more optimized version of `git clone` that will attempt to use a local
    cache of the repository and update it instead of cloning an entire repo
    from scratch.

## Patched Commands

- `2g clone <repo.git>/[file_path]`
    By adding a file path to a clone command, you are able to clone a single
    file from a git repository.
    Warning: This file will not have any Git history attached to it because it
    was not cloned as part of a repository.
    Example: Clone the CUDA .gitignore file into your repository you would
    run `2g clone https://github.com/github/gitignore/blob/main/CUDA.gitignore`
