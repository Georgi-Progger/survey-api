FROM golang:latest as builder

ARG CGO_ENABLED=0
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY .env .env

RUN go build -o /app/main ./cmd/hrtechno

CMD ["/app/main"]
