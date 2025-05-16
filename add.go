package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spenserblack/git-blame-ignore/pkg/git"
	"github.com/spf13/cobra"
)

const defaultRef = "HEAD"

var addCmd = &cobra.Command{
	Use:       "add [REF]",
	Short:     "Add a revision to the blame ignore revs file",
	Args:      cobra.MaximumNArgs(1),
	ValidArgs: []cobra.Completion{defaultRef},
	RunE: func(cmd *cobra.Command, args []string) error {
		git := git.Default()

		ref := defaultRef
		if len(args) > 0 {
			ref = args[0]
		}

		path := blameIgnoreRevsFile
		if path == "" {
			path = git.BlameIgnoreRevsFile()
		}
		if path == "" {
			root, err := git.Root()
			if err != nil {
				return err
			}
			path = filepath.Join(root, defaultIgnoreRevsFile)
		}
		rev, err := git.AsRev(ref)
		if err != nil {
			return err
		}

		f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer f.Close()
		// TODO Detect line ending that's used
		if _, err := f.WriteString(rev + "\n"); err != nil {
			return err
		}
		fmt.Fprintf(cmd.OutOrStdout(), "Added %s to %s\n", rev, path)

		return nil
	},
}
