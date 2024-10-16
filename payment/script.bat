set DATA_SOURCE_URL=root:verysecretpass@tcp(127.0.0.1:3306)/payment
set APPLICATION_PORT=3001
set ENV=development
go run ./cmd/main.go
