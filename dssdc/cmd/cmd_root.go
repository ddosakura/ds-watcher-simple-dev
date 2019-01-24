package cmd

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile string
	cfg     CFG

	// ether use log
	pLog bool

	// remoteURL in `get` command
	remoteURL      string
	remoteEntryURL string

	// publishWay
	allPublishWay bool

	rootCmd = &cobra.Command{
		Use:   "dssdc",
		Short: "A CLI for developer",
		Long: `Dssdc is a tool for developer to fresh page.
This application can call the http api to tell the manager where and when the developer is working.`,
	}
)

// Execute executes the root command.
func Execute() {
	rootCmd.Execute()
}

// Ver set the version.
func Ver(ver string) {
	rootCmd.Version = ver
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "./", "config file path")

	rootCmd.PersistentFlags().BoolVarP(&pLog, "log", "l", false, "print log message.")

	host, err := os.Hostname()
	if err != nil {
		host = "Unknow"
	}

	viper.SetDefault("project", "Unknow")
	viper.SetDefault("developer", host)
	viper.SetDefault("port", 3000)
	viper.SetDefault("localdb", "./dssdc.db")
	viper.SetDefault("monitor", Monitor{
		IncludeDirs: []string{"./src"},
		ExceptDirs:  []string{".git"},
		Types:       []string{".html", ".js", ".css"},
		UseWebPage:  true,
	})
	viper.SetDefault("command", Command{
		Exec:            []string{"echo fresh"},
		DelayMillSecond: 1,
	})
	viper.SetDefault("api", API{
		Root: "http://localhost:2000/",
	})

	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(devCmd)
	rootCmd.AddCommand(publishCmd)
	publishCmd.PersistentFlags().StringVarP(&publishMsg, "msg", "m", "", "git commit msg")
	publishCmd.PersistentFlags().BoolVarP(&allPublishWay, "all", "a", false, "use all publish way")
	rootCmd.AddCommand(packageCmd)
	rootCmd.AddCommand(getCmd)
	getCmd.PersistentFlags().StringVarP(&remoteURL, "remote-url", "u", "https://github.com", "remote url")
	getCmd.PersistentFlags().StringVarP(&remoteEntryURL, "entry", "e", "/ddosakura", "entry url")
	rootCmd.AddCommand(updateCmd)

	rootCmd.SetVersionTemplate(`{{with .Name}}{{printf "%s " .}}{{end}}{{printf "version %s\n" .Version}}`)
}

func initConfig() {
	if !pLog {
		w, e := os.OpenFile("./dssdc.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
		if e != nil {
			er(e)
		}
		log.SetOutput(w)
	}

	viper.AddConfigPath(cfgFile)
	viper.SetConfigName("dssdc")

	if err := viper.ReadInConfig(); err == nil {
		log.Println("Find config file:", viper.ConfigFileUsed())
		cfg = CFG{
			APIs:        viper.Sub("api"),
			Commands:    viper.Sub("command"),
			Monitors:    viper.Sub("monitor"),
			ProjectName: viper.GetString("project"),
			Developer:   viper.GetString("developer"),
			LocalDB:     viper.GetString("localdb"),
			Port:        viper.GetInt("port"),
		}
		// log.Println(cfg)

		_, e := os.Stat(cfg.LocalDB)
		if e == nil {
			if strings.HasSuffix(cfg.LocalDB, ".db") {
				initRepo(cfg.LocalDB)
			}
		}
	} else {
		log.Println(err)
	}
	/*
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			er(err)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".dssdc")

		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	*/
}
