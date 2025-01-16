FROM alpine:latest

RUN apk update && apk add --no-cache \
	bash \
	git \
	go

ARG API_PORT

WORKDIR /root

COPY certificat.crt /cert/
COPY private.key /cert/

# Build api
COPY src/* .
RUN go mod tidy
RUN go build -o api .

EXPOSE ${API_PORT}

CMD ["./api"]
