# Etapa de compilación
FROM golang:1.23.3-alpine3.20 AS builder

WORKDIR /app

# Instalar dependencias para la compilación
RUN apk add --no-cache ca-certificates git build-base

# Copiar go.mod y go.sum para cachear las dependencias
COPY go.mod go.sum ./

# Descargar las dependencias
RUN go mod download

# Copiar el resto del código fuente del backend
COPY . .

# Compilar el binario
# Usar CGO_ENABLED=0 para crear un binario estático,
# compatible con la imagen base alpine
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-s -w" \
    -o /app/bin/server ./cmd/server

# --- Etapa final ---
FROM alpine:3.20.3

# Instalar lo mínimo necesario y Goose para migraciones
RUN apk add --no-cache ca-certificates tzdata wget && \
    wget -O /usr/local/bin/goose https://github.com/pressly/goose/releases/download/v3.24.3/goose_linux_x86_64 && \
    chmod +x /usr/local/bin/goose

WORKDIR /app

# Crear usuario no-root seguro
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

# Copiar el binario y las migraciones desde la etapa de compilación
COPY --from=builder /app/bin/server .
COPY --from=builder /app/migrations ./migrations

# --- ESTA ES LA LÍNEA CLAVE ---
# Ajustar permisos para el binario y las migraciones.
# La bandera -R es crucial para cambiar los permisos de la carpeta y todo su contenido.
RUN chown -R appuser:appgroup server migrations

# Cambiar a usuario no-root
USER appuser

# Puerto expuesto
EXPOSE 8080

# Healthcheck para verificar que la app está viva
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --quiet --tries=1 --spider http://localhost:8080/health || exit 1

# Comando de ejecución: solo inicia el servidor
ENTRYPOINT ["./server"]
