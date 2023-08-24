compose-all:
	@docker-compose up --build

compose-down:
	@docker-compose down
	@docker-compose rm

compose-deps:
	@docker-compose up -d --build postgres redis

delete-db-volume:
	@docker volume rm cars_data