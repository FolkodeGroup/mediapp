
# Mediapp

Plataforma fullstack para gestión médica. Incluye frontend (React + Vite + TypeScript), backend (Go) y base de datos PostgreSQL, todo orquestado con Docker Compose.

---

## Tabla de contenidos

- [Requisitos previos](#requisitos-previos)
- [Cómo levantar el proyecto](#cómo-levantar-el-proyecto)
  - [Con Docker Compose (recomendado)](#con-docker-compose-recomendado)
  - [Manual (sin Docker)](#manual-sin-docker)
- [Flujos principales](#flujos-principales)
- [Ejecución de tests](#ejecución-de-tests)
- [Errores comunes](#errores-comunes)
- [Contacto y soporte](#contacto-y-soporte)

---

## Requisitos previos

- [Docker](https://www.docker.com/) y [Docker Compose](https://docs.docker.com/compose/) instalados
- Node.js >= 18 (para desarrollo local frontend)
- Go >= 1.21 (para desarrollo local backend)

---

## Cómo levantar el proyecto

### Con Docker Compose (recomendado)

1. Clona el repositorio:
   ```bash
   git clone <url-del-repo>
   cd mediapp
   ```
2. Copia los archivos de ejemplo de variables de entorno si existen (`.env.example` → `.env`).
3. Levanta todos los servicios:
   ```bash
   docker-compose up --build
   ```
   Esto inicia:
   - Frontend en [http://localhost:3000](http://localhost:3000) (o el puerto configurado)
   - Backend en [http://localhost:8080](http://localhost:8080)
   - Base de datos PostgreSQL

4. (Opcional) Aplica migraciones y seed manualmente si es necesario:
   ```bash
   docker-compose exec backend goose -dir backend/migrations postgres "postgres://mediapp_user:mediapp_password_2024@database:5432/mediapp_db?sslmode=disable" up
   ```

### Manual (sin Docker)

#### Backend
1. Instala dependencias:
   ```bash
   cd backend
   go mod download
   ```
2. Compila y ejecuta:
   ```bash
   go run ./cmd/server/main.go
   ```
3. Aplica migraciones y seed (requiere [goose](https://github.com/pressly/goose)):
   ```bash
   goose -dir migrations postgres "postgres://usuario:contraseña@localhost:5432/mediapp_db?sslmode=disable" up
   ```

#### Frontend
1. Instala dependencias:
   ```bash
   cd frontend
   npm install
   ```
2. Ejecuta en modo desarrollo:
   ```bash
   npm run dev
   ```

---

## Flujos principales

- **Inicio rápido:**
  - Clona el repo, instala dependencias y ejecuta con Docker Compose.
- **Desarrollo frontend:**
  - Usa `npm run dev` en `frontend/` para hot reload.
- **Desarrollo backend:**
  - Usa `go run ./cmd/server/main.go` en `backend/`.
- **Migraciones y seed:**
  - Usa goose para aplicar migraciones SQL y datos de prueba.

---

## Ejecución de tests

### Frontend

Este proyecto utiliza **Jest** y **React Testing Library** para pruebas unitarias de componentes.

Para ejecutar los tests de frontend, usa:

```sh
npx jest
```

O bien:

```sh
npm test
```

Los tests están en archivos `.test.tsx` dentro de `src/components`.

### Backend

Para ejecutar todos los tests del backend:

```sh
cd backend
go test ./...
```

---

## Errores comunes

- **El puerto ya está en uso:**
  - Cambia los puertos en `docker-compose.yml` o cierra el proceso que lo usa.
- **Problemas de permisos en Docker:**
  - Ejecuta con permisos adecuados o revisa la configuración de Docker Desktop.
- **No se conecta a la base de datos:**
  - Verifica variables de entorno y que el servicio `database` esté corriendo.
- **Fallo al instalar dependencias (npm/go):**
  - Borra `node_modules`/`go.sum` y reinstala.
- **Error de goose:**
  - Asegúrate de tener la versión correcta instalada y la cadena de conexión válida.

---

## Contacto y soporte

Para dudas, sugerencias o reportar bugs, abre un issue o contacta al equipo de Folkode Group.

## Ejecutar tests

Este proyecto utiliza **Jest** y **React Testing Library** para pruebas unitarias de componentes.

Para ejecutar los tests de frontend, usa el siguiente comando en la raíz del proyecto:

```sh
npx jest
```

También puedes usar el script de npm si está disponible:

```sh
npm test
```

> **Recomendación:** Instala la extensión "Jest" en VS Code para ejecutar y ver los resultados de los tests directamente en el editor.

Los tests se encuentran en archivos con la extensión `.test.tsx` dentro de la carpeta `src/components`.

---

## Ejecutar tests en el backend (Go)

Para ejecutar todos los tests del backend, usa el siguiente comando desde la raíz del proyecto o dentro de la carpeta `backend`:

```sh
go test ./...
```

Este comando ejecuta todos los archivos de test (`*_test.go`) en los subdirectorios del backend y muestra los resultados
