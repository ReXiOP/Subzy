package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "subzy",
	Short: "Subdomain takeover tool Mod by Sajid",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func Execute() {
	rootCmd.Execute()
}
