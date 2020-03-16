#pathshorten

A smaller representation of a full filepath.

### Using in bash prompt

Here is a simple example on how to use it in your bash/zsh prompt:

```sh
export PS1='$(pathshorten -s $PWD) \$ '
```

### How it works?

It works similarly to the function `pathshorten` in `vim`: it shows the first
letter of each directory of the path, except the latest one. For example, if
your `$HOME` directory is `/home/myuser` and you are trying to shorten the path
of a symlink directory at `/home/myuser/projects/pathshorten`, you have the 
following options:

|                       command                       |         result        |
| --------------------------------------------------- | --------------------- |
| `pathshorten /home/myuser/projects/pathshorten`     | `~/p/pathshorten`     |
| `pathshorten -s /home/myuser/projects/pathshorten`  | `~/p/pathshorten@`    |
| `pathshorten -a /home/myuser/projects/pathshorten`  | `/h/m/p/pathshorten`  |
| `pathshorten -sa /home/myuser/projects/pathshorten` | `/h/m/p/pathshorten@` |

### API

```
Prints a shortened version of a directory absolute path.

Usage:
  pathshorten [options] <path>
  pathshorten --help
  pathshorten --version

Options:
  -s --show-symlinks  Append a '@' to directories that are symbolic links.
  -a --absolute       Do not use '~' as a shortcut for $HOME.
  -h --help           Show this message.
  -v --version        Show version.
```
