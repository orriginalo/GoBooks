
# Book API

Простое REST API для управления книгами на Go + PostgreSQL

## 🚀 Запуск
```bash
docker-compose up --build
```

API доступно на `http://localhost:8080`

## 📚 Эндпоинты
- `GET /api/books` - список книг
- `POST /api/books` - добавить книгу
- `DELETE /api/books/:id` - удалить книгу

## Пример запроса
```bash
# Добавить книгу
curl -X POST -H "Content-Type: application/json" \
-d '{"title":"1984","author":"Orwell","year":1949}' \
http://localhost:8080/api/books
```

## ⚙️ Технологии
- Go 1.21
- Echo framework
- PostgreSQL
- Docker

## 🔧 Настройки БД
- Port - **5432**
- Username - postgres 
- DB Name - postgres
- Password - postgres

*конфигурационого файла нету, увы :(*