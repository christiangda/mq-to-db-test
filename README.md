# test-rmq

Testing RabbitMQ producer/consumer

## How to use

### Producers

Queue Producer

```bash
go run cmd/producer/producer.go --help
```

Queue Producer with Exchange

```bash
go run cmd/producer-with-exchange/producer.go --help
```

### Consumers

Queue consumer

```bash
go run cmd/consumer/consumer.go --help
```

## Message

```json
{
    "TYPE":"SQL",
    "CONTENT":{
        "SERVER":"localhost",
         "DB":"postgresql",
         "USER":"postgres",
         "PASS":"mysecretpassword",
         "SENTENCE":"SELECT pg_sleep(1.5);"
    },
    "DATE":"2020-01-01 00:00:01.000000-1",
    "APPID":"test",
    "ADITIONAL":null,
    "ACK": false,
    "RESPONSE":null
}
```
