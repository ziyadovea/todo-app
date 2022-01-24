.PHONY: build
build:
	docker-compose build todo-app

run:
	docker-compose up todo-app

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: migrate
migrate:
	migrate -path ./schema -database 'postgres://postgres:0000@localhost:5436/postgres?sslmode=disable' up
