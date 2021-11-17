FROM golang:1.17-alpine AS builder

WORKDIR /app

# Copy just the files needed for download modules to take advantage of caching in Docker for local development
# go.sum is used for cache invalidation
COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0
RUN GOOS=linux GOARCH=amd64 go build -o /opt/store-ads-service/service ./cmd/main.go

FROM alpine AS final

WORKDIR /bin
COPY --from=builder /opt/store-ads-service /opt/store-ads-service

WORKDIR /opt/store-ads-service
ENTRYPOINT ["/opt/store-ads-service/service"]
