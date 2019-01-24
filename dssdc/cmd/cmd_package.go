package cmd

import (
	"github.com/spf13/cobra"
)

var packageCmd = &cobra.Command{
	Use:   "package",
	Short: "Package project files.",
	Long:  `Package project files.`,

	Run: func(cmd *cobra.Command, args []string) {
		pkg(cfg.ProjectName + ".tar.gz")
	},
}
