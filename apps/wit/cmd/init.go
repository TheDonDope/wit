package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [path]",
	Short: "Initializes a new wit repository at the given path",
	Long: `Creates the necessary folder and file structure to store your repository on disk.

A folder called '.wit' will be created, which host all the files.`,
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
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
