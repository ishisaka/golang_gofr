.PHONY: redis-up
redis-up:
	docker run --name gofr-redis -p 2002:6379 -d redis:7.2

.PHONY: redis-down
redis-down:
	docker stop gofr-redis

.PHONY: run
run:
	go run main.go

.PHONY: test
test:
	go test -v .

.PHONY: build
build:
	docker image build . -t gofr-redis

.PHONY: up
up:
	docker compose up -d

.PHONY: down
down:
	docker compose down
