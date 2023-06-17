ifneq (,$(wildcard config/.env))
    include config/.env
    export
endif

run: stopair
	go run .
dbreset: dbrollback dbmigrate dbmock
stopair:
	-kill $$(lsof -ti:$(PORT)) 2>/dev/null
dbcreate:
	migrate create -ext sql -dir db/migrations/ -seq qrthrough
dbmigrate:
	migrate -database postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_NAME)?sslmode=disable -path ./db/migrations up
dbrollback:
	migrate -database postgres://root:arnan1234@localhost:5433/qrthrough?sslmode=disable -path ./db/migrations down

dbscript:
	go run ./cmd/automigrate/
dbmock:
	go run ./cmd/mockdb.go

dbalumni:
	go run ./cmd/alumni_csv/main.go

