FROM golang:latest

LABEL maintainer="janduymonroe <janduymonroe@gmail.com>"

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/api

EXPOSE 8080

CMD ["./main"]