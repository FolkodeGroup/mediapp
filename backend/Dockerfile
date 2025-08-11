# backend/Dockerfile

# --- Etapa de compilación ---
FROM golang:1.23.3-alpine3.20 AS builder

WORKDIR /app

# Instalar dependencias necesarias para la compilación
RUN apk add --no-cache ca-certificates git build-base

# Copiar go.mod y go.sum para aprovechar la cache de dependencias
COPY go.mod go.sum ./

# Descargar dependencias
RUN go mod download

# Copiar el código fuente
COPY . .

# Compilar el binario
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-s -w" \
    -o /app/bin/server ./cmd/server

# --- Etapa final ---
FROM alpine:3.20.3

# Instalar lo mínimo necesario en la imagen final
RUN apk add --no-cache ca-certificates tzdata wget

WORKDIR /app

# Crear usuario no-root seguro
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

# Copiar solo el binario desde la etapa de compilación
COPY --from=builder /app/bin/server .

# Ajustar permisos
RUN chown appuser:appgroup server

# Cambiar a usuario no-root
USER appuser

# Puerto expuesto
EXPOSE 8080

# Healthcheck para verificar que la app está viva
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --quiet --tries=1 --spider http://localhost:8080/health || exit 1

# Comando de ejecución
CMD ["./server"]
