compose-all:
	@docker-compose up -d --build

compose-down:
	@docker-compose down
	@docker-compose rm

compose-logs:
	@docker logs cars_postgres_1
	@docker logs cars_redis_1
	@docker logs -f cars_cars_1

compose-deps:
	@docker-compose up -d --build postgres redis

delete-db-volume:
	@docker volume rm cars_data