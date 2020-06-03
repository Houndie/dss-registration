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
COPY docker/startup.sh .
CMD ./startup.sh
