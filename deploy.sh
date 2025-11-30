#!/bin/bash
# ===========================================
# СКРИПТ ДЕПЛОЯ Copy Paste Service
# ===========================================
#
# Использование:
#   chmod +x deploy.sh
#   ./deploy.sh
#
# Требования:
#   - Docker и Docker Compose установлены
#   - Файл .env заполнен
#   - Домен указывает на IP сервера
#   - Порты 80 и 443 открыты

set -e

echo "🚀 Copy Paste Service - Деплой"
echo "================================"

# Проверка .env файла
if [ ! -f .env ]; then
    echo "❌ Файл .env не найден!"
    echo ""
    echo "Создайте его из примера:"
    echo "  cp .env.production.example .env"
    echo "  nano .env"
    exit 1
fi

# Загрузка переменных
source .env

# Проверка обязательных переменных
if [ -z "$DOMAIN" ]; then
    echo "❌ Переменная DOMAIN не задана в .env"
    exit 1
fi

if [ -z "$ACME_EMAIL" ]; then
    echo "❌ Переменная ACME_EMAIL не задана в .env"
    exit 1
fi

if [ "$DB_PASSWORD" = "CHANGE_ME_TO_SECURE_PASSWORD" ]; then
    echo "❌ Измените DB_PASSWORD в .env!"
    exit 1
fi

echo "📋 Конфигурация:"
echo "   Домен: $DOMAIN"
echo "   Email: $ACME_EMAIL"
echo ""

# Остановка старых контейнеров
echo "⏹️  Остановка старых контейнеров..."
docker compose -f docker-compose.prod.yml down 2>/dev/null || true

# Сборка и запуск
echo "🔨 Сборка образов..."
docker compose -f docker-compose.prod.yml build

echo "🚀 Запуск сервисов..."
docker compose -f docker-compose.prod.yml up -d

# Ожидание запуска
echo "⏳ Ожидание запуска сервисов..."
sleep 10

# Проверка статуса
echo ""
echo "📊 Статус контейнеров:"
docker compose -f docker-compose.prod.yml ps

echo ""
echo "✅ Деплой завершён!"
echo ""
echo "🌐 Ваше приложение доступно по адресу:"
echo "   https://$DOMAIN"
echo ""
echo "📝 Полезные команды:"
echo "   Логи:     docker compose -f docker-compose.prod.yml logs -f"
echo "   Статус:   docker compose -f docker-compose.prod.yml ps"
echo "   Стоп:     docker compose -f docker-compose.prod.yml down"
echo "   Рестарт:  docker compose -f docker-compose.prod.yml restart"

