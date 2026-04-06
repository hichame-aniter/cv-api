# Makefile
.PHONY: help docker-test-build docker-test-run

help:
	@echo "Available commands:"
	@echo "	make test				- Run tests and save coverage.out & coverage.html"
	@echo "	make docker-build-test	- Build tester image"
	@echo "	make docker-run-test	- Run tester image"

test:
	go test -v ./... -coverpkg=./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

docker-test-build:
	docker build --progress=plain --target tester -t cv-api-backend-test .

docker-test-run: docker-test-build
	docker stop cv-api-backend-test || true
	docker rm cv-api-backend-test || true
	docker run --name cv-api-backend-test cv-api-backend-test:latest