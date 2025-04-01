# Go Redis Queue Consumer for Laravel

This Go application demonstrates how to consume Redis queue jobs pushed by a Laravel application. It connects to Redis and processes jobs that were enqueued using Laravel's `QueueManager::pushRaw` method.

## How Jobs are pushed from Laravel

In the Laravel application, i'm dispatching jobs using the following pattern:

```php
use Illuminate\Queue\QueueManager;
use Illuminate\Support\Str;

$queueManager = app(QueueManager::class);

$payload = [
    'user_id' => Str::uuid()->toString(),
    'name' => fake()->name(),
    'age' => fake()->randomNumber(),
    'address' => fake()->address(),
];

// Using pushRaw so the job is not serialized but sent as raw JSON
// So in the Go consumer we can decode it directly without needing to unserialize it
$queueManager->pushRaw(json_encode($payload, JSON_THROW_ON_ERROR), 'my_queue_name');
```

## Configuration

The application is configured via environment variables in a `.env` file:

```dotenv
QUEUE_NAME=erp_database_queues:my_queue_name
WORKERS_COUNT=5
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0
```

Default values are set if no `.env` file is found.

## How It Works

The consumer:

1. Connects to the configured Redis instance
2. Starts multiple worker goroutines (configurable via `WORKERS_COUNT`)
3. Uses a single reader goroutine to pull jobs from Redis using `BLPOP`
4. Distributes jobs to workers through a channel
5. Processes jobs in parallel

## Benefits

- Significantly faster job processing compared to PHP
- Efficient resource utilization with concurrent workers
- Compatible with existing Laravel queue infrastructure
- Simple to deploy and maintain