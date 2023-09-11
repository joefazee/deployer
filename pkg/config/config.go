package config

import "github.com/spf13/viper"

const (

	// configFileName is the name of the config file in the root directory
	configFileName = ".env"
)

type Config struct {
	DBDriver         string `mapstructure:"DB_DRIVER"`
	DBSource         string `mapstructure:"DB_SOURCE"`
	HttServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`
	Environment      string `mapstructure:"ENVIRONMENT"`

	// redis config
	RedisAddress  string `mapstructure:"REDIS_ADDRESS"`
	RedisUsername string `mapstructure:"REDIS_USERNAME"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
	RedisDB       int    `mapstructure:"REDIS_DB"`
}

// Load the configurations from configFileName and bind them to Config
func Load(p string) (Config, error) {
	return loader(p, configFileName)
}

func loader(p string, env string) (cfg Config, err error) {

	viper.AddConfigPath(p)
	viper.SetConfigName(env)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&cfg)
	return
}
