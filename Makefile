install-mysql:
	docker run --name mysql8 -p 3306:3306  -e MYSQL_ROOT_PASSWORD=password -d mysql:8 mysqld --mysql-native-password=ON

bash-mysql:
	docker exec -it mysql8 bash
# fix err: Error 1045 (28000): Access denied for user 'mysql'@'172.17.0.1' (using password: YES)

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

go-test:
	go test ./... -v

build-image:
	docker build -t booklending:latest .

compose-up:
	docker compose up --build

gomock-install:
	go install github.com/golang/mock/mockgen@v1.6.0

mock:
	mockgen -package mockdb -destination db/mock/querier.go github.com/ismail118/booklending/db/sql Querier

.PHONY: install-mysql create-db install-migrate new-migrate migrate-up migrate-down bash-mysql go-test build-image compose-up