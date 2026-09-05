# Stage 1: Build the binary using heavy Go compiler
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY src/main.go .
RUN go mod init enterprise-app && go build -o enterprise-app main.go

# Stage 2: Super lightweight production image
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/enterprise-app .
EXPOSE 8080
CMD ["./enterprise-app"]
