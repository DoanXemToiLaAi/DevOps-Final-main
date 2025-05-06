# Build stage
FROM golang:1.20-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app

# Final stage
FROM alpine:latest

COPY --from=builder /build/app /usr/local/bin/app
RUN chmod +x /usr/local/bin/app

EXPOSE 4200

CMD ["/usr/local/bin/app"]
