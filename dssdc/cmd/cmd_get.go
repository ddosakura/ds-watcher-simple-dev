package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"see"},
	Short:   "See others project.",
	Long:    `See others project.`,

	Run: func(cmd *cobra.Command, args []string) {
		wd, err := os.Getwd()
		if err != nil {
			er(err)
		}

		if len(args) == 0 {
		} else if len(args) == 1 {
			arg := args[0]
			if arg[0] == '.' {
				arg = filepath.Join(wd, arg)
			}
			if filepath.IsAbs(arg) {
			} else {
			}
			log.Println(wd, arg)
		} else {
			er("please provide only one argument")
		}
	},
}
