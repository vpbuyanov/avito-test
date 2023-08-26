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

## Handlers
### JSON Request
URL `/api/user`
Method `POST`
```json
{
  "name": "your_name"
}
```
URL `/user/check`
Method `POST`
```json
{
  "user_id": 1
}
```
URL `/user/change`
Method `POST`
```json
{
  "user_id": 1,
  "add_user": [1, 2],
  "del_user": [3, 4]
}
```
URL `/segments`
Method `POST`
```json
{
  "slug": "AVITO_VOICE_MESSAGES"
}
```
URL `/segments`
Method `DELETE`
```json
{
  "slug": "AVITO_VOICE_MESSAGES"
}
```

### JSON Response
URL `/segments`
Method `POST`
```json
{
  "segments_id": 1,
  "slug": "AVITO_VOICE_MESSAGES",
  "status": "Success"
}
```
URL `/segments`
Method `DELETE`
```json
{
  "segments_id": 1,
  "slug": "AVITO_VOICE_MESSAGES",
  "status": "Success"
}
```
URL `/api/user`
Method `POST`
```json
{
  "user_id": 1,
  "status": "Success"
}
```
URL `/user/change`
Method `POST`
```json
{
  "user_id": 1,
  "segments": [1, 2],
  "Status": "Success"
}
```
URL `/user/check`
Method `POST`
```json
{
  "segments": [
    "AVITO_VOICE_1", 
    "AVITO_VOICE_Seva"
  ],
  "Status": "Success"
}
```