# Build stage
FROM golang:alpine as builder

ARG PATH="$PATH:$(go env GOPATH)/bin"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main cmd/main.go

## Run stage
FROM alpine:latest as run

WORKDIR /app

COPY --from=builder /app/main .
COPY example.env .

EXPOSE 8080

CMD ["./main"]