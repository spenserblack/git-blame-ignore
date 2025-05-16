package git_test

import (
	"errors"
	"testing"

	"github.com/spenserblack/git-blame-ignore/pkg/git"
)

func TestRoot(t *testing.T) {
	m := mock{
		output: str("/path/to/root/.git"),
	}
	git := git.New(&m)
	root, err := git.Root()
	if err != nil {
		t.Fatalf("err = %v, want nil", err)
	}
	want := "/path/to/root/"
	if root != want {
		t.Fatalf("root = %q, want %q", root, want)
	}
}

func TestBlameIgnoreRevsFile(t *testing.T) {
	tests := []struct {
		name string
		mock mock
		want string
	}{
		{
			name: "value is set",
			mock: mock{
				output: str(".ignore-revs"),
			},
			want: ".ignore-revs",
		},
		{
			name: "error is returned",
			mock: mock{
				output: fn(func(string, ...string) ([]byte, error) {
					return nil, errors.New("Something bad happened :(")
				}),
			},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			git := git.New(&tt.mock)
			got := git.BlameIgnoreRevsFile()
			if got != tt.want {
				t.Fatalf("BlameIgnoreRevsFile() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestAsRev(t *testing.T) {
	m := mock{
		output: str("1da177e4c3f41524e886b7f1b8a0c1fc7321cac2\n"),
	}
	want := "1da177e4c3f41524e886b7f1b8a0c1fc7321cac2"
	git := git.New(&m)
	rev, _ := git.AsRev("HEAD")

	if rev != want {
		t.Fatalf("AsRev() = %q, want %q", rev, want)
	}
}

type mock struct {
	// Output is the output that should be returned.
	output response
	// Subcommand is the subcommand that was called.
	Subcommand string
	// Args are the arguments passed to the subcommand.
	Args []string
}

// Implements Interface.
func (m *mock) Run(subcommand string, arg ...string) error {
	m.Subcommand = subcommand
	m.Args = arg
	_, err := m.output.Output(subcommand, arg...)
	return err
}

// Implements Interface.
func (m *mock) Output(subcommand string, arg ...string) ([]byte, error) {
	m.Run(subcommand, arg...)
	return m.output.Output(subcommand, arg...)
}

type response interface {
	// Output gets the mock response output.
	Output(subcommand string, arg ...string) ([]byte, error)
}

// str is a simple response that always returns the same string.
type str string

// Output implements response.
func (s str) Output(string, ...string) ([]byte, error) {
	return []byte(s), nil
}

// commands is a response that maps the command to the response that should be returned.
// Arguments are ignored.
type commands map[string]string

// Output implements response.
func (c commands) Output(subcommand string, _ ...string) ([]byte, error) {
	return []byte(c[subcommand]), nil
}

// fn is a response that runs an arbitrary function to get the appropriate output.
type fn func(subcommand string, arg ...string) ([]byte, error)

// Output implements response.
func (f fn) Output(subcommand string, arg ...string) ([]byte, error) {
	return f(subcommand, arg...)
}
