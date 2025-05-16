package git

import "os/exec"

type command struct{}

func (command) cmd(subcommand string, arg ...string) *exec.Cmd {
	args := make([]string, 0, len(arg)+1)
	args = append(args, subcommand)
	args = append(args, arg...)
	return exec.Command("git", args...)
}

// Run implements the Interface interface.
func (c command) Run(subcommand string, arg ...string) error {
	return c.cmd(subcommand, arg...).Run()
}

// Output implements the Interface interface.
func (c command) Output(subcommand string, arg ...string) ([]byte, error) {
	return c.cmd(subcommand, arg...).Output()
}
