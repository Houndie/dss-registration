FROM golang:1.14.3-alpine3.11 as builder

ARG FRONTEND

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo ./cmd/server/

FROM alpine:3.11
RUN apk add --no-cache netcat-openbsd

WORKDIR /root/
COPY --from=builder /app/server .

EXPOSE 80
CMD until nc -z $DSS_STORAGE_HOST $DSS_STORAGE_PORT; do \
		echo "Waiting for database ..."; \
		sleep 1; \
	done; \
	./server
