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
dev:
	docker-compose -f docker-compose.yaml -f docker-compose.debug.yaml up

.PHONY: deploy-index
deploy-index:
	gcloud datastore indexes create ./build/index.yaml

.PHONY: clean-index
clean-index:
	gcloud datastore indexes cleanup ./build/index.yaml
