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
