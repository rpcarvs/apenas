package cmd

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	libFlag     bool
	packageFlag bool
)

func copyJustfile(fs embed.FS, srcPath string, outPath string) error {
	data, err := fs.ReadFile(srcPath)
	if err != nil {
		return fmt.Errorf("error reading embedded justfile")
	}

	if err := os.WriteFile(outPath, data, 0o644); err != nil {
		return fmt.Errorf("error writing justfile")
	}
	return nil
}

// pythonCmd represents the python command
var pythonCmd = &cobra.Command{
	Use:   "python [project_name]",
	Short: "Scaffold the python project using uv",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		projectName := args[0]
		// default
		style := "app"
		switch {
		case libFlag:
			style = "lib"
		case packageFlag:
			style = "package"
		}

		fmt.Println("Preparing python project...")

		execCmd := exec.Command("uv",
			"init",
			projectName,
			"--"+style,
		)

		execCmd.Stdout = os.Stdout
		execCmd.Stderr = os.Stderr

		// executing the _cmd
		if err := execCmd.Run(); err != nil {
			return fmt.Errorf("failed to run uv init for %q: %w", projectName, err)
		}

		// copying justfile
		src := "assets/justfile_python"
		out := filepath.Join(projectName, "justfile")

		if err := copyJustfile(assets, src, out); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(pythonCmd)

	pythonCmd.Flags().BoolVar(&libFlag, "lib", false, "Create a lib")
	pythonCmd.Flags().BoolVar(&packageFlag, "package", false, "Create a package")
	// this does nothing. Adding just for completness
	pythonCmd.Flags().Bool("app", false, "Create an app")
}
