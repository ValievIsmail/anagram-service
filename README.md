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