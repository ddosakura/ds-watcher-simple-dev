package cmd

import (
	"github.com/spf13/viper"
)

type Monitor struct {
	Types       []string `yaml:"types"`
	IncludeDirs []string `yaml:"includeDirs"`
	ExceptDirs  []string `yaml:"exceptDirs"`
	UseWebPage  bool     `yaml: "useWebPage"`
}

type Command struct {
	Exec            []string `yaml:"exec"`
	DelayMillSecond int      `yaml:"delayMillSecond"`
}

type API struct {
	Root     string `yaml:"root"`
	Upload   string `yaml:"upload"`
	Notifier string `yaml:"notifier"`
}

type CFG struct {
	APIs        *viper.Viper
	Commands    *viper.Viper
	Monitors    *viper.Viper
	ProjectName string
	Developer   string
	LocalDB     string
	Port        int
	Proxy       map[string]string
}
