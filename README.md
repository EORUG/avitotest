# гайд по установке
один из подов может не подняться при первом запуске, тогда следует потушить все командой docker-compose down и запустить compose up еще раз
```bash
docker-compose up --build
```
# примеры запросов
1. получение сегментов пользователя
```bash
curl --location 'localhost:8080/GetUserInfo' \
--header 'content-type: application/json' \
--data '{"id":1}'
```
2 Метод создания сегмента
```bash
curl --location 'localhost:8080/CreateSegment' \
--header 'content-type: application/json' \
--data '{"name":"name3",
"persent":1}'
```
3 Метод добавления и удаления сегментов у пользователя
```bash
curl --location 'localhost:8080/ChangeSegment' \
--header 'content-type: application/json' \
--data '{"toAdd":["name2"],
"toDelete":[],
"userId":1,
"TTL":"2023-08-31 22:00:00"}'
```
4 Метод удаления сегментов
```bash
curl --location 'localhost:8080/DeleteSegment' \
--header 'content-type: application/json' \
--data '{"name":"name3"}'
```
# Структура проекта
модели находятся в папке models

контроллеры лежат в папке handler

взаимодействие с базой данных в папке db

# Вопросы по ходу разработки
1 Было не очень понятно как чистить по ttl так что я написал клинер в отдельной горутине 

2
