# 🏥 MediApp - Sistema de Gestión Médica

## 🚀 Cómo ejecutar en modo desarrollo

### Opción 1: Desarrollo Local (Recomendado)

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

3. **Base de datos**:
   ```bash
   docker run --name mediapp-postgres \
     -e POSTGRES_USER=mediapp_user \
     -e POSTGRES_PASSWORD=mediapp_password_2024 \
     -e POSTGRES_DB=mediapp_db \
     -p 5432:5432 \
     -d postgres:16-alpine
   ```

### Opción 2: Desarrollo con Docker

```bash
# Levantar todo el stack de desarrollo
npm run dev:docker

# O en segundo plano
npm run dev:docker-detached

# Para parar
npm run stop:dev
```

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

## 🐳 Docker para Producción

```bash
# Levantar en producción
npm run prod

# O en segundo plano
npm run prod:detached

# Para parar
npm run stop
```

## 📁 Estructura del Proyecto

```
mediapp/
├── frontend/           # React + TypeScript + Vite
│   ├── src/
│   ├── public/
│   ├── Dockerfile      # Producción
│   ├── Dockerfile.dev  # Desarrollo
│   └── vite.config.ts  # Configuración Vite
├── backend/            # Go + Gin
│   ├── cmd/
│   ├── internal/
│   ├── Dockerfile      # Producción
│   ├── Dockerfile.dev  # Desarrollo
│   └── .air.toml       # Hot reload Go
├── docker-compose.yml      # Producción
├── docker-compose.dev.yml  # Desarrollo
└── .env                    # Variables de entorno
```

## 🔧 Configuración de Desarrollo

### Frontend (Vite)
- **Puerto**: 3000
- **Hot Reload**: ✅
- **Proxy API**: `/api` → `http://localhost:8080`

### Backend (Go + Air)
- **Puerto**: 8080
- **Hot Reload**: ✅ (con Air)
- **Base de datos**: PostgreSQL en puerto 5432

### Base de Datos
- **PostgreSQL**: 16-alpine
- **Puerto**: 5432
- **Usuario**: mediapp_user
- **Contraseña**: mediapp_password_2024
- **Base de datos**: mediapp_db

## 🌐 URLs de Desarrollo

- **Frontend**: http://localhost:3000
- **Backend API**: http://localhost:8080
- **Base de datos**: localhost:5432

## 🛠️ Herramientas de Desarrollo

- **Vite**: Build tool y dev server para React
- **Air**: Hot reload para Go
- **MSW**: Mock Service Worker para testing
- **ESLint + Prettier**: Linting y formateo
- **Tailwind CSS**: Styling

## 📝 Comandos Útiles

```bash
# Ver logs de Docker en desarrollo
docker-compose -f docker-compose.dev.yml logs -f

# Acceder al contenedor del backend
docker exec -it mediapp-backend-dev sh

# Acceder al contenedor del frontend
docker exec -it mediapp-frontend-dev sh

# Limpiar contenedores y volúmenes
docker-compose -f docker-compose.dev.yml down -v
docker system prune -f
```

## 🚨 Solución de Problemas

### El frontend no arranca
1. Verificar que estás en el directorio `frontend/`
2. Ejecutar `npm install`
3. Verificar que el puerto 3000 esté libre

### El backend no arranca
1. Verificar que Go esté instalado
2. Ejecutar `go mod download` en `backend/`
3. Verificar que el puerto 8080 esté libre

### Problemas con Docker
1. Verificar que Docker esté corriendo
2. Limpiar contenedores: `docker-compose down -v`
3. Reconstruir: `docker-compose build --no-cache`
