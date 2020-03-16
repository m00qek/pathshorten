package main

import (
	"os"
	"os/user"
	"path"
	"path/filepath"
	"strings"
)

const SEPARATOR = string(filepath.Separator)
const ROOT = string(filepath.Separator)
const SYMLINK = "@"
const HOME = "~"

func homedir() (string, error) {
	user, err := user.Current()
	if nil != err {
		return "", err
	}

	return user.HomeDir, nil
}

func symbolicHome(home string, dirpath string) string {
	return strings.Replace(dirpath, home, HOME, 1)
}

func absoluteHome(home string, dirpath string) string {
	return strings.Replace(dirpath, HOME, home, 1)
}

func shortname(text string) string {
	var runes = []rune(text)

	if string(runes[0]) == "." {
		return string(runes[:2])
	}

	return string(runes[0])
}

func isSymlink(home string, directory string) bool {
	fileInfo, err := os.Lstat(absoluteHome(home, directory))
	if nil != err {
		return false
	}

	return fileInfo.Mode()&os.ModeSymlink > 0
}

func splitPath(dirpath string) []string {
	newPath := []string{}
	for _, dir := range strings.Split(dirpath, SEPARATOR) {
		if dir != "" {
			newPath = append(newPath, dir)
		}
	}

	return newPath
}

func pathshorten(home string, dirpath string, symlinks bool) string {
	if dirpath == ROOT || len(dirpath) == 0 {
		return ROOT
	}

	basepath := splitPath(dirpath)
	last := len(basepath) - 1

	normalpath, shortpath := "", ""
	if shortname(dirpath) == ROOT {
		normalpath, shortpath = ROOT, ROOT
	}

	for index, directory := range basepath {
		normalpath = path.Join(normalpath, directory)

		if index != last {
			directory = shortname(directory)
		}

		if symlinks && isSymlink(home, normalpath) {
			directory += SYMLINK
		}

		shortpath = path.Join(shortpath, directory)
	}

	return shortpath
}
