package main

import (
	"github.com/ardiantovn/redisqueue" // Import the producer package
)

func main() {
	// Create instance of RedisQueue
	redisQueue1 := redisqueue.NewRedisQueue("worker1")

	// Consume everything from "TaskName" with HandlerFunc
	redisQueue1.Consume("topic1", redisqueue.HandlePrint)
}
