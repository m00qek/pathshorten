package main

import (
	"path"
	"path/filepath"
	"strings"
)

const SEPARATOR = string(filepath.Separator)
const ROOT = string(filepath.Separator)
const SYMLINK = "@"
const HOME = "~"

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

func splitPath(dirpath string) []string {
	newPath := []string{}
	for _, dir := range strings.Split(dirpath, SEPARATOR) {
		if dir != "" {
			newPath = append(newPath, dir)
		}
	}

	return newPath
}

func pathshorten(home string, dirpath string, isSymlink SymlinkPredicate) string {
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

		if isSymlink(home, normalpath) {
			directory += SYMLINK
		}

		shortpath = path.Join(shortpath, directory)
	}

	return shortpath
}
