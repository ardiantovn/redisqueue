package main

import (
	"github.com/ardiantovn/redisqueue" // Import the producer package
)

func main() {
	// Create instance of RedisQueue
	redisQueue := redisqueue.NewRedisQueue()

	// Consume
	redisQueue.Consume()
}
