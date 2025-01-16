FROM alpine:latest

RUN apk update && apk add --no-cache \
	bash \
	git \
	go

ARG API_PORT

WORKDIR /root

# Build api
COPY src/* .
RUN go mod tidy
RUN go build -o crawler .

CMD ["./crawler"]
