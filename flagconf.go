// Package flagconf provides utilities to parse the flags of GoLang's flag package from the configuration reader or string or file.

package flagconf

import (
	"flag"
	"io"
)

// Parse parses flags from reader and appends given arguments.
func Parse(fs *flag.FlagSet, r io.Reader, args []string) error {
	a := Arguments{}
	if e := a.Read(r); e != nil {
		return e
	}
	return fs.Parse(append(a, args...))
}

// ParseString parses flags from string and appends given arguments.
func ParseString(fs *flag.FlagSet, s string, args []string) error {
	a := Arguments{}
	if e := a.ReadString(s); e != nil {
		return e
	}
	return fs.Parse(append(a, args...))
}

// ParseFile parses flags from file and appends given arguments.
func ParseFile(fs *flag.FlagSet, path string, args []string) error {
	a := Arguments{}
	if e := a.ReadFile(path); e != nil {
		return e
	}
	return fs.Parse(append(a, args...))
}
