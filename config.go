package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// loads the configuration from the .env file
func loadConfig() (config Config, err error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	viper.SetDefault("QUEUE_NAME", "job-queue")
	viper.SetDefault("WORKERS_COUNT", 1)
	viper.SetDefault("REDIS_HOST", "localhost")
	viper.SetDefault("REDIS_PORT", "6379")
	viper.SetDefault("REDIS_PASS", "")
	viper.SetDefault("REDIS_DB", 0)

	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err = viper.Unmarshal(&config)

	if err != nil {
		panic(fmt.Errorf("unable to decode the .env file: %w", err))
	}

	return config, nil
}
