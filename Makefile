include .env

MAIN_DB_DSN=$(MAIN_DB_DIALECT)://$(MAIN_DB_USERNAME):$(MAIN_DB_PASSWORD)@$(MAIN_DB_HOST):$(MAIN_DB_PORT)/$(MAIN_DB_NAME)?sslmode=$(MAIN_DB_SSLMODE)

create-schema:
	go run -mod=mod entgo.io/ent/cmd/ent new $(name)

generate-ent:
	go generate ./ent

create-migration:
	atlas migrate diff $(name) --dir "file://ent/migrate/migrations" --to "ent://ent/schema" --dev-url "docker://postgres/17.2/$(MAIN_DB_NAME)_shadow"

apply-migration:
	atlas migrate apply --dir "file://ent/migrate/migrations" --url "$(MAIN_DB_DSN)"