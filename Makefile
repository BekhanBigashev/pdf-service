up:
	docker compose up

ps:
	$(PROD_COMPOSE) ps

# Остановка всех контейнеров
down:
	docker compose down

# Очистка (остановка + удаление контейнеров, сетей и томов)
clean:
	docker compose down -v