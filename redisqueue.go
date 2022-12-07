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

type RedisQueue struct {
	Redis        *redis.Client
	QueueFactory taskq.Factory
	MainQueue    taskq.Queue
	TaskInstance taskq.Task
}

func NewRedisQueue() *RedisQueue {
	// Create a redis client
	redis := redis.NewClient(&redis.Options{
		Addr: ":6379",
	})

	// Create a queue factory
	queueFactory := redisq.NewFactory()

	// Create a queue
	mainQueue := queueFactory.RegisterQueue(&taskq.QueueOptions{
		Name:  "api-worker",
		Redis: redis, // go-redis client
	})

	return &RedisQueue{
		Redis:        redis,
		QueueFactory: queueFactory,
		MainQueue:    mainQueue,
	}
}

// Create handler
func (r *RedisQueue) HandlePrint(c context.Context, message string) error {
	// Process message here
	log.Println(message)
	return nil
}

func (r *RedisQueue) Produce(TaskName string, data map[string]string) {
	flag.Parse()

	ctx := context.Background()

	// Convert data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	// Register a task.
	var TaskInstance = taskq.RegisterTask(&taskq.TaskOptions{
		Name:    TaskName,
		Handler: r.HandlePrint,
	})

	// Start producing
	for {
		// Call the task with JSON data
		err := r.MainQueue.Add(TaskInstance.WithArgs(ctx, jsonData))
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println(data)
		}
		time.Sleep(time.Second)
	}
}

func (r *RedisQueue) Consume(TaskName string) {
	flag.Parse()

	c := context.Background()

	// Register a task.
	taskq.RegisterTask(&taskq.TaskOptions{
		Name:    TaskName,
		Handler: r.HandlePrint,
	})

	// Start consuming
	err := r.QueueFactory.StartConsumers(c)
	if err != nil {
		log.Fatal(err)
	}

	// Waiting for signal to stop program
	sig := r.WaitSignal()
	log.Println(sig.String())

	err = r.QueueFactory.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func (r *RedisQueue) WaitSignal() os.Signal {
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
