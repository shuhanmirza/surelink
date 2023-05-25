package util

import "github.com/spf13/viper"

type GlobalConfig struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	RedisUrl      string `mapstructure:"REDIS_URL"`
}

type SecretConfig struct {
	LinkPreviewApiKey string `mapstructure:"LINKPREVIEW_APIKEY"`
}

func LoadGlobalConfig(path string) (config GlobalConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func LoadSecretConfig(path string) (config SecretConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("secret")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
