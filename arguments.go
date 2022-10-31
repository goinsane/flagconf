package flagconf

import (
	"bufio"
	"io"
	"os"
	"strings"
	"unicode"
)

// Arguments holds command line arguments that will be read.
type Arguments []string

// Read reads arguments from io.Reader and appends them.
func (a *Arguments) Read(r io.Reader) error {
	args := make([]string, 0, 4096)
	scnr := bufio.NewScanner(r)
	for scnr.Scan() {
		line := strings.TrimSpace(scnr.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		idx := strings.IndexFunc(line, func(r rune) bool {
			return unicode.IsSpace(r)
		})
		if idx < 0 {
			args = append(args, "-"+line+"=")
			continue
		}
		args = append(args, "-"+line[:idx]+"="+tryUnquote(strings.TrimSpace(line[idx+1:])))
	}
	if e := scnr.Err(); e != nil {
		return e
	}
	*a = append(*a, args...)
	return nil
}

// ReadString reads arguments from string and appends them.
func (a *Arguments) ReadString(s string) error {
	return a.Read(strings.NewReader(s))
}

// ReadFile reads arguments from file and appends them.
func (a *Arguments) ReadFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return a.Read(f)
}
