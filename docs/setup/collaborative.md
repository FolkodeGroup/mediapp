# 🏥 MediApp - Configuración de Desarrollo Colaborativo

## 🚀 Inicio Rápido

### Pre-requisitos
- Docker & Docker Compose v2.39+
- Git
- Acceso a Supabase (colaborativo)

### 📋 Configuración Inicial

1. **Clonar el repositorio:**
```bash
git clone <repository-url>
cd mediapp
```

2. **Configurar variables de entorno:**
```bash
# El archivo .env ya está configurado para Supabase
# Verificar que la DATABASE_URL es correcta
cp .env.example .env
```

3. **Levantar los servicios de desarrollo:**
```bash
# Detener servicios anteriores (si existen)
docker compose -f docker-compose.dev.yml down --remove-orphans

# Limpiar redes no utilizadas
docker network prune -f

# Construir y levantar servicios
docker compose -f docker-compose.dev.yml up --build
```

### 🔧 Docker Compose v2.39+ Features

Este proyecto usa las características más recientes de Docker Compose:

- **Watch mode**: Los cambios en código se sincronizan automáticamente
- **Environment variables**: Configuración centralizada en `.env`
- **Service dependencies**: Gestión automática de dependencias
- **Health checks**: Verificación automática del estado de servicios

### 🌐 URLs de Desarrollo

- **Frontend**: http://localhost:3000
- **Backend API**: http://localhost:8080
- **Health Check**: http://localhost:8080/health
- **Swagger Docs**: http://localhost:8080/swagger/index.html

### 🗄️ Base de Datos (Supabase)

El proyecto está configurado para usar **Supabase** como base de datos compartida:

- **Ventajas del enfoque colaborativo:**
  - Todos los desarrolladores usan la misma base de datos
  - Esquema consistente entre todos los entornos
  - Datos compartidos para pruebas
  - Sin necesidad de configurar PostgreSQL local

- **Variables de conexión:**
  - `DATABASE_URL`: Configurada en `.env` 
  - Conectividad automática con SSL requerido
  - Pooling de conexiones habilitado

### � **Estado de Conectividad: 100% ✅**

**12/12 tablas conectadas exitosamente:**

| Tabla | Columnas | Registros | Estado | Descripción |
|-------|----------|-----------|--------|-------------|
| `auditorias` | 5 | 0 | ✅ | Logs de auditoría del sistema |
| `consultorios` | 2 | 0 | ✅ | Información de consultorios |
| `datos_personales` | 5 | 0 | ✅ | Datos personales encriptados |
| `historia_clinica_version` | 9 | 0 | ✅ | Versiones de historias clínicas |
| `historias_clinicas` | 4 | 0 | ✅ | Historias clínicas principales |
| `pacientes` | 10 | 0 | ✅ | Información de pacientes |
| `permisos` | 2 | 7 | ✅ | Sistema de permisos |
| `recetas_medicas` | 6 | 0 | ✅ | Recetas médicas digitales |
| `rol_permiso` | 2 | 0 | ✅ | Relación roles-permisos |
| `roles` | 2 | 3 | ✅ | Roles del sistema |
| `turnos` | 5 | 0 | ✅ | Sistema de turnos |
| `usuarios` | 8 | 0 | ✅ | Usuarios del sistema |

### 🔌 **API Endpoints Disponibles**

#### **Gestión de Pacientes**
```bash
# Listar todos los pacientes
GET http://localhost:8080/api/v1/pacientes

# Obtener paciente específico
GET http://localhost:8080/api/v1/pacientes/{id}
```

#### **Diagnóstico y Verificación**
```bash
# Verificar conectividad completa con Supabase
GET http://localhost:8080/api/v1/test/supabase

# Conectar y analizar todas las tablas
GET http://localhost:8080/api/v1/connect/all-tables

# Inspeccionar estructura de tabla específica
GET http://localhost:8080/api/v1/inspect/tables?table=pacientes
GET http://localhost:8080/api/v1/inspect/tables?table=usuarios
GET http://localhost:8080/api/v1/inspect/tables?table=historia_clinica_version
```

#### **Autenticación**
```bash
# Login
POST http://localhost:8080/login

# Endpoint protegido (requiere JWT)
GET http://localhost:8080/protected
```

### �🔄 Comandos de Desarrollo

```bash
# Reiniciar servicios
docker compose -f docker-compose.dev.yml restart

# Ver logs en tiempo real
docker compose -f docker-compose.dev.yml logs -f

# Logs de un servicio específico
docker compose -f docker-compose.dev.yml logs -f backend-dev

# Detener servicios
docker compose -f docker-compose.dev.yml down

# Rebuild completo
docker compose -f docker-compose.dev.yml up --build --force-recreate
```

### 🔨 Watch Mode

Docker Compose v2.39+ incluye watch mode automático:

- **Frontend**: Cambios en `src/` se sincronizan automáticamente
- **Backend**: Cambios en código Go activan reconstrucción automática
- **Air**: Hot reload habilitado para Go

### � **Datos del Sistema**

#### **Roles Disponibles:**
- `medico` (ID: 1)
- `admin` (ID: 2) 
- `recepcionista` (ID: 3)

#### **Permisos Disponibles:**
- `ver_pacientes`
- `editar_pacientes`
- `ver_historia`
- `crear_turnos`
- `gestionar_usuarios`
- `auditar_sistema`
- `firmar_recetas`

### �🐛 Troubleshooting

#### Problemas de Conexión a Supabase
```bash
# Verificar variables de entorno
docker compose -f docker-compose.dev.yml exec backend-dev env | grep DATABASE

# Verificar conectividad
docker compose -f docker-compose.dev.yml exec backend-dev ping aws-1-us-east-2.pooler.supabase.com

# Verificar todas las tablas
curl -s http://localhost:8080/api/v1/connect/all-tables | jq '.summary'
```

#### Limpiar Completamente Docker
```bash
# Detener todos los contenedores
docker compose -f docker-compose.dev.yml down --remove-orphans

# Eliminar imágenes locales del proyecto
docker rmi mediapp-frontend-dev mediapp-backend-dev

# Limpiar sistema Docker
docker system prune -f
```

#### Regenerar Contenedores
```bash
# Forzar reconstrucción completa
docker compose -f docker-compose.dev.yml build --no-cache
docker compose -f docker-compose.dev.yml up --force-recreate
```

### 🧪 **Testing y Verificación**

#### Verificar Health Check
```bash
curl http://localhost:8080/health
# Respuesta esperada: {"db": true, "status": "ok"}
```

#### Verificar Conectividad Completa
```bash
curl -s http://localhost:8080/api/v1/connect/all-tables | jq '.summary'
# Debería mostrar: "connection_rate": 100, "successful_connections": 12
```

#### Inspeccionar Tabla Específica
```bash
# Ejemplo: inspeccionar tabla de usuarios
curl -s "http://localhost:8080/api/v1/inspect/tables?table=usuarios" | jq '.'
```

### 🤝 Trabajo en Equipo

- **Base de datos compartida**: Todos usan la misma instancia de Supabase
- **Variables de entorno**: Sincronizadas en el repositorio (sin credenciales sensibles)
- **Docker Compose**: Configuración idéntica para todos los desarrolladores
- **Hot reload**: Desarrollo ágil con cambios en tiempo real
- **API consistente**: Endpoints documentados y probados

### 📝 Notas Importantes

1. **No modificar** la `DATABASE_URL` sin coordinación del equipo
2. **El archivo `.env` está en el repositorio** pero `.env.local` puede usarse para overrides locales
3. **JWT_SECRET_KEY** debe ser consistente entre todos los desarrolladores
4. **Ports 3000 y 8080** deben estar libres en tu máquina local
5. **Nombres de tablas**: Usar nombres exactos de Supabase (ej: `historia_clinica_version` no `historia_clinica_versiones`)

### 🏗️ **Arquitectura del Proyecto**

```
mediapp/
├── frontend/              # React + TypeScript + Vite
│   ├── src/
│   ├── public/
│   └── Dockerfile.dev
├── backend/               # Go + Gin + pgx
│   ├── cmd/server/        # Main server
│   ├── internal/
│   │   ├── handlers/      # API handlers
│   │   ├── db/           # Database connection
│   │   ├── auth/         # JWT authentication
│   │   └── config/       # Configuration
│   └── Dockerfile.dev
├── docker-compose.dev.yml # Development environment
├── .env                  # Environment variables
└── SETUP-COLABORATIVO.md # This file
```

### 📈 **Métricas de Rendimiento**

- **Conexión a DB**: < 30s timeout
- **API Response**: < 1s promedio
- **Pool de conexiones**: Optimizado para concurrencia
- **Hot reload**: < 3s para cambios en código
- **Conectividad**: 100% de tablas accesibles

### 🆘 Contacto

Si tienes problemas con la configuración, contacta al equipo de desarrollo o revisa los logs:

```bash
# Ver logs del backend
docker compose -f docker-compose.dev.yml logs -f backend-dev

# Ver logs del frontend  
docker compose -f docker-compose.dev.yml logs -f frontend-dev
```
