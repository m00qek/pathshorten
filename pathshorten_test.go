package main

import (
	"testing"
)

func assert(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("Expecting \"%s\" but got \"%s\"", expected, actual)
	}
}

func assertArray(t *testing.T, expected []string, actual []string) {
	for index, element := range expected {
		if element != actual[index] {
			t.Errorf("Expecting \"%s\" but got \"%s\". \nMore specifically, at index \"%d\", was expecting \"%s\" but got \"%s\"",
				expected, actual, index, element, actual[index])
		}
	}
}

func TestSymbolicHomeWithAbsoluteHome(t *testing.T) {
	value := symbolicHome("/home/quincas", "/home/quincas/theories/humanitas")
	assert(t, "~/theories/humanitas", value)
}

func TestSymbolicHomeWithSymbolicHome(t *testing.T) {
	value := symbolicHome("/home/quincas", "~/theories/humanitas")
	assert(t, "~/theories/humanitas", value)
}

func TestAbsoluteHomeWithSymbolicHome(t *testing.T) {
	value := absoluteHome("/home/quincas", "/home/quincas/theories/humanitas")
	assert(t, "/home/quincas/theories/humanitas", value)
}

func TestAbsoluteHomeWithAbsoluteHome(t *testing.T) {
	value := absoluteHome("/home/quincas", "~/theories/humanitas")
	assert(t, "/home/quincas/theories/humanitas", value)
}

func TestShortNameOfCommonDirectory(t *testing.T) {
	value := shortname("theories")
	assert(t, "t", value)
}

func TestShortNameOfHiddenDirectory(t *testing.T) {
	value := shortname(".theories")
	assert(t, ".t", value)
}

func TestSplitPath(t *testing.T) {
	value := splitPath("/home/quincas/theories/humanitas")
	assertArray(t, []string{"home", "quincas", "theories", "humanitas"}, value)
}

func TestPathshortenWithoutSymlinks(t *testing.T) {
	value := pathshorten("/home/quincas",
		"/home/quincas/theories/humanitas",
		func(home string, dirpath string) bool { return false })

	assert(t, "/h/q/t/humanitas", value)
}

func TestPathshortenWithSymlinks(t *testing.T) {
	value := pathshorten("/home/quincas",
		"/home/quincas/theories/humanitas",
		func(home string, dirpath string) bool {
			return dirpath == "/home/quincas/theories"
		})

	assert(t, "/h/q/t@/humanitas", value)
}
