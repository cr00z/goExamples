# goRabbitMQ

RabbitMQ + Fiber

## Build

```
docker-compose up -d --build
docker-compose down
```

## Test

Sender: `curl --url 'http://localhost:3000/send?msg=hello'`

RabbitMQ: `http://localhost:15672/ - guest:guest`

Consumer: `docker logs 6127714afea9`