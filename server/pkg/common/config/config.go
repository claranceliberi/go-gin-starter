package config

import "github.com/spf13/viper"

type Config struct {
    Port  string `mapstructure:"PORT"`
	DBHost string `mapstructure:"DB_HOST"`
	DBPort string `mapstructure:"DB_PORT"`
	DBUser string `mapstructure:"DB_USER"`
	DBPass string `mapstructure:"DB_PASSWORD"`
	DBName string `mapstructure:"DB_NAME"`
}

// LoadConfig reads config from file or environment variables.
func LoadConfig() (c Config, err error) {
    viper.AddConfigPath("./pkg/common/config/envs")
    viper.SetConfigName("dev")
    viper.SetConfigType("env")

    viper.AutomaticEnv()

	// If a config file is found, read it in.
    err = viper.ReadInConfig()

    if err != nil {
        return
    }

	// Unmarshal the config into the c variable.
    err = viper.Unmarshal(&c)

    return
}