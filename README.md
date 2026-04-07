# Copy Paste Service

Простой сервис для быстрого обмена текстом через короткие, человеко-читаемые ссылки.
Заметки автоматически удаляются через 3 часа.

## 🚀 Развертывание (Production)

Проект использует **глобальный экземпляр Traefik v3** для балансировки нагрузки и управления SSL сертификатами. 

### Предварительные требования

1. Работающий [Traefik v3](https://traefik.io/) с доступом к Docker-сети `web` (либо поменяйте имя сети в конце `docker-compose.prod.yml`).
2. Настроенный резолв сертификатов в глобальном Traefik с именем `letsencrypt`.

### Установка

1. Скопируйте шаблон конфигурации:
   ```bash
   cp .env.example .env
   ```
2. Отредактируйте `.env`, указав ваш домен и пароль от базы данных:
   ```env
   DOMAIN=paste.example.com
   DB_PASSWORD=your_secure_password
   ```
3. Запустите скрипт деплоя:
   ```bash
   chmod +x deploy.sh
   ./deploy.sh
   ```

После развертывания ваше приложение и API будут доступны по указанному домену (например, `https://paste.example.com`). Документация API будет доступна по адресу `https://paste.example.com/swagger/`.

---

## 🛠️ Локальная разработка

Для локального запуска (без Traefik) вы можете использовать стандартный docker-compose:

```bash
docker compose up --build
```

После запуска:
- **Frontend**: http://localhost:3000
- **Backend API**: http://localhost:8080
- **Swagger UI**: http://localhost:8080/swagger/

### Запуск фронтенда локально (Vite)

```bash
cd frontend
npm install
npm run dev
```

Создайте файл `frontend/.env.local`:
```env
VITE_API_URL=http://localhost:8080
VITE_APP_URL=http://localhost:5173
```

### Запуск бэкенда локально (Make)

```bash
cd backend
make run             # Запустить backend локально
make build           # Собрать бинарник
make sqlc            # Сгенерировать код sqlc
make migrate-up      # Применить миграции
```

---

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
├── docker-compose.yml       # Локальная оркестрация
├── docker-compose.prod.yml  # Оркестрация для Production (Traefik labels)
├── .env.example             # Шаблон конфигурации
└── deploy.sh                # Скрипт автоматического деплоя
```

## 📡 API

| Метод | Путь | Описание |
|-------|------|----------|
| `POST` | `/api/notes` | Создать заметку |
| `GET` | `/api/notes/{id}` | Получить заметку |
| `GET` | `/api/notes/{id}/raw` | Получить текст в сыром виде |
| `GET` | `/health` | Health check |

### Пример работы с API

```bash
# Создать заметку
curl -X POST https://paste.example.com/api/notes \
  -H "Content-Type: application/json" \
  -d '{"content": "Hello!"}'

# Ответ
{
  "id": "sunny-cat-42",
  "url": "https://paste.example.com/sunny-cat-42",
  "expires_at": "2024-01-15T18:30:00Z"
}
```

## 🐳 Сервисы

| Сервис | Описание |
|--------|----------|
| `frontend` | React SPA (Nginx) |
| `backend` | Go API (Чистая архитектура, HTTP/Router) |
| `postgres` | PostgreSQL 17 |
| `migrations` | Goose миграции |