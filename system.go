package main

import (
	"fmt"
	"os"
	"os/user"
)

type SymlinkPredicate = func(string, string) bool

func homedir() (string, error) {
	user, err := user.Current()
	if nil != err {
		return "", err
	}

	return user.HomeDir, nil
}

func isSymlink(home string, directory string) bool {
	fileInfo, err := os.Lstat(absoluteHome(home, directory))
	if nil != err {
		return false
	}

	return fileInfo.Mode()&os.ModeSymlink > 0
}

func couldNotGetHomedirError() {
	message := "ERROR: Could not get the current user and its $HOME directory."
	fmt.Fprintln(os.Stderr, message)
	os.Exit(13)
}
