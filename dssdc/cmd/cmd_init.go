package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a configure file for project",
	Long:  `Initialize (dssdc init) will create a example for configure file.`,

	Run: func(cmd *cobra.Command, args []string) {
		arg := getConfigPath()
		log.Println("create", arg)
		viper.WriteConfigAs(arg)
		db := viper.GetString("localdb")
		if db != "" {
			db = getDBPath(db)
			log.Println("init", db)
			initRepo(db)
		}
	},
}
