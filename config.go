package main

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

// loads the configuration from the .env file
func loadConfig() (config Config, err error) {
	err = godotenv.Load(".env")

	if err != nil {
		return config, err
	}

	// Load the configuration from the environment variables
	config.RedisHost = os.Getenv("REDIS_HOST")
	config.RedisPort = os.Getenv("REDIS_PORT")
	config.RedisPass = os.Getenv("REDIS_PASS")
	config.RedisDB, _ = strconv.Atoi(os.Getenv("REDIS_DB"))
	config.QueueName = os.Getenv("QUEUE_NAME")
	config.WorkersCount, _ = strconv.Atoi(os.Getenv("WORKERS_COUNT"))
	config.BufferSize, _ = strconv.Atoi(os.Getenv("BUFFER_SIZE"))

	return config, nil
}
