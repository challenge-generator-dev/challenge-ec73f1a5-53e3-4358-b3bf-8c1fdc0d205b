# Etapa de compilación
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Instalar herramientas necesarias para compilación
RUN apk add --no-cache git make protobuf

# Copiar go.mod y go.sum primero para aprovechar cache
COPY go.mod go.sum ./
RUN go mod download

# Copiar código fuente
COPY . .

# Generar código protobuf
RUN protoc --go_out=. --go-grpc_out=. pkg/proto/invoice.proto

# Compilar el binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /invoice-service ./cmd/invoice-service

# Etapa final - imagen minima
FROM alpine:3.19

WORKDIR /app

# Instalar certificadas necesarias para Alpine
RUN apk add --no-cache ca-certificates openssl

# Copiar el binary desde la etapa de build
COPY --from=builder /invoice-service /app/invoice-service

# Copiar configuración si existe
COPY config/ /app/config/

# Exponer puerto del servicio gRPC
EXPOSE 50051

# Health check para Kubernetes
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD grpc_health_probe -addr=localhost:50051 || exit 1

# Instalar grpc_health_probe
RUN apk add --no-cache grpc-tools

# Ejecutar el servicio
CMD ["/app/invoice-service"]