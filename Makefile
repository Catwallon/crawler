all: build

build:
	docker-compose --project-name search-engine build

start:
	docker-compose --project-name search-engine up --build

stop:
	docker-compose down

.PHONY: all build start stop