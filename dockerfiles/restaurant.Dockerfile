FROM golang:1.24.2-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY . .
RUN go mod download

ENV CONFIG_PATH=./cmd/restaurant-service/config/config.yaml

CMD ["go", "run", "./cmd/restaurant-service/main.go"]