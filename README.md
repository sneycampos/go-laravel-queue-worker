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

$queueManager->pushRaw(json_encode($payload, JSON_THROW_ON_ERROR), 'my_queue_name');
```

## How the Go Consumer Works

The Go consumer:

1. Connects to your Redis instance
2. Uses multiple workers (50, check the code) to process jobs concurrently
3. Processes JSON payloads with the structure defined in the `JobPayload` struct

## Benefits

- Significantly faster job processing compared to PHP
- Efficient resource utilization with concurrent workers
- Compatible with existing Laravel queue infrastructure
- Simple to deploy and maintain