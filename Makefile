include .env
export

export PROJECT_ROOT=$(shell pwd)


pg-up:
	@docker-compose up app-postges

pg-down:
	@docker-compose down app-postges

pg-cleanup:
	@read -p "Очистить все volume файлы окружения? [y/N]: " ans; \
	if ["$$ans" = "y"]; then \
		docker-compose down app-postges && \
		rm -rf out/pgdata; \
		echo "Файлы окружения очищены"; \
	else \
		echo "Очистка окружения отменена"; \
	fi

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Отсутствует обязательный параметр 'seq'"; \
		exit 1; \
	fi

	@docker compose run --rm app-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)";

migrate-up:
	@nake migrate-action action=up

migrate-down:
	@nake migrate-action action=down

migrate-action:
	@if [-z "$(action)"]; then \
		echo "Отсутствует обязательный параметр 'action'";
		exit 1; \
	fi;

	docker compose run --rm app-postgres-migrate \
		-path migrations \
		-database postrges://${POSTGRES_USER}:${POSTGRES_PASSWORD}app-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		"$(action)"


swagger-up:
	docker run --rm -d -e PORT=8080 -p 8080:8080 \
	-v ${PROJECT_ROOT}/swagger.conf:/etc/nginx/conf.d/default.conf \
	docker.swagger.io/swaggerapi/swagger-editor
