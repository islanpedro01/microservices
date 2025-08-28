FROM golang:1.24.5 AS builder
WORKDIR /protos
COPY ../microservices-proto .
WORKDIR /usr/src/app
COPY . .
RUN cd ./microservices/payment && \
    go mod edit -replace=github.com/islanpedro01/microservices-proto/golang=/protos && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /payment ./cmd/main.go

FROM scratch
COPY --from=builder /payment ./payment
CMD ["./payment"]