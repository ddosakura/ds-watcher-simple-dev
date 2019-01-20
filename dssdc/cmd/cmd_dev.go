package cmd

import (
	"github.com/spf13/cobra"
)

var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "Dev mode.",
	Long:  `Watch the change of files.`,

	Run: func(cmd *cobra.Command, args []string) {
		if cfg.Port > 0 {
			go initFreshing()
		}
		initWatcher()
	},
}
