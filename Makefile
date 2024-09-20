include .env
export

createmigration:
ifeq ($(name),)
	@echo "Please, provide a migration name. Ex: make createmigration name=my_migration"
	@exit 1
endif
	migrate create -ext=sql -dir=internal/infra/database/postgres/sql/migrations -seq $(name)