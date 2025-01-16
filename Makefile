ifneq (,$(wildcard .env))
	include .env
	export
endif

all: build

build: _cp-certificates
	docker-compose --project-name search-engine up --build --no-start

start: _cp-certificates
	docker-compose --project-name search-engine up --build

stop:
	docker-compose down

_cp-certificates:
	cp $(SSL_CERTIFICATE) api/certificat.crt
	cp $(SSL_CERTIFICATE_KEY) api/private.key
	cp $(SSL_CERTIFICATE) nginx/certificat.crt
	cp $(SSL_CERTIFICATE_KEY) nginx/private.key

.PHONY: all build start stop