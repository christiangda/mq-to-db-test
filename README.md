# test-rmq

Testing RabbitMQ producer/consumer

## How to use

### Producers

Queue Producer

```bash
go run cmd/producer/producer.go --help
```

Queue Producer with Exchange to send a defined number of messages

```bash
go run cmd/producer-with-exchange/producer.go --help
```

Queue Producer with Exchange to send a defined number of messages per seconds

```bash
go run cmd/producer-with-rate/producer.go --help
```

```bash
go run cmd/producer-with-rate/producer.go -timeRandom -messageRate 10
```

```bash
go run ./cmd/producer-with-rate/producer.go -messageRate 1 -timeRandom --debug
```

```bash
go run ./cmd/producer-with-rate/producer.go -messageRate 1 -timeRandom -timeRandomMin 1 -timeRandomMax 11
```

### Consumers

Queue consumer

```bash
go run cmd/consumer/consumer.go --help
```

## Default Message sent

```json
{
    "TYPE":"SQL",
    "CONTENT":{
        "SERVER":"",
         "DB":"",
         "USER":"",
         "PASS":"",
         "SENTENCE":"SELECT pg_sleep(1.0);"
    },
    "DATE":"2020-01-01 00:00:01.000000-1",
    "APPID":"test",
    "ADITIONAL":null,
    "ACK": false,
    "RESPONSE":null
}
```
