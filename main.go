// main.go
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

type JobPayload struct {
	UserId  string `json:"user_id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Laravel creates queues with the prefix env.APP_NAME + "_queues:" + queue name
	// For example, if APP_NAME is "erp_database" and queue name is "my_queue_name",
	// the queue name will be "erp_database_queues:my_queue_name"
	queueName := "erp_database_queues:my_queue_name"

	ctx := context.Background()

	numWorkers := 50
	for i := 0; i < numWorkers; i++ {
		go worker(ctx, client, queueName, i)
	}

	// Keep main goroutine running
	select {}
}

func worker(ctx context.Context, client *redis.Client, queueName string, workerID int) {
	for {
		// When queue is empty: Block at BLPop for up to 1 second → Continue loop
		result, err := client.BLPop(ctx, time.Second, queueName).Result()
		if err != nil {
			if errors.Is(err, redis.Nil) {
				fmt.Printf("Worker %d: No jobs in queue\n", workerID)
				continue
			}

			log.Printf("Worker %d: Error popping from queue: %v", workerID, err)
			time.Sleep(time.Second)
			continue
		}

		if len(result) < 2 {
			continue
		}

		// result[0] is the queue name, result[1] is the job JSON
		jobJSON := result[1]

		fmt.Println(jobJSON)

		// Parse the job
		var job JobPayload
		err = json.Unmarshal([]byte(jobJSON), &job)
		if err != nil {
			log.Printf("Worker %d: Error unmarshalling job JSON: %v", workerID, err)
			continue
		}

		// Log the job details
		fmt.Printf("Worker %d processing job: %+v\n", workerID, job)

		// process the job
		processJob(workerID, job)

		// are available: Process job → Wait 500ms → Get next job
		time.Sleep(time.Second)
	}
}

func processJob(workerID int, job JobPayload) {
	fmt.Printf("Worker %d processing job %s\n", workerID, job.UserId)

	// Simulate some work
	time.Sleep(5 * time.Second)
	fmt.Printf("Worker %d finished job %s\n", workerID, job.UserId)
}
