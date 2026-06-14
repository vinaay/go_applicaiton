# Stage 1: Build
FROM golang:1.25.0 as builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main ./cmd/server/main.go

# Stage 2: Run
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY configs ./configs
CMD ["./main"]