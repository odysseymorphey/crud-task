# Beatiful CRUD realizatioin
## Установка

```bash
git clone github.com/odysseymorphey/crud-task.git
cd crud-task
docker-compose up -d
```
## Ручки
`GET /api/users` Получить всех пользователей \
`GET /api/users/:id` Получить пользователя по ID \
`POST /api/users` Добавить пользователя \
`PUT /api/users/` Изменить данные о пользователе \
`DELETE /api/users/:id` Удалить пользовтеля по ID

## Тестирование
1. Поднять композ
2. Открыть файл `tests/api.http`
3. Запускать тесты в логическом порядке, то есть прежде чем что-то удалить/изменить/получить, нужно это что-то создать