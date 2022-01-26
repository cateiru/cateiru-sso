.PHONY: all
all: clean test

.PHONY: test
test:
	docker-compose up -d
	go test -v ./...
	docker-compose down

.PHONY: test-no-cache
test-no-cache:
	docker-compose up -d
	go test -v ./... -count=1
	docker-compose down

.PHONY: dev
test-no-cache:
	docker-compose -f docker-compose.debug.yaml up
