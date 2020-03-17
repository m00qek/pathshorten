# pathshorten

A smaller representation of a full filepath.

### Instalation

```
go get github.com/m00qek/pathshorten
```

### Using in bash prompt

Here is a simple example on how to use it in your bash/zsh prompt:

```sh
export PS1='$(pathshorten -s $PWD) \$ '
```

### How it works?

It works similarly to the function `pathshorten` in `vim`: it shows the first
letter of each directory of the path, except the latest one. For example, if
your `$HOME` directory is `/home/quincas` and you are trying to shorten the path
of a symlink directory at `/home/quincas/.theories/humanitas`, you have the 
following options:

|                       command                       |         result        |
| --------------------------------------------------- | --------------------- |
| `pathshorten /home/quincas/.theories/humanitas`     | `~/.t/humanitas`      |
| `pathshorten -s /home/quincas/.theories/humanitas`  | `~/.t/humanitas@`     |
| `pathshorten -a /home/quincas/.theories/humanitas`  | `/h/q/.t/humanitas`   |
| `pathshorten -sa /home/quincas/.theories/humanitas` | `/h/q/.t/humanitas@`  |

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
