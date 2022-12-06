# REST microservices examples

REST API:

```
POST   /task/              :  создаёт задачу и возвращает её ID
GET    /task/<taskid>      :  возвращает одну задачу по её ID
GET    /task/              :  возвращает все задачи
DELETE /task/<taskid>      :  удаляет задачу по ID
GET    /tag/<tagname>      :  возвращает список задач с заданным тегом
GET    /due/<yy>/<mm>/<dd> :  возвращает список задач, запланированных на указанную дату
```

1. Simple, no-dep
2. Gorilla/mux
3. Gin

### Запуск

`SERVERPORT=1234 go run .`

### Тест

```
SERVERADDR=localhost:1234

curl -iL -w "\n" -X POST -H "Content-Type: application/json" --data '{"text":"task first","tags":["todo", "life"], "due":"2016-01-02T15:04:05+00:00"}' ${SERVERADDR}/task/
curl -iL -w "\n" -X POST -H "Content-Type: application/json" --data '{"text":"buy milk","tags":["todo"], "due":"2016-01-03T15:04:05+00:00"}' ${SERVERADDR}/task/

curl -iL -w "\n" ${SERVERADDR}/task/1
curl -iL -w "\n" ${SERVERADDR}/task
curl -iL -w "\n" -X DELETE ${SERVERADDR}/task/2
curl -iL -w "\n" ${SERVERADDR}/tag/todo/
curl -iL -w "\n" ${SERVERADDR}/due/2016/01/03
curl -iL -w "\n" -X DELETE ${SERVERADDR}/task/
```

![simple](https://github.com/cr00z/goRest/blob/main/images/simple.png?raw=true)
