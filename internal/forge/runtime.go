package forge

import (
	"errors"
	"io"
	"io/fs"
	"os"
	"os/exec"
)

// NewRuntime is foobar
func NewRuntime() *Runtime {
	return &Runtime{
		Environment: NewEnvironment(),
		FS:          os.DirFS("."),
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
	FS          fs.FS
}

// Execute is foobar
func (runtime Runtime) Execute(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stderr = runtime.Stderr
	cmd.Stdin = runtime.Stdin
	cmd.Stdout = runtime.Stdout

	return cmd.Run()
}

// ReadInDefaultEnvironmentFiles is foobar
func (runtime Runtime) ReadInDefaultEnvironmentFiles() error {
	defaultFiles := []string{
		// Values already set in the Environment will not be changed
		".env.local", // Not tracked in git, first priority
		".env",       // Defaults if not set other
	}

	for _, defaultFile := range defaultFiles {
		if err := runtime.ReadInEnvironmentFile(defaultFile); err != nil {
			return err
		}
	}

	return nil
}

// ReadInEnvironmentFile is foobar
func (runtime Runtime) ReadInEnvironmentFile(fileName string) error {
	file, err := runtime.FS.Open(fileName)
	if err != nil {
		// File does not exist errors are okay
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}

		return err
	}

	fileContentBytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	fileContents := string(fileContentBytes)

	return runtime.Environment.ImportEnvFileContents(fileContents)
}
