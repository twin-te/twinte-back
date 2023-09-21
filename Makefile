include .env

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

logs:
	docker compose logs app

bash:
	docker compose exec -it app /bin/bash

tidy:
	go mod tidy

test:
	go test -count=1 ./...

serve:
	go run . serve

buf-gen:
	rm -rf ./api/gen && buf generate ./twinte-proto

# ex.) make migrate-create name=foo
migrate-create:
	migrate create -dir ./db/migrations -ext sql -seq -digits 6 ${name}

migrate-up:
	migrate -database $(DB_URL) -path ./db/migrations up

migrate-down:
	migrate -database $(DB_URL) -path ./db/migrations down -all

# ex.) make migrate-force version=1
migrate-force:
	migrate -database $(DB_URL) -path ./db/migrations force ${version}

sqlboiler:
	sqlboiler psql -c sqlboiler.toml
