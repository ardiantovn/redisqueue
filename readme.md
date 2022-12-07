This repo contains an example to use [vmihailenco/taskq/v3](github.com/vmihailenco/taskq/v3).

# QuickStart
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

# Note
1. If you create 2 producers with same workers name and different topic name,

example:

```
-> on `producer1/main.go`
workerName= "worker1"
topicName= "topic1"

-> on `producer2/main.go`
workerName= "worker1"
topicName= "topic2"

-> on `consumer1/main.go`
workerName= "worker1"
topicName= "topic1"

-> on `consumer2/main.go`
workerName= "worker1"
topicName= "topic2"
```

then you will have to enable two consumers at the same time. If you only enable
a consumer (suppose that you only enable `consumer1/main.go`), then you will get this error message:

```
taskq: 2022/12/07 11:45:00 consumer.go:616: task="topic2" failed (will retry=1 in dur=30s): taskq: unknown task="topic2"
```

2. If you create a producer and 2 consumers

example:

```
-> on `producer1/main.go`
workerName= "worker1"
topicName= "topic1"

-> on `consumer1/main.go`
workerName= "worker1"
topicName= "topic1"

-> on `consumer2/main.go`
workerName= "worker1"
topicName= "topic1"
```

and you enable `consumer1/main.go` and `consumer2/main.go` respectively. The message will be only
consumed by `consumer1/main.go`. If you disable `consumer1/main.go`, then the message will be
consumed by `consumer2/main.go`.

3. If you create two producers with different worker names and same topic name

example:

```
-> on `producer1/main.go`
workerName= "worker1"
topicName= "topic1"

-> on `producer2/main.go`
workerName= "worker2"
topicName= "topic1"

-> on `consumer1/main.go`
workerName= "worker1"
topicName= "topic1"

-> on `consumer2/main.go`
workerName= "worker2"
topicName= "topic1"
```

then you will get the following output

```
-> on `consumer1/main.go`
{"message":"hello"}

-> on `consumer2/main.go`
{"message":"good bye"}
```




