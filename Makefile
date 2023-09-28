

run:
	go run main.go app

migrate:
	go run main.go migrate

seed:
	go run main.go seed

# Генерим прото файлы
proto:
	./scripts/gen_wallet_storage_proto.sh
	./scripts/gen_wallet_storage_ds_proto.sh

migrate-create:
	migrate create -ext sql -dir ./cmd/migrate/migrations/ -seq $(name)
