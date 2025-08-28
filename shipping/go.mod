module github.com/islanpedro01/microservices/order

go 1.24.4

require (
	github.com/islanpedro01/microservices-proto/golang/shipping v0.0.0-20250828190131-9285d94049c2
	gorm.io/driver/mysql v1.6.0
)

require (
	golang.org/x/net v0.43.0 // indirect
	golang.org/x/sys v0.35.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250826171959-ef028d996bc1 // indirect
	google.golang.org/protobuf v1.36.8 // indirect
)

replace github.com/islanpedro01/microservices-proto/golang/payment => ../../microservices-proto/golang/payment

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0
	github.com/islanpedro01/microservices-proto/golang/order v0.0.0-20250807035753-c642a4b89af5
	github.com/islanpedro01/microservices-proto/golang/payment v0.0.0-00010101000000-000000000000
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/text v0.28.0 // indirect
	google.golang.org/grpc v1.75.0
	gorm.io/gorm v1.30.0
)
