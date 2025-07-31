# Build stage
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY src/ ./src/
WORKDIR /app/src
RUN go mod download
RUN go build -o /app/cp-tracker-be main.go

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/cp-tracker-be .
EXPOSE 8080
CMD ["./cp-tracker-be"]
