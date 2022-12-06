package redisqueue

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/vmihailenco/taskq/v3"
	"github.com/vmihailenco/taskq/v3/redisq"
)

var Redis = redis.NewClient(&redis.Options{
	Addr: ":6379",
})

// Create a queue factory.
var QueueFactory = redisq.NewFactory()

// Create a queue.
var MainQueue = QueueFactory.RegisterQueue(&taskq.QueueOptions{
	Name:  "api-worker",
	Redis: Redis, // go-redis client
})

// Register a task.
var HelloTask = taskq.RegisterTask(&taskq.TaskOptions{
	Name: "hello",
	Handler: func(c context.Context, message string) error {
		// Proses pesan di sini
		log.Println(message)
		return nil
	},
})

func Produce() {
	flag.Parse()

	ctx := context.Background()

	// Mengubah data ke dalam bentuk JSON
	data := map[string]string{"message": "hello"}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	// Start producing
	for {
		// Call the task with JSON data
		err := MainQueue.Add(HelloTask.WithArgs(ctx, jsonData))
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second)
	}
}

func Consume() {
	flag.Parse()

	c := context.Background()

	// Memulai konsumer untuk antrian
	err := QueueFactory.StartConsumers(c)
	if err != nil {
		log.Fatal(err)
	}

	// Menunggu sinyal untuk menghentikan program
	sig := WaitSignal()
	log.Println(sig.String())

	err = QueueFactory.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func WaitSignal() os.Signal {
	ch := make(chan os.Signal, 2)
	signal.Notify(
		ch,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	for {
		sig := <-ch
		switch sig {
		case syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM:
			return sig
		}
	}
}
