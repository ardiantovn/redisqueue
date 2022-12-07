package main

import (
	"github.com/ardiantovn/redisqueue" // Import the producer package
)

func main() {
	// Create instance of RedisQueue
	redisQueue := redisqueue.NewRedisQueue("worker1")

	// Produce "data" into "TaskName" with HandlerFunc
	data := map[string]string{"message": "hello"}
	redisQueue.Produce("hello", redisqueue.HandlePrint, data)
}
