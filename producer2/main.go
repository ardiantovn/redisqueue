package main

import (
	"github.com/ardiantovn/redisqueue" // Import the producer package
)

func main() {
	// Create instance of RedisQueue
	redisQueue := redisqueue.NewRedisQueue("worker2")

	// Produce "data" into "TaskName" with HandlerFunc
	data := map[string]string{"message": "good bye"}
	redisQueue.Produce("topic2", redisqueue.HandlePrint, data)
}
