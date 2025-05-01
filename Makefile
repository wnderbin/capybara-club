go-run-service-user:
	CONFIG_PATH=./cmd/user-service/config/config.yaml go run ./cmd/user-service/main.go
go-run-service-restaurant:
	CONFIG_PATH=./cmd/restaurant-service/config/config.yaml go run ./cmd/restaurant-service/main.go
go-run-service-admin:
	CONFIG_PATH=./cmd/admin-service/config/config.yaml go run ./cmd/admin-service/main.go
go-run-service-order:
	CONFIG_PATH=./cmd/order-service/config/config.yaml go run ./cmd/order-service/main.go