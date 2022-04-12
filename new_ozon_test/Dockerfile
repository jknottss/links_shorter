FROM golang:latest
WORKDIR /new_ozon_test
COPY . .
RUN go mod download
RUN go build -o ozon_test .
CMD ./ozon_test