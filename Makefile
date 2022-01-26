VERSION=1.0.0

irita-api:
	go build -o bin/irita-api

docker:
	docker build -t irita-api:$(VERSION) -f docker/Dockerfile .
	docker tag irita-api:$(VERSION) irita-api:latest

start:
	docker-compose -f docker/docker-compose.yaml up -d

stop:
	docker-compose -f docker/docker-compose.yaml down -v

.PHONY: docker