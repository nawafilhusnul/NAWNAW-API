# Makefile

# Variables
NEW_DOMAIN_SCRIPT = ./bash/generate.sh

# Targets
.PHONY: create delete

create:
	@$(NEW_DOMAIN_SCRIPT) create $(domain)

delete:
	@$(NEW_DOMAIN_SCRIPT) delete $(domain)

MIGRATION_DIR := $(shell jq -r '.migration.dir' config.json)
DB_DRIVER := $(shell jq -r '.database.driver' config.json)
DB_HOST := $(shell jq -r '.database.host' config.json)
DB_PORT := $(shell jq -r '.database.port' config.json)
DB_NAME := $(shell jq -r '.database.database' config.json)
DB_USERNAME := $(shell jq -r '.database.username' config.json)
DB_PASSWORD := $(shell jq -r '.database.password' config.json)

MYSQL_CONN_STRING := $(DB_DRIVER)://$(DB_USERNAME):$(DB_PASSWORD)@tcp\($(DB_HOST):$(DB_PORT)\)/$(DB_NAME)


migrate:
	@read -p "Enter migration sequence: " seq; \
	echo "running ... migrate create -ext sql -dir $(MIGRATION_DIR) -seq $$seq"; \
	migrate create -ext sql -dir $(MIGRATION_DIR) -seq $$seq

migrate-up:
	echo "running ... migrate -source file://$(MIGRATION_DIR) -database $(MYSQL_CONN_STRING) up"; \
	migrate -source file://$(MIGRATION_DIR) -database $(MYSQL_CONN_STRING) up