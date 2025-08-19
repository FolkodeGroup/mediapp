# 🏥 MediApp - Sistema de Gestión Médica

## � Guía de Desarrollo Actualizada (Docker Compose v2.39+)

### ⚠️ Importante: Usamos Docker Compose moderno
Este proyecto usa **Docker Compose v2.39+** con sintaxis moderna. No necesitas instalar `docker-compose` por separado si tienes Docker Desktop o Docker Engine reciente.

### 🔧 Verificar tu versión de Docker Compose
```bash
# Verificar versión (debería ser v2.39+)
docker compose version

# Si tienes Docker Desktop, ya tienes la versión correcta
# Si usas Linux server, asegúrate de tener Docker Engine reciente
```

## 🚀 Cómo ejecutar en modo desarrollo

### Opción 1: Desarrollo con Docker (Recomendado) 🐳

```bash
# Levantar todo el stack de desarrollo
docker compose -f docker-compose.dev.yml up --build

# O en segundo plano
docker compose -f docker-compose.dev.yml up --build -d

# Para parar
docker compose -f docker-compose.dev.yml down
```

### Opción 2: Desarrollo Local

1. **Frontend**:
   ```bash
   cd frontend
   npm install
   npm run dev
   ```
   El frontend estará disponible en: http://localhost:3000

2. **Backend** (en otra terminal):
   ```bash
   cd backend
   go mod download
   go run cmd/server/main.go
   ```
   El backend estará disponible en: http://localhost:8080

   **Nota**: En desarrollo local necesitarás configurar las variables de entorno para Supabase.

### Opción 3: Scripts desde la raíz

```bash
# Desarrollo frontend únicamente
npm run dev

# Instalar dependencias
npm run install:frontend
npm run install:backend

# Build para producción
npm run build

# Linting y formateo
npm run lint
npm run format
```

## �️ Base de Datos

**Importante**: Este proyecto usa **Supabase** como base de datos compartida, no PostgreSQL local.

- **Proveedor**: Supabase (PostgreSQL en la nube)
- **Conectividad**: 100% (12/12 tablas)
- **Configuración**: Variables en `.env`
- **Beneficios**: Colaborativo, sin setup local, datos compartidos

## 📁 Estructura del Proyecto

```
mediapp/
├── frontend/                    # React + TypeScript + Vite
│   ├── src/
│   ├── public/
│   ├── Dockerfile.dev          # Desarrollo
│   └── vite.config.ts
├── backend/                    # Go + Gin + Supabase
│   ├── cmd/server/
│   ├── internal/
│   │   ├── handlers/           # API handlers
│   │   ├── db/                 # Supabase connection
│   │   └── auth/               # JWT auth
│   ├── Dockerfile.dev          # Desarrollo
│   └── .air.toml               # Hot reload
├── docker-compose.dev.yml      # Desarrollo (v2.39+)
├── .env                        # Variables de Supabase
├── SETUP-COLABORATIVO.md       # Guía del equipo
└── API-ENDPOINTS.md            # Documentación API
```

## 🔧 Configuración de Desarrollo

### Frontend (Vite + React)
- **Puerto**: 3000
- **Hot Reload**: ✅ (Watch mode Docker)
- **Proxy API**: Configurado automáticamente

### Backend (Go + Air + Supabase)
- **Puerto**: 8080
- **Hot Reload**: ✅ (Air + Watch mode Docker)
- **Base de datos**: Supabase PostgreSQL
- **Conectividad**: 100% (12/12 tablas)

### Docker Compose v2.39+ Features
- **Watch mode**: Cambios automáticos sin rebuild
- **IPv4 networking**: Conectividad optimizada
- **Environment files**: Variables centralizadas
- **Health checks**: Verificación automática

## 🌐 URLs de Desarrollo

- **Frontend**: http://localhost:3000
- **Backend API**: http://localhost:8080
- **Health Check**: http://localhost:8080/health
- **Swagger Docs**: http://localhost:8080/swagger/index.html
- **Conectividad**: http://localhost:8080/api/v1/connect/all-tables

## 🛠️ Herramientas de Desarrollo

- **Docker Compose v2.39+**: Orquestación moderna
- **Vite**: Build tool rápido para React
- **Air**: Hot reload para Go
- **Supabase**: Base de datos colaborativa
- **JWT**: Autenticación
- **Zap**: Structured logging
- **pgx/v5**: Driver PostgreSQL optimizado

## 📝 Comandos Útiles (Docker Compose v2.39+)

### Gestión de Servicios
```bash
# Levantar servicios
docker compose -f docker-compose.dev.yml up

# Levantar en background
docker compose -f docker-compose.dev.yml up -d

# Rebuild y levantar
docker compose -f docker-compose.dev.yml up --build

# Parar servicios
docker compose -f docker-compose.dev.yml down

# Parar y limpiar volúmenes
docker compose -f docker-compose.dev.yml down -v
```

### Logs y Debugging
```bash
# Ver logs en tiempo real
docker compose -f docker-compose.dev.yml logs -f

# Logs del backend únicamente
docker compose -f docker-compose.dev.yml logs -f backend-dev

# Logs del frontend únicamente
docker compose -f docker-compose.dev.yml logs -f frontend-dev

# Ver estado de contenedores
docker compose -f docker-compose.dev.yml ps
```

### Acceso a Contenedores
```bash
# Acceder al contenedor del backend
docker compose -f docker-compose.dev.yml exec backend-dev sh

# Acceder al contenedor del frontend
docker compose -f docker-compose.dev.yml exec frontend-dev sh

# Ejecutar comando en backend
docker compose -f docker-compose.dev.yml exec backend-dev go version
```

### Limpieza y Mantenimiento
```bash
# Rebuild completo sin cache
docker compose -f docker-compose.dev.yml build --no-cache

# Recrear contenedores
docker compose -f docker-compose.dev.yml up --force-recreate

# Limpiar sistema Docker
docker system prune -f

# Limpiar redes Docker
docker network prune -f
```

## 🧪 Testing y Verificación

### Verificar que todo funciona
```bash
# Health check
curl http://localhost:8080/health

# Conectividad Supabase (debería ser 100%)
curl -s http://localhost:8080/api/v1/connect/all-tables | jq '.summary'

# Inspeccionar tabla específica
curl -s "http://localhost:8080/api/v1/inspect/tables?table=pacientes" | jq '.'
```

### Verificar Watch Mode
```bash
# Hacer un cambio en cualquier archivo .go o .tsx/.ts
# Los contenedores deberían recompilar automáticamente
docker compose -f docker-compose.dev.yml logs -f backend-dev
```

## 🚨 Solución de Problemas

### El stack no arranca
```bash
# 1. Verificar Docker
docker compose version  # Debería ser v2.39+

# 2. Limpiar todo
docker compose -f docker-compose.dev.yml down --remove-orphans
docker system prune -f

# 3. Rebuild completo
docker compose -f docker-compose.dev.yml up --build --force-recreate
```

### Problemas de conectividad con Supabase
```bash
# Verificar variables de entorno
docker compose -f docker-compose.dev.yml exec backend-dev env | grep DATABASE

# Test de conectividad
curl http://localhost:8080/api/v1/test/supabase
```

### El frontend no carga
1. Verificar que el puerto 3000 esté libre
2. Verificar logs: `docker compose -f docker-compose.dev.yml logs frontend-dev`
3. Rebuild: `docker compose -f docker-compose.dev.yml up --build frontend-dev`

### El backend no responde
1. Verificar que el puerto 8080 esté libre
2. Verificar logs: `docker compose -f docker-compose.dev.yml logs backend-dev`
3. Verificar health check: `curl http://localhost:8080/health`

### Watch mode no funciona
1. Verificar que estás en Docker Compose v2.39+
2. Los cambios deben estar en `src/` (frontend) o en archivos `.go` (backend)
3. Revisar logs para ver si detecta cambios

## 🏆 Estado Actual del Proyecto

- ✅ Docker Compose v2.39+ configurado
- ✅ 100% conectividad con Supabase (12/12 tablas)
- ✅ Hot reload funcionando en ambos servicios
- ✅ Watch mode activado
- ✅ Health checks configurados
- ✅ API endpoints documentados
- ✅ Entorno colaborativo listo
