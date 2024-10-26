#run main application
.PHONY: run/api
run/api:
	@echo 'Running Appliation'
	@go run ./cmd/api/ -port=4000 -env=development -db-dsn=$(SIGNIN_DB_DSN)


#enter database
.PHONY: db/sql
db/psql:
	psql $(SIGNIN_DB_DSN) -U signin

#create migrations
.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating Database Migrations for $(name)'
	migrate create -seq -ext=.sql -dir=./migrate $(name)

#up migrations
.PHONY: db/migrations/up
db/migrations/up:
	@echo 'running up migrations'
	migrate -path ./migrate -database $(SIGNIN_DB_DSN) up

#down migrations
.PHONY: db/migrations/down
db/migrations/down:
	@echo 'running up migrations'
	migrate -path ./migrate -database $(SIGNIN_DB_DSN) down