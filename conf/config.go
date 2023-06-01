package conf

import (
	"fmt"
	"kvm-dashboard/utils"

	"github.com/spf13/viper"
)

var C *Config

type Config struct {
	AppConf      AppConf
	InfluxDBConf InfluxDBConf
	RpcConf      RpcConf
	User         User
}

type AppConf struct {
	Name    string
	Version string
	Host    string
	Port    string
	Mode    string
}

type InfluxDBConf struct {
	URL    string
	Token  string
	Org    string
	Bucket string
}

type RpcConf struct {
	SeverHost  string
	SeverPort  string
	ClientHost string
	ClientPort string
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func InitConf() *Config {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("conf")
	viper.SetConfigName("app")

	if err := viper.ReadInConfig(); err != nil {
		utils.Log.Panic(fmt.Errorf("fatal error config file: %w", err))
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	utils.Log.Info("Using config file:", viper.ConfigFileUsed())

	config := &Config{
		AppConf: AppConf{
			Name:    viper.GetString("app.name"),
			Version: viper.GetString("app.version"),
			Host:    viper.GetString("app.host"),
			Port:    viper.GetString("app.port"),
			Mode:    viper.GetString("app.mode"),
		},
		InfluxDBConf: InfluxDBConf{
			URL:    viper.GetString("database.influxdb.url"),
			Token:  viper.GetString("database.influxdb.token"),
			Org:    viper.GetString("database.influxdb.org"),
			Bucket: viper.GetString("database.influxdb.bucket"),
		},
		RpcConf: RpcConf{
			SeverHost:  viper.GetString("rpc.server.host"),
			SeverPort:  viper.GetString("rpc.server.port"),
			ClientHost: viper.GetString("rpc.client.host"),
			ClientPort: viper.GetString("rpc.client.port"),
		},
		User: User{
			Username: viper.GetString("user.username"),
			Password: viper.GetString("user.password"),
		},
	}

	return config
}
