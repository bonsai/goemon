# Multi-stage build for Goemon Orchestrator
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o goemon ./src/cmd/goemon

# Final image
FROM alpine:latest
RUN apk add --no-cache ca-certificates python3 py3-pip

WORKDIR /app
COPY --from:builder /app/goemon .
COPY --from:builder /app/src/core/agent/python ./src/core/agent/python
COPY --from:builder /app/src/core/bridge/python-bridge.go ./src/core/bridge/python-bridge.go

# Expose port for future API (ADR-004)
EXPOSE 8080

CMD ["./goemon"]
