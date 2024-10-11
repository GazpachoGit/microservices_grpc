set DATA_SOURCE_URL=root:verysecretpass@tcp(127.0.0.1:3306)/order
set APPLICATION_PORT=3000
set ENV=development
go run ./cmd/main.go

grpcurl -plaintext localhost:3000 describe Order
grpcurl -plaintext -d {\"user_id\":123,\"order_items\":[{\"product_code\":\"prod\",\"quantity\":4,\"unit_price\":12}]} localhost:3000 Order/Create
