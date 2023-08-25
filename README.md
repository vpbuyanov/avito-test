# START

### DB and API
``` shell
docker compose -f docker-compose.run.yml up postgres server -d
```
---

### migrations
```shell
go run migrations/create-migrations.go
```
---