#!/bin/bash
# ===========================================
# СКРИПТ ДЕПЛОЯ Copy Paste Service
# ===========================================

set -e

echo "🚀 Copy Paste Service - Деплой"
echo "================================"

# Проверка .env файла
if [ ! -f .env ]; then
    echo "❌ Файл .env не найден!"
    echo ""
    echo "Создайте его:"
    echo "  nano .env"
    echo ""
    echo "Содержимое:"
    echo "  DOMAIN=paste.example.com"
    echo "  DB_PASSWORD=secure_password"
    exit 1
fi

# Загрузка переменных
set -a
source .env
set +a

# Проверка обязательных переменных
if [ -z "$DOMAIN" ]; then
    echo "❌ Переменная DOMAIN не задана в .env"
    exit 1
fi

if [ -z "$DB_PASSWORD" ]; then
    echo "❌ Задайте DB_PASSWORD в .env!"
    exit 1
fi

echo "📋 Конфигурация:"
echo "   Домен: $DOMAIN"
echo ""

# Остановка старых контейнеров
echo "⏹️  Остановка старых контейнеров..."
docker compose -f docker-compose.prod.yml down 2>/dev/null || true

# Сборка образов
echo "🔨 Сборка образов..."
docker compose -f docker-compose.prod.yml build

# Запуск
echo "🚀 Запуск сервисов..."
docker compose -f docker-compose.prod.yml up -d

# Ожидание запуска
echo "⏳ Ожидание запуска сервисов..."
sleep 15

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
