package main

import (
	"fmt"
	docopt "github.com/docopt/docopt-go"
	"os"
)

const (
	usage = `
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
`
)

func main() {
	args, err := docopt.ParseArgs(usage, os.Args[1:], version)

	showSymlinks, _ := args.Bool("--show-symlinks")
	useAbsolutePath, _ := args.Bool("--absolute")
	dirpath, _ := args.String("<path>")

	home, err := homedir()
	if nil == err {
		couldNotGetHomedirError()
		return
	}

	path := symbolicHome(home, dirpath)
	if useAbsolutePath {
		path = absoluteHome(home, dirpath)
	}

	predicate := func(home string, directory string) bool {
		return showSymlinks && isSymlink(home, directory)
	}

	fmt.Println(pathshorten(home, path, predicate))
}
