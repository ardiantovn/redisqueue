package main

import (
	"github.com/ardiantovn/redisqueue" // Import the producer package
)

func main() {
	// Create instance of RedisQueue
	redisQueue := redisqueue.NewRedisQueue()

	// Call the produce function in producer.go
	redisQueue.Consume("hello")
}
