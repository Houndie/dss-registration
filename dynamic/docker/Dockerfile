FROM golang:1.14.3-alpine3.11 as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo ./cmd/server/

FROM alpine:3.11

WORKDIR /root/
COPY --from=builder /app/server .
COPY docker/config.toml .

EXPOSE 80

ENTRYPOINT ["./server"]
