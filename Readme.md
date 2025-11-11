# QA API

Минимальный REST API для вопросов и ответов на Go + PostgreSQL.

##  Запуск

```bash
git clone https://github.com/ashurov-imomali/qa-api.git
cd qa-api

docker-compose up --build
```

После запуска:

* API доступен на `http://localhost:8080`
* PostgreSQL — на `localhost:5432`

Миграции применяются автоматически при старте контейнера.

## API

### Вопросы

* `GET /questions` — список всех вопросов
* `POST /questions` — создать вопрос (`{"text":"текст вопроса"}`)
* `GET /questions/{id}` — вопрос с ответами
* `DELETE /questions/{id}` — удалить вопрос и ответы

### Ответы

* `POST /questions/{id}/answers` — добавить ответ (`{"user_id":"uuid","text":"ответ"}`)
* `GET /answers/{id}` — получить ответ
* `DELETE /answers/{id}` — удалить ответ

## Пример запроса

```bash
# создать вопрос
curl -X POST -H "Content-Type: application/json" \
  -d '{"text":"Как запустить проект?"}' \
  http://localhost:8080/questions

# добавить ответ
curl -X POST -H "Content-Type: application/json" \
  -d '{"user_id":"123e4567-e89b-12d3-a456-426614174000","text":"Через docker-compose up"}' \
  http://localhost:8080/questions/1/answers
```

## Технологии

* Go (net/http, GORM)
* PostgreSQL
* goose (миграции, запускаются в коде)
* Docker, docker-compose

## Автор

[ashurov-imomali](https://github.com/ashurov-imomali)
