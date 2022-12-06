package main

import (
	"github.com/ardiantovn/redisqueue" // Import the producer package
)

func main() {
	// Call the produce function in producer.go
	redisqueue.Produce()
}
