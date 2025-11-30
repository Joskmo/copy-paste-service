# 🚀 Деплой Copy Paste Service

## Требования

- VPS/сервер с Ubuntu 22.04+ (или другой Linux)
- Docker и Docker Compose
- Домен, указывающий на IP сервера
- Открытые порты: 80, 443

## Быстрый старт

### 1. Установка Docker (если не установлен)

```bash
# Установка Docker
curl -fsSL https://get.docker.com | sh

# Добавление пользователя в группу docker
sudo usermod -aG docker $USER

# Перезайдите в систему для применения изменений
exit
```

### 2. Клонирование проекта

```bash
git clone <your-repo-url> copy-paste-service
cd copy-paste-service
```

### 3. Настройка домена

Создайте A-запись в DNS:
```
paste.example.com → IP_ВАШЕГО_СЕРВЕРА
```

Подождите 5-10 минут для распространения DNS.

### 4. Конфигурация

```bash
# Скопируйте пример конфигурации
cp .env.production.example .env

# Отредактируйте файл
nano .env
```

Заполните переменные:
```env
DOMAIN=paste.example.com
ACME_EMAIL=admin@example.com
DB_PASSWORD=ваш_безопасный_пароль
```

### 5. Запуск

```bash
# Сделайте скрипт исполняемым
chmod +x deploy.sh

# Запустите деплой
./deploy.sh
```

### 6. Проверка

```bash
# Статус контейнеров
docker compose -f docker-compose.prod.yml ps

# Логи
docker compose -f docker-compose.prod.yml logs -f

# Проверка SSL
curl -I https://paste.example.com
```

## Архитектура

```
                    ┌─────────────┐
                    │   Traefik   │ ← Let's Encrypt SSL
                    │  (reverse   │
        :80/:443 ── │   proxy)    │
                    └──────┬──────┘
                           │
              ┌────────────┴────────────┐
              │                         │
        ┌─────┴─────┐             ┌─────┴─────┐
        │ Frontend  │             │  Backend  │
        │  (nginx)  │             │   (Go)    │
        │   :80     │             │  :8080    │
        └───────────┘             └─────┬─────┘
                                        │
                                  ┌─────┴─────┐
                                  │ PostgreSQL│
                                  │   :5432   │
                                  └───────────┘
```

## Роутинг

| Path | Сервис |
|------|--------|
| `/api/*` | Backend |
| `/health` | Backend |
| `/swagger/*` | Backend |
| `/*` | Frontend |

## Управление

```bash
# Остановка
docker compose -f docker-compose.prod.yml down

# Перезапуск
docker compose -f docker-compose.prod.yml restart

# Обновление (пересборка)
docker compose -f docker-compose.prod.yml up -d --build

# Логи конкретного сервиса
docker compose -f docker-compose.prod.yml logs -f backend
docker compose -f docker-compose.prod.yml logs -f frontend
docker compose -f docker-compose.prod.yml logs -f traefik

# Очистка (удаление volumes)
docker compose -f docker-compose.prod.yml down -v
```

## Бэкап базы данных

```bash
# Создание бэкапа
docker exec copy-paste-postgres pg_dump -U postgres copypaste > backup.sql

# Восстановление
cat backup.sql | docker exec -i copy-paste-postgres psql -U postgres copypaste
```

## Мониторинг SSL

Сертификаты Let's Encrypt обновляются автоматически Traefik'ом.

Проверка сертификата:
```bash
echo | openssl s_client -servername paste.example.com -connect paste.example.com:443 2>/dev/null | openssl x509 -noout -dates
```

## Troubleshooting

### SSL сертификат не выдаётся

1. Проверьте, что домен указывает на сервер:
   ```bash
   dig paste.example.com
   ```

2. Проверьте порты:
   ```bash
   sudo netstat -tulpn | grep -E ':80|:443'
   ```

3. Проверьте логи Traefik:
   ```bash
   docker compose -f docker-compose.prod.yml logs traefik
   ```

### Backend не запускается

Проверьте логи:
```bash
docker compose -f docker-compose.prod.yml logs backend
```

### Миграции не применились

```bash
docker compose -f docker-compose.prod.yml logs migrations
```

## Безопасность

- ✅ Все соединения шифруются (HTTPS)
- ✅ HTTP автоматически редиректится на HTTPS
- ✅ База данных недоступна извне
- ✅ Заметки удаляются через 3 часа
- ⚠️ Используйте надёжный пароль для базы данных
- ⚠️ Регулярно обновляйте образы

