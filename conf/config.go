package conf

import (
	"fmt"
	"kvm-dashboard/utils"

	"github.com/spf13/viper"
)

type Config struct {
	AppConf AppConf
}

type AppConf struct {
	Name    string
	Version string
	Host    string
	Port    string
	Mode    string
}

func InitConf() *Config {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("conf")
	viper.SetConfigName("app")

	if err := viper.ReadInConfig(); err != nil {
		utils.LogWithPanic(fmt.Errorf("fatal error config file: %w", err))
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	utils.LogWithInfo("Using config file:", viper.ConfigFileUsed())

	config := &Config{
		AppConf: AppConf{
			Name:    viper.GetString("app.name"),
			Version: viper.GetString("app.version"),
			Host:    viper.GetString("app.host"),
			Port:    viper.GetString("app.port"),
			Mode:    viper.GetString("app.mode"),
		},
	}

	return config
}
