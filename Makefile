BIN_NAME=ozon_test

build:
		go mod download
		go build -o ${BIN_NAME} ./cmd

psql:
		docker-compose --profile psql  up --build

inmemory:
		docker-compose --profile inmemory  up --build

test:
		go test ./internal/createlink ./internal/handler

clean:
		go clean
		rm ${BIN_NAME}