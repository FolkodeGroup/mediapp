# frontend/Dockerfile
FROM node:22.11.0-alpine3.20 AS builder

WORKDIR /app

# Instalar dependencias del sistema
RUN apk add --no-cache libc6-compat

# Copiar package files
COPY package*.json ./

# Instalar dependencias
RUN npm ci && npm cache clean --force

# Copiar código fuente
COPY . .

# Construir la aplicación
RUN npm run build

# Etapa de producción
FROM nginx:1.27.2-alpine

# Copiar archivos construidos
COPY --from=builder /app/dist /usr/share/nginx/html

# Copiar configuración de nginx
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Exponer puerto
EXPOSE 80

# Iniciar nginx (sin crear usuario, usar el default)
CMD ["nginx", "-g", "daemon off;"]