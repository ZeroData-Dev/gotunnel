package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:     "gotunnel",
	Short:   "Gotunnel Application",
	Long:    "This is a simple TCP tunnel application built with Go.",
	Example: "gotunnel --help",
	Version: "1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Welcome to Gotunnel Application!. Use --help for more information.")
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "Display the version of the application")

	// Add subcommands
	rootCmd.AddCommand(versionCmd)
}
