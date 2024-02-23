up:
	docker compose up -d

down:
	docker compose down

start:
	docker compose start

restart:
	docker compose restart

stop:
	docker compose stop

bash-app:
	docker compose exec -it app /bin/bash

bash-db:
	docker compose exec -it db /bin/bash

psql:
	docker compose exec -it db psql -U postgres -d twinte_db

tidy:
	go mod tidy

test:
	go test -count=1 ./...

# e.g. make migrate-create name=foo
migrate-create:
	migrate create -dir ./db/migrations -ext sql -seq -digits 6 ${name}

# e.g. make migrate-up db_url=${DB_URL}
migrate-up:
	migrate -database ${db_url} -path ./db/migrations up;

# e.g. make migrate-down db_url=${DB_URL}
migrate-down:
	migrate -database ${db_url} -path ./db/migrations down -all

# e.g. make migrate-force db_url=${DB_URL} version=1
migrate-force:
	migrate -database ${db_url} -path ./db/migrations force ${version}

gorm-gen:
	rm -rf ./db/gen && go run ./db/gorm_gen.go

buf-gen:
	rm -rf ./api/rpcgen && buf generate --template ./twinte-proto/buf.gen.yaml ./twinte-proto
