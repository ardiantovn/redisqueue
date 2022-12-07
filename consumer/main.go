package main

import (
	"github.com/ardiantovn/redisqueue" // Import the producer package
)

func main() {
	// Create instance of RedisQueue
	redisQueue := redisqueue.NewRedisQueue("api-worker")

	// Consume everything from "TaskName" with HandlerFunc
	redisQueue.Consume("hello", redisqueue.HandlePrint)
}
