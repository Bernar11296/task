FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN apk add build-base && go build -o test cmd/test-task/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app .
CMD ["./test"]