package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

var ctx = context.Background()
var config Config
var redisClient *redis.Client

func init() {
	var err error
	config, err = loadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort),
		Password: config.RedisPass,
		DB:       config.RedisDB,
	})

	// Test the redis connection
	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}
}

func main() {
	// This channel will be used to send jobs to the workers
	jobChannel := make(chan *Job)

	// Start workers
	numWorkers := config.WorkersCount
	for i := 0; i < numWorkers; i++ {
		go worker(jobChannel, i)
	}

	// Queue reader
	// This goroutine will read from the queue and send jobs to the jobChannel
	// Using a single goroutine to read from the queue and send jobs to the channel help us avoid
	// stressing the Redis server with too many BLPop requests
	go func() {
		for {
			// When queue is empty: Block at BLPop for up to 1 second → Continue loop
			result, err := redisClient.BLPop(ctx, time.Second, config.QueueName).Result()
			if err != nil {
				if errors.Is(err, redis.Nil) {
					fmt.Printf("Queue reader: No jobs in queue\n")
					continue
				}

				log.Printf("Queue reader: Error popping job from the Queue: %v", err)
				time.Sleep(time.Second)
				continue
			}

			if len(result) < 2 {
				continue
			}

			// result[0] is the queue name, result[1] is the job JSON
			jobJSON := result[1]

			var payload JobPayload
			err = json.Unmarshal([]byte(jobJSON), &payload)
			if err != nil {
				log.Printf("Queue reader: Error unmarshalling job JSON: %v", err)
				continue
			}

			jobChannel <- &Job{Payload: payload}
		}
	}()

	// Keep main goroutine running
	select {}
}

func worker(jobChannel <-chan *Job, workerID int) {
	for job := range jobChannel {
		fmt.Printf("Worker %d received job %s\n", workerID, job.Payload.UserID)
		processJob(job.Payload, workerID)
		time.Sleep(time.Second)
	}
}

func processJob(job JobPayload, workerID int) {
	fmt.Printf("Worker %d processing job %s\n", workerID, job.UserID)

	// Simulate some work
	time.Sleep(5 * time.Second)
	fmt.Printf("Worker %d finished job %s\n", workerID, job.UserID)
}
