# Переменные
DEV_COMPOSE = docker compose -f docker-compose.dev.yml
PROD_COMPOSE = docker compose -f docker-compose.prod.yml

# Запуск в dev-режиме
dev:
	$(DEV_COMPOSE) up --build

# Запуск в prod-режиме
prod:
	$(PROD_COMPOSE) up --build -d

# Остановка всех контейнеров (dev + prod)
down:
	docker compose -f docker-compose.dev.yml -f docker-compose.prod.yml down

# Очистка (остановка + удаление контейнеров, сетей и томов)
clean:
	docker compose -f docker-compose.dev.yml -f docker-compose.prod.yml down -v