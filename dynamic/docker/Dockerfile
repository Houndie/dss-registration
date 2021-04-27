FROM golang:1.16

RUN apt update && apt install -y netcat
RUN wget https://github.com/cortesi/modd/releases/download/v0.8/modd-0.8-linux64.tgz && tar xzf modd-0.8-linux64.tgz && cp modd-0.8-linux64/modd /go/bin/ && rm -rf modd-0.8-linux64.tgz modd-0.8-linux64
RUN mkdir /app
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
CMD docker/startup.sh
