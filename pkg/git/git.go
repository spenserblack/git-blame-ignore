// Package git provides the interface for running Git commands.
package git

import (
	"path"
	"strings"
)

// Interface is the main git interface. It can be a Git command, or a mock for testing.
type Interface interface {
	// Run runs a command.
	Run(subcommand string, arg ...string) error
	// Output runs the command and returns its standard output.
	Output(subcommand string, arg ...string) ([]byte, error)
}

// Git is the main git command.
type Git struct {
	// Cmd is the command interface.
	cmd Interface
}

// New returns a new git interface.
func New(i Interface) Git {
	return Git{
		cmd: i,
	}
}

// Default returns the default git interface.
func Default() Git {
	return Git{
		cmd: command{},
	}
}

// Root gets the root git directory based on the current working directory.
func (git Git) Root() (string, error) {
	output, err := git.cmd.Output("rev-parse", "--git-dir")
	if err != nil {
		return "", err
	}
	root, _ := path.Split(string(output))
	return root, nil
}

// GetConfig gets a git configuration value. It will resolve local, global, and
// system configuration.
func (git Git) GetConfig(keys ...string) (string, error) {
	key := configKey(keys)
	output, err := git.cmd.Output("config", "--get", key)
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// BlameIgnoreRevsFile gets the configured file used to ignore revs in blame view.
//
// Returns an empty string if it is not set, or if there is any other reason that it
// can't be fetched. Use GetConfig for more control.
func (git Git) BlameIgnoreRevsFile() string {
	path, _ := git.GetConfig("blame", "ignoreRevsFile")
	return path
}

// AsRev converts a reference to a rev.
func (git Git) AsRev(ref string) (string, error) {
	output, err := git.cmd.Output("rev-parse", ref)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

// configKey converts a collection of nested keys into a single configuration key.
func configKey(keys []string) string {
	return strings.Join(keys, ".")
}
