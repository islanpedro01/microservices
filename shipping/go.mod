module github.com/islanpedro01/microservices/shipping

go 1.24.4

require github.com/islanpedro01/microservices-proto/golang/shipping v0.0.0-20250828190131-9285d94049c2

require (
	golang.org/x/net v0.43.0 // indirect
	golang.org/x/sys v0.35.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250826171959-ef028d996bc1 // indirect
	google.golang.org/protobuf v1.36.8 // indirect
)

replace github.com/islanpedro01/microservices-proto/golang/payment => ../../microservices-proto/golang/payment

require (
	golang.org/x/text v0.28.0 // indirect
	google.golang.org/grpc v1.75.0
)
