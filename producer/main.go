package main

import (
	"github.com/ardiantovn/redisqueue" // Import the producer package
)

func main() {
	// Call the produce function in producer.go
	data := map[string]string{"message": "hello"}
	redisqueue.Produce(data)
}
