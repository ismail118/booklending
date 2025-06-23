install-mysql:
	docker run --name mysql8 -p 3306:3306  -e MYSQL_ROOT_PASSWORD=password -d mysql:8

create-db:
	docker exec mysql8 mysql -u root -p'password' -e "CREATE DATABASE IF NOT EXISTS booklending;"

install-migrate:
	brew install golang-migrate

new-migrate:
	migrate create -ext sql -dir db/migration -seq $(name)

migrate-up:
	migrate -path db/migration -database "mysql://root:password@tcp(localhost:3306)/booklending?tls=false" up

migrate-down:
	migrate -path db/migration -database "mysql://root:password@tcp(localhost:3306)/booklending?tls=false" down

migrate-force:
	migrate -path db/migration -database "mysql://root:password@tcp(localhost:3306)/booklending?tls=false" force $(version)

.PHONY: install-mysql create-db install-migrate new-migrate migrate-up migrate-down