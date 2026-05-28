# -------------------------
# Stage 1: Build Stage
# -------------------------
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o atlas-server ./cmd/server

# -------------------------
# Stage 2: Runtime Stage
# -------------------------
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/atlas-server .

EXPOSE 8080

CMD ["./atlas-server"]