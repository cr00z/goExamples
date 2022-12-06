# goTodoList

* Router: gin
* DB: jmoiron/sqlx, postgres, migrations, trnsactions
* Configs: viper, godotenv
* Logs: logrus
* Auth: JWT
* Documentation: Swagger
* Tests: gomock

## Run

```
docker pull postgress
docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d --rm postgres
```

### Migrate

```
brew install golang-migrate
migrate create -ext sql -dir ./schema -seq init
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up
docker exec -it todo-db /bin/bash

root@75143aee9109:/# psql -U postgres
postgres=# \d

                 List of relations
 Schema |        Name        |   Type   |  Owner
--------+--------------------+----------+----------
 public | lists_items        | table    | postgres
 public | lists_items_id_seq | sequence | postgres
 public | schema_migrations  | table    | postgres
 public | todo_items         | table    | postgres
 public | todo_items_id_seq  | sequence | postgres
 public | todo_lists         | table    | postgres
 public | todo_lists_id_seq  | sequence | postgres
 public | users              | table    | postgres
 public | users_id_seq       | sequence | postgres
 public | users_lists        | table    | postgres
 public | users_lists_id_seq | sequence | postgres
(11 rows)

migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' down
postgres=# select * from schema_migrations;
postgres=# update schema_migrations set version = '000001', dirty=false;
```

## Swagger

```
export PATH=$PATH:$HOME/go/bin
go get -u github.com/swaggo/swag/cmd/swag
swag init -g cmd/main.go

```

## Tests

В файле service/service.go дописываем:

`//go:generate mockgen -source=service.go -destination=mocks/mock.go`

```
export PATH=$PATH:$HOME/go/bin
go install github.com/golang/mock/mockgen@v1.6.0
go generate
```
