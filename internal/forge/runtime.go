package forge

import (
	"io"
	"os"
)

// NewRuntime is foobar
func NewRuntime() *Runtime {
	return &Runtime{
		Environment: NewEnvironment(),
		Stdout:      os.Stdout,
		Stderr:      os.Stderr,
		Stdin:       os.Stdin,
	}
}

// Runtime is foobar
type Runtime struct {
	Environment Environment
	Stdout      io.Writer
	Stderr      io.Writer
	Stdin       io.Reader
}
