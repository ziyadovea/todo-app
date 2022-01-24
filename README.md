# REST API на Go (Golang) для создания списков дел и управления ими.

## В приложении есть следующие возможности по следующим эндпоинтам:
- POST "/auth/sign-up" - регистрация пользователя
- POST "/auth/sign-in" - аутентификация пользователя <br>
К следующим адресам доступ имеет только авторизованный в системе пользователь:
- POST "/api//lists" - создание нового списка
- GET "/api//lists"  - получение всех списков
- DELETE "/api/lists" - удаление всех списков
- PUT "/api//lists/:list_id" - обновление списка с id, равным :list_id
- GET "/api//lists/:list_id" - получение списка с id, равным :list_id
- DELETE "/api/lists/:list_id" - удаление списка с id, равным :list_id
- POST "/api//lists/:list_id/items" - создание нового элементка списка
- GET "/api//lists/:list_id/items" - получение всех элементов списка 
- DELETE "/api//lists/:list_id/items" - удаление всех элементов списка
- PUT "/api//lists/:list_id/items/:item_id" - обновление элемента списка с id, равным :item_id
- GET "/api//lists/:list_id/:item_id" - получение элемента списка с id, равным :item_id
- DELETE "/api//lists/:list_id/:item_id" - удаление элемента списка с id, равным :item_id

## Примечания к проекту:
- Архитектура REST API
- Фреймворк [gin-gonic/gin](https://github.com/gin-gonic/gin)
- Чистая архитектура (handler -> service -> repository)
- PostgreSQL в качестве БД. Для работы с БД используется пакет [sqlx](https://github.com/jmoiron/sqlx). Генерация файлов миграций. 
- Docker, Docker-compose, Makefile 
- Конфигурация приложения с помощью пакета [viper]("https://github.com/spf13/viper"). Работа с переменными окружения.
- Регистрация и аутентификация. Работа с JWT. 
- Middleware
- Graceful Shutdown
- Частичное покрытие кода тестами с использованием пакета [testify](https://github.com/stretchr/testify), а также с использованием моков - [mock](https://github.com/golang/mock).

### Для запуска приложения используется docker-compose, вызов команд описан в makefile:

```
make build && make run
```

Если приложение запускается впервые, требуется также применить команду для миграции БД:

```
make migrate
```
