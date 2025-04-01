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
	QueueName    string
	WorkersCount int
	RedisHost    string
	RedisPort    string
	RedisPass    string
	RedisDB      int
	BufferSize   int
}
