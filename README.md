# Copy Paste Service

Простой сервис для быстрого обмена текстом через короткие, человеко-читаемые ссылки.
Заметки автоматически удаляются через 3 часа.

## 🚀 Быстрый старт

```bash
docker compose up --build
```

После запуска:
- **Frontend**: http://localhost:3000
- **Backend API**: http://localhost:8080
- **Swagger UI**: http://localhost:8080/swagger/

## 📁 Структура проекта

```
.
├── backend/                 # Go backend
│   ├── cmd/server/          # Точка входа
│   ├── internal/            # Внутренние пакеты
│   ├── migrations/          # SQL миграции
│   ├── api/                 # OpenAPI спецификация
│   ├── Dockerfile           # Backend образ
│   └── Dockerfile.migrations # Миграции образ
├── frontend/                # React frontend
│   ├── src/
│   │   ├── pages/           # Страницы
│   │   ├── api.ts           # API клиент
│   │   └── config.ts        # Конфигурация
│   ├── Dockerfile           # Frontend образ
│   └── nginx.conf           # Nginx конфиг
└── docker-compose.yml       # Оркестрация
```

## ⚙️ Конфигурация домена

### Для Production

В `docker-compose.yml` измените следующие переменные:

```yaml
backend:
  environment:
    # URL для ссылок в API ответах
    BASE_URL: https://paste.example.com

frontend:
  build:
    args:
      # URL бэкенда
      VITE_API_URL: https://api.paste.example.com
      # Домен фронтенда для ссылок
      VITE_APP_URL: https://paste.example.com
```

### Для локальной разработки

Frontend (Vite dev server):
```bash
cd frontend
npm install
npm run dev
```

Создайте файл `frontend/.env.local`:
```
VITE_API_URL=http://localhost:8080
VITE_APP_URL=http://localhost:5173
```

## 🔧 Makefile команды

```bash
cd backend
make run             # Запустить backend локально
make build           # Собрать бинарник
make sqlc            # Сгенерировать код sqlc
make migrate-up      # Применить миграции
```

## 📡 API

| Метод | Путь | Описание |
|-------|------|----------|
| `POST` | `/api/notes` | Создать заметку |
| `GET` | `/api/notes/{id}` | Получить заметку |
| `GET` | `/api/notes/{id}/raw` | Получить текст |
| `GET` | `/health` | Health check |

### Пример

```bash
# Создать заметку
curl -X POST http://localhost:8080/api/notes \
  -H "Content-Type: application/json" \
  -d '{"content": "Hello!"}'

# Ответ
{
  "id": "sunny-cat-42",
  "url": "http://localhost:3000/sunny-cat-42",
  "expires_at": "2024-01-15T18:30:00Z"
}
```

## 🐳 Сервисы

| Сервис | Порт | Описание |
|--------|------|----------|
| `frontend` | 3000 | React SPA |
| `backend` | 8080 | Go API |
| `postgres` | 5432 | PostgreSQL |
| `migrations` | - | Goose миграции |
