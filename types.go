package main

type JobPayload struct {
	UserID  string `json:"user_id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

type Job struct {
	Payload JobPayload
}

type Config struct {
	QueueName    string `mapstructure:"QUEUE_NAME"`
	WorkersCount int    `mapstructure:"WORKERS_COUNT"`
	RedisHost    string `mapstructure:"REDIS_HOST"`
	RedisPort    string `mapstructure:"REDIS_PORT"`
	RedisPass    string `mapstructure:"REDIS_PASS"`
	RedisDB      int    `mapstructure:"REDIS_DB"`
}
