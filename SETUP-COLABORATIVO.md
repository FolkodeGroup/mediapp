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

### 🔄 Comandos de Desarrollo

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

### 🐛 Troubleshooting

#### Problemas de Conexión a Supabase
```bash
# Verificar variables de entorno
docker compose -f docker-compose.dev.yml exec backend-dev env | grep DATABASE

# Verificar conectividad
docker compose -f docker-compose.dev.yml exec backend-dev ping aws-1-us-east-2.pooler.supabase.com
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

### 🤝 Trabajo en Equipo

- **Base de datos compartida**: Todos usan la misma instancia de Supabase
- **Variables de entorno**: Sincronizadas en el repositorio (sin credenciales sensibles)
- **Docker Compose**: Configuración idéntica para todos los desarrolladores
- **Hot reload**: Desarrollo ágil con cambios en tiempo real

### 📝 Notas Importantes

1. **No modificar** la `DATABASE_URL` sin coordinación del equipo
2. **El archivo `.env` está en el repositorio** pero `.env.local` puede usarse para overrides locales
3. **JWT_SECRET_KEY** debe ser consistente entre todos los desarrolladores
4. **Ports 3000 y 8080** deben estar libres en tu máquina local

### 🆘 Contacto

Si tienes problemas con la configuración, contacta al equipo de desarrollo.
