# 🏥 MediApp - Guía de Desarrollo Colaborativo

## ⚠️ Importante: Usamos Docker Compose moderno
Este proyecto usa **Docker Compose v2.39+** con sintaxis moderna. No necesitas instalar `docker-compose` por separado si tienes Docker Desktop o Docker Engine reciente.

### 🔧 Verificar tu versión de Docker Compose
```bash
# Verificar versión (debería ser v2.39+)
docker compose version

# Si tienes Docker Desktop, ya tienes la versión correcta
# Si usas Linux server, asegúrate de tener Docker Engine reciente
```

---

## 🚀 Cómo ejecutar en modo desarrollo

### Opción 1: Desarrollo con Docker (Recomendado) 🐳
1. **Levantar todo el stack de desarrollo**:
   ```bash
   docker compose -f docker-compose.dev.yml up --build
   ```

2. **Levantar en segundo plano**:
   ```bash
   docker compose -f docker-compose.dev.yml up --build -d
   ```

3. **Parar servicios**:
   ```bash
   docker compose -f docker-compose.dev.yml down
   ```

4. **Rebuild completo**:
   ```bash
   docker compose -f docker-compose.dev.yml up --build --force-recreate
   ```

5. **Verificar logs**:
   ```bash
   docker compose -f docker-compose.dev.yml logs -f
   ```

---

### Opción 2: Desarrollo Local
1. **Frontend**:
   ```bash
   cd frontend
   npm install
   npm run dev
   ```
   El frontend estará disponible en: [http://localhost:3000](http://localhost:3000)

2. **Backend** (en otra terminal):
   ```bash
   cd backend
   go mod download
   go run cmd/server/main.go
   ```
   El backend estará disponible en: [http://localhost:8080](http://localhost:8080)

3. **Configurar variables de entorno**:
   Asegúrate de que las variables en el archivo `.env` estén correctamente configuradas para conectar con Supabase.

---

## 🔧 Configuración de Servicios

### 1. Backend (Go)
1. **Instalar dependencias**:
   ```bash
   cd backend
   go mod download
   ```

2. **Configurar migraciones**:
   - Asegúrate de que el servidor PostgreSQL esté corriendo.
   - Aplica las migraciones y el seed inicial:
     ```bash
     goose -dir migrations postgres "postgres://usuario:contraseña@localhost:5432/mediapp_db?sslmode=disable" up
     ```

3. **Ejecutar el servidor**:
   ```bash
   go run ./cmd/server/main.go
   ```
   El backend estará disponible en: [http://localhost:8080](http://localhost:8080).

4. **Verificar el estado del backend**:
   - Health Check: [http://localhost:8080/health](http://localhost:8080/health)
   - Swagger Docs: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

---

### 2. Frontend (React + Vite)
1. **Instalar dependencias**:
   ```bash
   cd frontend
   npm install
   ```

2. **Ejecutar en modo desarrollo**:
   ```bash
   npm run dev
   ```
   El frontend estará disponible en: [http://localhost:3000](http://localhost:3000).

3. **Proxy API**:
   - El frontend está configurado para redirigir automáticamente las solicitudes a la API del backend.

---

### 3. Base de Datos (Supabase)
1. **Conexión a Supabase**:
   - Asegúrate de que las variables de entorno en [`.env`](.env ) estén configuradas correctamente:
     ```env
     DATABASE_URL=postgres://usuario:contraseña@host:puerto/dbname
     ```

2. **Verificar conectividad**:
   ```bash
   curl -s http://localhost:8080/api/v1/connect/all-tables | jq '.summary'
   ```
   Deberías obtener un resumen con `connection_rate: 100`.

---

## 🛠️ Herramientas de Desarrollo
- **Hot Reload**:
  - Backend: Configurado con `Air` para recargar automáticamente al detectar cambios en archivos `.go`.
  - Frontend: Configurado con Vite para recargar automáticamente al detectar cambios en `src/`.

- **Health Checks**:
  - Backend: [http://localhost:8080/health](http://localhost:8080/health)

- **Testing**:
  - Backend:
    ```bash
    cd backend
    go test ./...
    ```
  - Frontend:
    ```bash
    npm test
    ```

---

## 📝 Notas Importantes
1. **Puertos**:
   - Backend: 8080
   - Frontend: 3000
2. **Variables de entorno**:
   - Asegúrate de que [`.env`](.env ) esté correctamente configurado.
3. **JWT**:
   - La clave `JWT_SECRET_KEY` ser consistente entre todos los desarrolladores.
4. **Base de datos compartida**:
   - Todos los desarrolladores usan la misma instancia de Supabase para mantener consistencia.

---

## 🚨 Solución de Problemas
1. **El backend no responde**:
   - Verifica el puerto 8080.
   - Revisa los logs:
     ```bash
     docker compose -f docker-compose.dev.yml logs backend-dev
     ```
   - Verifica el health check:
     ```bash
     curl http://localhost:8080/health
     ```

2. **El frontend no carga**:
   - Verifica el puerto 3000.
   - Revisa los logs:
     ```bash
     docker compose -f docker-compose.dev.yml logs frontend-dev
     ```

3. **Problemas de conectividad con Supabase**:
   - Verifica las variables de entorno:
     ```bash
     docker compose -f docker-compose.dev.yml exec backend-dev env | grep DATABASE
     ```
   - Prueba la conectividad:
     ```bash
     curl http://localhost:8080/api/v1/test/supabase
     ```

---

## 🌐 URLs de Desarrollo
- **Frontend**: [http://localhost:3000](http://localhost:3000)
- **Backend API**: [http://localhost:8080](http://localhost:8080)
- **Health Check**: [http://localhost:8080/health](http://localhost:8080/health)
- **Swagger Docs**: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
