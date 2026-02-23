// Package cmd holds the CLI base command
package cmd

import (
	"embed"
	"os"

	"github.com/spf13/cobra"
)

//go:embed assets/*
var assets embed.FS

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "apenas",
	Short: "A simple CLI to scaffold Rust, Go and Python projects",
	Long: `The CLI will set a minimal dir structure and copy a
	corresponding justfile with common recipes for the selected
	language.`,
	Run: func(cmd *cobra.Command, args []string) {
		// prints the help message by default
		cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
