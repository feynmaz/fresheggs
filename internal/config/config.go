package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	AppName     string      `mapstructure:"app_name"`
	LogLevel    int         `mapstructure:"log_level"`
	Server      Server      `mapstructure:"server"`
	EthClient   EthClient   `mapstructure:"eth_client"`
	Web3Account Web3Account `mapstructure:"web3_account"`
}

type Server struct {
	ListenPort int    `mapstructure:"listen_port"`
	Env        string `mapstructure:"app_env"`
}

type EthClient struct {
	Url string `mapstructure:"url"`
}

type Web3Account struct {
	Address    string `mapstructure:"address"`
	PrivateKey string `mapstructure:"private_key"`
}

func Load(path string) (*Config, error) {
	setDefaults()

	var c Config

	if path != "" {
		_, err := os.Lstat(path)
		if err != nil {
			return nil, err
		}
		viper.SetConfigFile(path)

		err = viper.ReadInConfig()
		if err != nil {
			return nil, err
		}
	}

	err := viper.Unmarshal(&c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func setDefaults() {
	viper.SetDefault("log_level", 1)
	viper.SetDefault("server.listen_port", 8080)
	viper.SetDefault("server.env", "dev")
}
