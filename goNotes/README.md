* CI/CD: CircleCI / GitHub Actions, docker-compose
* Debug: Opentracing (Zipkin)
* Search and filter: Elasticsearch + Logstash
* ETL: Apache Airflow
* Logs: Telegraf / FileBeat
* Service Discovering: Consul

# NoteService

Golang, REST API, JSON, MongoDB

# FileService

S3 MinIO

# CategoryService

Python 3, Flask, Neo4j (Cypher)

# TagsService



# APIService

RESTful, Swagger

# SearchService

Python 3, Elasticsearch

# UserService

REST + PostgreSQL

# WebApplication

Vue.js / React.js

Эндпоинты:

- аутентификации и регистрации
- GET /api/categories получение древовидного списка категорий   
- POST /api/categories создать новую категорию
- GET /notes?category_id=? список заметок текущей категории
