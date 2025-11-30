// ===========================================
// КОНФИГУРАЦИЯ ПРИЛОЖЕНИЯ
// ===========================================

// API URL - URL бэкенда
// В development: http://localhost:8080
// В production: задать через переменную окружения VITE_API_URL
export const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080';

// APP URL - домен фронтенда для формирования ссылок
// В development: http://localhost:5173
// В production: задать через переменную окружения VITE_APP_URL
// Например: https://paste.example.com
export const APP_URL = import.meta.env.VITE_APP_URL || window.location.origin;

