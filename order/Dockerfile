FROM golang:1.24.5 AS builder
WORKDIR /protos
COPY ../microservices-proto .
WORKDIR /usr/src/app
COPY . .
RUN cd ./microservices/order && \
    go mod edit -replace=github.com/islanpedro01/microservices-proto/golang=/protos && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /order ./cmd/main.go

FROM scratch
COPY --from=builder /order ./order
CMD ["./order"]