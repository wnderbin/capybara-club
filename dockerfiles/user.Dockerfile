FROM golang:1.24.2-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY . .
RUN go mod download

CMD ["go", "run", "./cmd/user-service/main.go"]