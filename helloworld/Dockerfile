# First stage: Build the binary
FROM golang:1.20 AS builder

WORKDIR /app
COPY . .

RUN go build -o myapp main.go

# Second stage: Copy the built binary to a minimal image
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/myapp .

CMD ["./myapp"]
