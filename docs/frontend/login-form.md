# Autenticación Local: Migración de Mock Service Worker a Backend Real

### Resumen del Cambio

- Este documento describe la migración del flujo de login en el frontend de MediApp, pasando de un entorno simulado con Mock Service Worker (MSW) a un entorno real utilizando el backend en Go corriendo en localhost:8080. Ahora, todas las peticiones de autenticación se realizan contra el servidor real, permitiendo pruebas y desarrollo más cercanos a producción.

### Puntos Principales del Cambio

- 1.  Eliminación de Mock Service Worker (MSW)
      Se eliminó cualquier referencia, importación o setup de MSW en el frontend.
      Las peticiones de login ahora se envían directamente al backend real usando Axios.
      Ya no se interceptan ni mockean respuestas en el navegador.
- 2.  Configuración de CORS en el Backend
      Se agregó y configuró el middleware CORS en el backend Go (Gin) para aceptar peticiones desde el frontend local (http://localhost:3000).

> El backend ahora responde correctamente a las peticiones OPTIONS (preflight) y POST, permitiendo la autenticación real desde el navegador SIN USAR MSW.

Fragmento relevante en main.go:

```go
router.Use(cors.New(cors.Config
{AllowOrigins:[]string{"http://localhost:3000"}, 
AllowMethods:     []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},    
AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},    
ExposeHeaders:    []string{"Content-Length"},    
AllowCredentials: true,    MaxAge:           12 * time.Hour,}))
```

Unificación de Campos de Login
   Se revisó y unificó el campo de usuario en el frontend y backend para asegurar que ambos usen el mismo nombre (username y password).
   El formulario de login y la función de autenticación esperan y envían los mismos campos que el backend espera.
4. Manejo de Errores según Código HTTP
   El frontend ahora muestra mensajes personalizados según el código HTTP recibido del backend:
   400: Solicitud inválida.
   401: Credenciales incorrectas o usuario bloqueado.
   500: Error interno del servidor.
   Network Error: Problemas de conexión con el backend.
   Esto mejora la experiencia de usuario y facilita la depuración.
   Instrucciones para Probar el Proyecto en Local
5. Backend (Go)
   Asegúrate de tener el backend corriendo en localhost:8080.
   Verifica que el middleware CORS esté correctamente configurado como se muestra arriba.
   Inicia el backend:

cd backendgo run ./cmd/server/main.go 2. Frontend (React)
Asegúrate de que el frontend corra en localhost:3000 (o ajusta el CORS en el backend si usas otro puerto).
Inicia el frontend:

cd frontendnpm installnpm run dev
Accede a la ruta de login y prueba el flujo de autenticación con usuarios válidos y no válidos. 3. Verificación de Flujo
Al enviar el formulario de login, la petición debe ir a http://localhost:8080/login.
Si el backend responde correctamente, deberías ser redirigido al dashboard.
Si hay errores, el mensaje mostrado dependerá del código HTTP recibido.
Si ves errores de CORS en la consola del navegador, revisa la configuración del backend y asegúrate de que ambos servidores estén corriendo en los puertos correctos.
Notas para el Equipo
Si necesitas cambiar el puerto del frontend, recuerda actualizar la configuración CORS en el backend.
Para ambientes de staging o producción, ajusta los orígenes permitidos en CORS según corresponda.
Si agregas nuevos endpoints protegidos, asegúrate de que también respondan correctamente a las peticiones OPTIONS.
Elimina cualquier referencia a MSW en nuevos archivos o tests.
