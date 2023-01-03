package log

import (
	l "log"
)

const DefaultPrefix string = ""
const DefaultFlags int = 0

// Prefix configures std log package with a prefix
func Prefix(prefix string) {
	l.SetPrefix(prefix)
}

// Flags configures std log package with features
// like date, time, and file name
// see https://pkg.go.dev/log#SetFlags for options
func Flags(flags int) {
	l.SetFlags(flags)
}
