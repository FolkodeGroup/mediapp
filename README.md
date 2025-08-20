
# Mediapp

Plataforma fullstack para la gestión de pacientes en clínicas. Incluye frontend (React + Vite + TypeScript), backend (Go) y base de datos PostgreSQL.

---

## 📚 Tabla de Contenidos

1.  [Requisitos Previos](#1-requisitos-previos)
2.  [Inicio Rápido](#2-inicio-r%C3%A1pido)
    * [Configurar Variables de Entorno](#configurar-variables-de-entorno)
    * [Backend (Go)](#backend-go)
    * [Frontend (React + Vite)](#frontend-react--vite)
    * [Base de Datos (PostgreSQL)](#base-de-datos-postgresql)
3.  [Verificación y URLs Útiles](#3-verificaci%C3%B3n-y-urls-%C3%BAtiles)
4.  [Herramientas de Desarrollo](#4-herramientas-de-desarrollo)
    * [Hot Reload](#hot-reload)
    * [Health Checks](#health-checks)
5.  [Ejecución de Tests](#5-ejecuci%C3%B3n-de-tests)
    * [Tests del Frontend](#tests-del-frontend)
    * [Tests del Backend](#tests-del-backend)
6.  [Guía para Colaboradores](#6-gu%C3%ADa-para-colaboradores)
7.  [Despliegue](#7-despliegue)
8.  [Errores Comunes](#8-errores-comunes)
9.  [Contacto y Soporte](#9-contacto-y-soporte)

---

## 1. Requisitos Previos

Asegúrate de tener instalado lo siguiente en tu sistema:

* **Go**: Versión **1.21** o superior.
* **Node.js**: Versión **18** o superior (incluye `npm`).
* **PostgreSQL**: Un servidor de base de datos PostgreSQL funcionando localmente.
* **`goose`**: Herramienta para gestionar las migraciones de la base de datos de Go. Si no lo tienes, puedes instalarlo con:
    ```bash
    go install [github.com/pressly/goose/v3/cmd/goose@latest](https://github.com/pressly/goose/v3/cmd/goose@latest)
    ```

---

## 2. Inicio Rápido

Sigue estos pasos para levantar el proyecto en tu máquina local.

### Configurar Variables de Entorno

Crea un archivo llamado `.env` en el directorio **raíz del proyecto** (donde se encuentra este README.md). Configura las variables necesarias para la conexión con Supabase/PostgreSQL. Puedes usar el archivo `.env.example` como plantilla.

```ini
# Ejemplo de .env
DATABASE_URL="postgres://usuario:contraseña@localhost:5432/mediapp_db?sslmode=disable"
JWT_SECRET_KEY="tu_clave_secreta_jwt_aqui"
```

### Backend (Go)

1.  Navega al directorio del backend:
    ```bash
    cd backend
    ```
2.  Descarga las dependencias de Go:
    ```bash
    go mod download
    ```
3.  Aplica las migraciones de la base de datos. Asegúrate de que tu servidor PostgreSQL esté corriendo. **Reemplaza `usuario`, `contraseña`, `localhost:5432` y `mediapp_db` con tus propios valores de conexión:**
    ```bash
    goose -dir migrations postgres "postgres://usuario:contraseña@localhost:5432/mediapp_db?sslmode=disable" up
    ```
4.  Inicia el servidor backend:
    ```bash
    go run ./cmd/server/main.go
    ```
    El backend estará disponible en `http://localhost:8080`.

### Frontend (React + Vite)

1.  Abre una **nueva terminal** y navega al directorio del frontend:
    ```bash
    cd frontend
    ```
2.  Instala las dependencias de Node.js:
    ```bash
    npm install
    ```
3.  Inicia el servidor de desarrollo de React:
    ```bash
    npm run dev
    ```
    El frontend estará disponible en `http://localhost:3000`.

    * **Proxy API**: El frontend está configurado para redirigir automáticamente las solicitudes a la API del backend en `http://localhost:8080`.

### Base de Datos (PostgreSQL)

La aplicación backend se conectará a la base de datos PostgreSQL configurada en tu archivo `.env`. Puedes verificar la conectividad de la API con un `curl` (requiere `jq` para formatear la salida):

```bash
curl -s http://localhost:8080/api/v1/connect/all-tables | jq '.summary'
```
Deberías obtener un resumen con `connection_rate: 100`.

---

## 3. Verificación y URLs Útiles

Una vez que ambos servicios estén corriendo, puedes acceder a:

* **Frontend**: `http://localhost:3000`
* **Backend API**: `http://localhost:8080`
* **Health Check**: `http://localhost:8080/health` (Para verificar el estado del backend)
* **Swagger Docs**: `http://localhost:8080/swagger/index.html` (Documentación interactiva de la API)

---

## 4. Herramientas de Desarrollo

### Hot Reload

* **Backend**: Configurado para recargar automáticamente al detectar cambios en archivos `.go` (requiere una herramienta como [Air](https://github.com/cosmtrek/air) si deseas un hot reload automático para Go, de lo contrario, necesitarás reiniciar el comando `go run` manualmente).
* **Frontend**: Configurado con Vite para recargar automáticamente al detectar cambios en `src/`.

### Health Checks

* **Backend**: `http://localhost:8080/health`

---

## 5. Ejecución de Tests

### Tests del Frontend

Este proyecto utiliza **Jest** y **React Testing Library** para pruebas unitarias de componentes.

Para ejecutar los tests del frontend, usa:

```bash
cd frontend
npm test
```

Los tests están en archivos `*.test.tsx` dentro de la carpeta `src/components`.

### Tests del Backend

Para ejecutar todos los tests del backend:

```bash
cd backend
go test ./...
```

Este comando ejecutará todos los archivos de test (`*_test.go`) en los subdirectorios del backend y mostrará los resultados.

---

## 6. Guía para Colaboradores

Esta sección está destinada a los miembros del equipo que contribuyen directamente al proyecto. Sigue estas directrices para asegurar un flujo de trabajo colaborativo eficiente:

1.  **Clona el repositorio**: Descarga la copia principal del repositorio a tu máquina local:
    ```bash
    git clone <url-del-repo>
    cd mediapp
    ```
2.  **Configura el entorno de desarrollo**: Sigue las instrucciones detalladas en la sección [Inicio Rápido](#2-inicio-r%C3%A1pido) para configurar y ejecutar el proyecto localmente.
3.  **Crea una nueva rama**: Antes de empezar a trabajar en una nueva funcionalidad o corrección, crea una rama específica desde la rama `main` (o la rama de desarrollo principal, según el flujo de trabajo del equipo):
    ```bash
    git checkout main # Asegúrate de estar en la rama principal
    git pull origin main # Sincroniza con los últimos cambios
    git checkout -b <tipo>/<descripcion-corta-de-la-tarea>
    ```
    (Ejemplos de `<tipo>`: `feature/`, `bugfix/`, `docs/`, `refactor/`). Por ejemplo: `git checkout -b feature/login-form-validation`.
4.  **Realiza tus cambios**: Escribe tu código, realiza las pruebas necesarias y asegúrate de que todo funcione correctamente y cumpla con los estándares de calidad del proyecto.
5.  **Haz commits descriptivos**: Guarda tus cambios con mensajes de commit claros y concisos que expliquen lo que hiciste. Usa un formato como:
    ```bash
    git add .
    git commit -m "feat: añade la funcionalidad de validación del formulario de login"
    ```
    (Prefiere `feat:` para nuevas funcionalidades, `fix:` para correcciones, `docs:` para documentación, `refactor:` para refactorizaciones, etc.)
6.  **Sube tus cambios**: Empuja tu rama al repositorio principal:
    ```bash
    git push origin <tipo>/<descripcion-corta-de-la-tarea>
    ```
7.  **Abre un Pull Request (PR)**: Dirígete al repositorio en GitHub y abre un Pull Request desde tu rama hacia la rama `main` (o la rama de integración). Describe detalladamente los cambios realizados, el problema que resuelve o la funcionalidad que añade para facilitar la revisión.

---

## 7. Despliegue

Esta sección describe los pasos generales para desplegar la aplicación en un entorno de producción o staging.

1.  **Build del Frontend**:
    Navega al directorio `frontend` y genera los archivos estáticos de producción:
    ```bash
    cd frontend
    npm run build
    ```
    Esto creará una carpeta `dist` con los archivos optimizados para despliegue.

2.  **Build del Backend**:
    Navega al directorio `backend` y compila el binario ejecutable del servidor:
    ```bash
    cd backend
    go build -o server ./cmd/server/main.go
    ```
    Esto creará un archivo ejecutable llamado `server` en el directorio `backend`.

3.  **Configuración del Entorno de Producción**:
    Asegúrate de que las variables de entorno para el entorno de producción (ej. `DATABASE_URL`, `JWT_SECRET_KEY`) estén correctamente configuradas en el servidor de destino.

4.  **Ejecutar los Binarios**:

    * **Para el Backend**: Copia el binario `server` al servidor de producción y ejecútalo (puedes usar un gestor de procesos como systemd o Supervisor para mantenerlo en ejecución).
    * **Para el Frontend**: Sube el contenido de la carpeta `frontend/dist` a un servidor web (como Nginx, Apache o un servicio de hosting de archivos estáticos) que se encargará de servir los archivos a los usuarios.

*Nota: Los pasos de despliegue específicos pueden variar significativamente dependiendo del proveedor de hosting, la infraestructura utilizada y las prácticas de CI/CD.*

---

## 8. Errores Comunes

* **El puerto ya está en uso**:
    * Verifica si otro proceso está usando el puerto `3000` (frontend) o `8080` (backend). Puedes identificar y cerrar el proceso o cambiar los puertos en la configuración del proyecto (ej. `vite.config.ts` para frontend, o el código Go para el backend).
* **No se conecta a la base de datos**:
    * Verifica que tu servidor PostgreSQL esté corriendo.
    * Revisa que las variables de entorno en tu archivo `.env` estén correctamente configuradas (`DATABASE_URL`).
    * Asegúrate de que las credenciales (usuario, contraseña) y el nombre de la base de datos sean correctos.
* **Fallo al instalar dependencias (npm/go)**:
    * Para problemas con `npm install`, intenta borrar la carpeta `node_modules` y el archivo `package-lock.json` en `frontend`, y luego ejecuta `npm install` de nuevo.
    * Para problemas con `go mod download`, intenta borrar el archivo `go.sum` en `backend` y ejecuta `go mod download` nuevamente.
* **Error de `goose` al aplicar migraciones**:
    * Verifica que tienes `goose` instalado correctamente (`goose version`).
    * Asegúrate de que la cadena de conexión de PostgreSQL que pasas a `goose` sea válida y que el usuario tenga los permisos necesarios para crear tablas.

---

## 9. Contacto y Soporte

Para cualquier duda, sugerencia o para reportar un bug, por favor, abre un [issue en GitHub](enlace-a-issues-de-tu-repo) o contacta directamente al equipo de Folkode Group.