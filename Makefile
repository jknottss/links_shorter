BIN_NAME=ozon_test

build:
		go mod download
		go build -o ${BIN_NAME} .

psql:
		docker-compose --profile psql  up --build

inmemory:
		docker-compose --profile inmemory  up --build

test:
		go test ./createlink... ./handler

clean:
		go clean
		rm ${BIN_NAME}