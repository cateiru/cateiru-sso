.PHONY: all
all: clean test

.PHONY: test
test:
	docker-compose up -d
	go test -v ./...
	docker-compose down --rmi all
