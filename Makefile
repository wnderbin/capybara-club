go-run-service-user:
	CONFIG_PATH=./user-service/config/config.yaml go run ./cmd/user-service/main.go
go-run-service-restaurant:
	CONFIG_PATH=./restaurant-service/config/config.yaml go run ./cmd/restaurant-service/main.go
go-run-service-admin:
	CONFIG_PATH=./admin-service/config/config.yaml go run ./cmd/admin-service/main.go
go-run-service-order:
	CONFIG_PATH=./order-service/config/config.yaml go run ./cmd/order-service/main.go