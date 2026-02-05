package cmd

import (
	"github.com/glitchedslimy/vito/internal/ui"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new project from a template",
	Run: func(cmd *cobra.Command, args []string) {
		// Init Bubble Tea interface
		ui.StartTemplateSelector()
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
