package cmd

import "github.com/spf13/cobra"

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Display the version of the application",
	Long:    "This command displays the current version of the application.",
	Example: "gotunnel version",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("Gotunnel version: %s\n", "1.0.0") // Replace with actual version variable if available
	},
}
