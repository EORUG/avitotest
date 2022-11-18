# гайд по установке
```bash
docker-compose up --build
```
# примеры запросов
1. получение баланса пользователя
```bash
curl --location --request GET 'localhost:8080/users/1'
```
2 Метод начисления /users/{userID}
```bash
curl --location --request PUT 'localhost:8080/users/1' \
--header 'Content-Type: application/json' \
--data-raw '{
    "cash" : 10
}'
```
3 Метод резервирования средств
```bash
curl --location --request POST 'localhost:8080/purchases/Purchase/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "userid" : 1,
    "servisesid" : 2,
    "orderid" : 1,
    "cash" : 99
}'
```
4 Метод признания выручки
```bash
curl --location --request PUT 'localhost:8080/purchases/Purchase/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "userid" : 1,
    "servisesid" : 2,
    "orderid" : 1,
    "cash" : 99
}'
```
# Структура проекта
модели находятся в папке models

контроллеры лежат в папке handler

взаимодействие с базой данных в папке db

миграции базы накатываются в db/db.go



