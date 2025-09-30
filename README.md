# Handmade Shop Bot 🧶

### Cодержание
- [📄 Краткое описание](#-крактое-описание)
- [🚀 Стек технологий](#-стек-технологий)
- [🛠️ Установка и запуск](#️-установка-и-запуск)
- [🧪 Тесты](#-тесты)
- [📜 Makefile команды](#-makefile-команды)
- [📈 Prometheus config](#-prometheus-config)
---
## 📄Крактое Описание
Сервис и Telegram-бот для интернет-магазина hand-made: каталог товаров, оформление заказов, интеграция с платёжным шлюзом, админ-панель через бота. Бот позволяет просматривать каталог, покупать товары, админам — добавлять/редактировать продукты и управлять заказами.
Основные возможности:
- 📦 Управление товарами (CRUD)
- 🛒 Создание и просмотр заказов
- 🔗 Интеграция с платёжным сервисом (создание платежа, проверка статуса)
- 🔔 Уведомления админа и покупателя через Telegram
- 📊 Экспорт метрик для Prometheus
  
## 🚀 Стек технологий
- Go (чистые usecase-слои)
- PostgreSQL — основная БД
- Telebot v4 — Telegram SDK
- Docker & docker-compose — локальный запуск
- Prometheus / Grafana — метрики и дашборды
- estify/mock — unit-тесты и моки

## 🛠️ Установка и запуск

1. Клонировать проект
```bash
git clone https://github.com/iviv660/handmade-shop-bot.git
cd handmade-shop-bot
```
2. Создать бота через BotFather и вставить токен в .env
3. Запуск стека:
```bash
make up
```
Поднимутся: 
- service — твой Go-сервис с метриками (:2112/metrics);
- prometheus — интерфейс на http://localhost:9090;
- postgres — база данных;
- postgres_exporter — метрики PostgreSQL (:9187/metrics). 

Проверка:
- Метрики сервиса: http://localhost:2112/metrics 
- Метрики Postgres: http://localhost:9187/metrics 
- Prometheus UI: http://localhost:9090
4. Остановка:
```bash
make down 
```

## 🧪 Тесты
Запуск unit-тестов:
```bash 
make test
```

## 📜 Makefile команды

- make up — поднять весь стек (docker compose up -d)
- make down — остановить стек
- make logs — логи всех контейнеров
- make ps — статус контейнеров
- make test — запустить Go-тесты

## 📈 Prometheus config
prometheus/prometheus.yml уже содержит scrape-jobs для:
- самого Prometheus,
- Go-сервиса,
- Postgres Exporter.
