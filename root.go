package main

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "git-blame-ignore",
	Short: "Manage your blame ignore revs file",
}

// defaultIgnoreRevsFile is the default path, relative to the repository root, to the
// blame ignore revs file.
const defaultIgnoreRevsFile = ".git-blame-ignore-revs"

// blameIgnoreRevsFile is the path to the blame ignore revs file.
var blameIgnoreRevsFile = ""

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.PersistentFlags().StringVarP(
		&blameIgnoreRevsFile,
		"ignore-revs-file",
		"f",
		"",
		"Path to the blame ignore revs file",
	)
}
