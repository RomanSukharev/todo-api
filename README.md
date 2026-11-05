# todo-api

Учебный REST API для Todo: слои как в Effective Mobile (config → model → migrations → repository → service → handler), PostgreSQL, позже Swagger.

## Последовательность файлов (как пишем и зачем)

Порядок важен: каждый следующий слой опирается на предыдущий.

| Шаг | Файл / папка | Статус | Зачем |
|-----|----------------|--------|--------|
| 1 | `go.mod` | ✅ | Имя модуля, зависимости |
| 2 | `cmd/todo-api/main.go` | ✅ | Точка входа: собрать и запустить |
| 3 | `internal/config/config.go` | ✅ | Env → настройки (порт, `DATABASE_URL`) |
| 4 | `.env` / `.env.example` | ✅ | Локальные значения конфига |
| 5 | `internal/model/todo.go` | ✅ | Структура Todo (общий язык слоёв) |
| 6 | `docker-compose.yml` | ✅ | PostgreSQL для разработки |
| 7 | `migrations/*.up.sql` / `*.down.sql` | ✅ | Схема БД (таблица `todos`), формат migrate |
| 8 | `internal/repository/` (+ dao, queries, pkg/database) | ✅ | SQL CRUD в стиле Effective Mobile |
| 9 | `internal/service/` | ⏳ | Бизнес-правила (валидация, ErrNotFound) |
| 10 | `internal/handler/` | ⏳ | HTTP JSON, статус-коды |
| 11 | роутер в `main` или `internal/server/` | ⏳ | Маршруты `/api/v1/todos` |
| 12 | Swagger (`swag` + `docs/`) | ⏳ | Документация API |
| 13 | этот `README.md` | ✅ | Как поднять и в каком порядке шли |

```text
main
  → config
  → db
  → repository  → PostgreSQL (схема из migrations)
  → service
  → handler
  → HTTP Listen
```

## Структура проекта

```text
todo-api/
├── cmd/todo-api/main.go
├── internal/
│   ├── config/
│   ├── model/
│   ├── errs/
│   ├── repository/
│   │   ├── repository.go
│   │   ├── todo/
│   │   ├── dao/
│   │   └── queries/
│   ├── service/       # дальше
│   └── handler/       # дальше
├── pkg/database/postgres/
├── migrations/
│   ├── 000001_create_todos.up.sql
│   └── 000001_create_todos.down.sql
├── docker-compose.yml
├── .env.example
├── go.mod
└── README.md
```

## Как запустить сейчас

### 1. Postgres

```bash
docker compose up -d
```

БД: `todo` / пользователь `todo` / пароль `todo`  
Порт на хосте: **5433** (внутри контейнера 5432).

### 2. Env

Скопируй пример при необходимости:

```bash
cp .env.example .env
```

`DATABASE_URL` должен смотреть на `localhost:5433`.

### 3. Миграции

Инструмент: [golang-migrate](https://github.com/golang-migrate/migrate) (формат `.up.sql` / `.down.sql`).

```bash
# установка CLI (один раз)
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# накат
migrate -path migrations -database "postgres://todo:todo@localhost:5433/todo?sslmode=disable" up

# откат на 1 шаг
migrate -path migrations -database "postgres://todo:todo@localhost:5433/todo?sslmode=disable" down 1
```
### 4. Приложение

```bash
go run ./cmd/todo-api
```

Сейчас печатает хост/порт из конфига. После repository/service/handler здесь поднимется HTTP-сервер.

## Эндпоинты (план)

| Метод | Путь | Описание |
|-------|------|----------|
| `POST` | `/api/v1/todos` | Создать |
| `GET` | `/api/v1/todos` | Список |
| `GET` | `/api/v1/todos/{id}` | По id |
| `PUT` | `/api/v1/todos/{id}` | Обновить |
| `DELETE` | `/api/v1/todos/{id}` | Удалить |

Swagger: `/swagger/index.html` (когда подключим).

## Что дальше

Следующий слой: **`internal/service`** (usecase) — правила поверх repository.
