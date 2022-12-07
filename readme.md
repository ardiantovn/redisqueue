This repo contains an example to use [vmihailenco/taskq/v3](github.com/vmihailenco/taskq/v3).
## run redis on port 6379

## run producer1

```
go run producer1/main.go
```

## run consumer1

```
go run consumer1/main.go
```

## run producer2

```
go run producer2/main.go
```

## run consumer2

```
go run consumer2/main.go
```

## Note
If you create 2 producers with same workers name and different topic name,

example:

```
-> producer1/main.go
workerName= "worker1"
topicName= "topic1"

-> producer2/main.go
workerName= "worker1"
topicName= "topic2"
```

then you have to enable two consumers at the same time. If you only enable
a consumer, then you will get this error message:

```
taskq: 2022/12/07 11:45:00 consumer.go:616: task="topic2" failed (will retry=1 in dur=30s): taskq: unknown task="topic2"
```



