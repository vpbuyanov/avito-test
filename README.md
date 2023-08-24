# START

## Start db and create migrations
``` shell
docker compose -f docker-compose.postgres.yml up -d 
```    
  
---
### windows
```shell
scoop install migrate 
```

### MacOS
```shell
brew install golang-migrate
```

### Linux
```shell
 curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-arm64.tar.gz | tar xvz
``` 
---  

### migrations
```shell
export POSTGRESQL_URL=postgres://postgres:postgres@0.0.0.0:5432/avitoTest?sslmode=disable
migrate -database ${POSTGRESQL_URL} -path migrations up
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