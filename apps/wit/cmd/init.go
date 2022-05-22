package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// InitialStash ...
var InitialStash string

// SeparateWitDir ...
var SeparateWitDir string

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [<directory>]",
	Short: "Create an empty Wit repository or reinitialize an existing one",
	Long: `This command creates an empty Wit repository - basically a .wit directory with subdirectories for
stashes, reports and the like. An initial stash without any inventory will be created (see
the --initial-stash option below for its name).

If the $WIT_DIR environment variable is set then it specifies a path to use instead of ./.wit for
the base of the repository.

Running wit init in an existing repository is safe. It will not overwrite things that are already
there. The primary reason for rerunning wit init is to pick up newly added templates (or to move
the repository to another place if --separate-wit-dir is given).`,
	Run: func(cmd *cobra.Command, args []string) {
		var path string
		if len(args) == 0 {
			path = "."
		} else if len(args) == 1 {
			path = args[0]
		}
		fmt.Println(fmt.Printf("init called, path given: %q", path))
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	initCmd.Flags().BoolP("quiet",
		"q",
		false,
		"Only print error and warning messages; all other output will be suppressed.",
	)

	initCmd.Flags().StringVar(&SeparateWitDir,
		"seperate-wit-dir",
		".wit",
		`Instead of initializing the repository as directory to either $WIT_DIR or ./.wit/,
create a text file there containing the path to the actual repository.
This file acts as filesystem-agnostic Wit symbolic link to the repository.

If this is reinitialization, the repository will be moved to
the specified path.`,
	)

	initCmd.Flags().StringVarP(&InitialStash,
		"initial-stash",
		"b",
		"main",
		`Use the specified name for the initial stash in the newly created repository.
If not specified, fall back to the default name (currently main; the name can be
customized via the init.defaultStash configuration variable).`)
}
