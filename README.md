## Quick start

---
### Run
Для запуска приложения и его зависимостей
``` shell
docker compose -f docker-compose.run.yml up -d
```


### Debug
Для запуска зависимостей без запуска самого приложения
``` shell
docker compose -f docker-compose.debug.yml up -d
```

## Handlers

---
### JSON Request
URL `/api/user`
Method `POST`
```json
{
  "name": "your_name"
}
```
URL `/api/user/check`
Method `POST`
```json
{
  "user_id": 1
}
```
URL `/api/user/change`
Method `POST`
```json
{
  "user_id": 1,
  "add_user": [1, 2],
  "del_user": [3, 4]
}
```
URL `/api/segments`
Method `POST`
```json
{
  "slug": "AVITO_VOICE_MESSAGES"
}
```
URL `/api/segments`
Method `DELETE`
```json
{
  "slug": "AVITO_VOICE_MESSAGES"
}
```

### JSON Response
URL `/api/segments`
Method `POST`
```json
{
  "segments_id": 1,
  "slug": "AVITO_VOICE_MESSAGES",
  "status": "Success"
}
```
URL `/api/segments`
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
URL `/api/user/change`
Method `POST`
```json
{
  "user_id": 1,
  "segments": [1, 2],
  "Status": "Success"
}
```
URL `/api/user/check`
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