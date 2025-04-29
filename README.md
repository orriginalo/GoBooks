# Book API

Простое REST API для управления книгами на Go + PostgreSQL с примером фронтенда

## 🚀 Запуск
```bash
docker-compose up --build
```

API доступно на `http://localhost:8080`  
Фронтенд доступен на `http://localhost:8080/index.html`

## 📚 Эндпоинты
- `GET /api/books` - получить список книг
- `POST /api/books` - добавить новую книгу
- `DELETE /api/books/:id` - удалить книгу по ID

## 🖥 Фронтенд
Включён простой HTML-интерфейс (`index.html`), который:
- Отображает список книг
- Позволяет добавлять новые книги
- Даёт возможность удалять существующие

## Пример запроса
```bash
# Добавить книгу
curl -X POST -H "Content-Type: application/json" \
-d '{"title":"1984","author":"Orwell","year":1949}' \
http://localhost:8080/api/books
```

## ⚙️ Технологии
- Backend: Go 1.21 + Echo framework
- Database: PostgreSQL
- Frontend: Vanilla JS + HTML
- Контейнеризация: Docker

## 🔧 Настройки БД
- **Порт:** 5432
- **Пользователь:** postgres
- **База данных:** postgres
- **Пароль:** postgres

*Настройки задаются напрямую в коде (main.go) и docker-compose.yml*