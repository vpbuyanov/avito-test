# START

## Start db and create migrations
``` shell
docker compose -f docker-compose.postgres.yml up -d 
```
---

### migrations
```shell
go run migrations/create-migrations.go
```
---
## Start Api
*Create a file config.yml in the project root*  
*Fill in this file as in the example config.example.yml*

# run in cmd
```shell
go run cmd/api/main.go
```
  
# run in docker
```shell
docker compose build
docker compose up
```