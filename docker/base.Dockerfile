FROM golang:1.24.0-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o server ./cmd/app/main.go
RUN ls -l /app/server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .

ENTRYPOINT ["./server"]