BIN_NAME=ozon_test

build:
		go mod download
		go build -o ${BIN_NAME} .

run psql:
		docker-compose --profile psql  up --build

run inmemory:
		docker-compose --profile inmemory  up --build

test:

clean:
		go clean
		rm ${BIN_NAME}