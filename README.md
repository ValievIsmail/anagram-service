### Anagram-service 
Сервис для поиска анаграм в словаре

Request:

1. Загрузить словарь:

```
curl localhost:8080/load -d '["foobar", "aabb", "baba", "boofar", "test"]'
```

POST http://localhost:8080/load
BODY
```
["foobar", "barfoo", "listen", "silent"]
```

2. Найти анаграму

```
curl 'localhost:8080/get?word=test'
```

GET http://localhost:8080/get?word=listen

Команды для сервиса:

1. make go - запустить сервис
2. make test - прогнать тесты
3. make build - собрать bin
4. make docker-build - собраз docker-image
5. make docker-run - запустить сервис в docker контейнере