# Search Engine

A search engine with a crawler, featuring a React frontend and a Go backend.

## Requirements
- `make`
- `docker`
- `docker-compose`

## Build
1. Edit the `env-example` file to configure the required environment variables.  
2. Rename the file to `.env`.  

3. Then, build the project by running:  
```bash
make build
```

## Run
I recommend using [Portainer](https://github.com/portainer/portainer) to manage the stack. You can then choose to launch the stack with or without the crawler, depending on your needs. Additionally, you can dynamically modify the `CRAWLER_START_URL` before starting the crawler.  

To start everything without Portainer:  
```bash
make start
```

To stop all services:  
```bash
make stop
```