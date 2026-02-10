package cmd

import (
	"fmt"
	"os"

	"github.com/glitchedslimy/vito/internal/ui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vito",
	Short: "",
	Long:  ``,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.AddTemplateFunc("styleCommand", func(s string) string {
		return ui.CommandStyle.Render(s)
	})
	cobra.AddTemplateFunc("styleTitle", func(s string) string {
		return ui.TitleStyle.Render(s)
	})
	cobra.AddTemplateFunc("descStyle", func(s string) string {
		return ui.DescStyle.Render(s)
	})

	rootCmd.SetUsageTemplate(ui.UsageTemplate)
}
