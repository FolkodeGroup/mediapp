

# 🧪 Testing de API

> Se testearon y documentaron los siguientes endpoints de la API de MediApp utilizando el cliente de Insomnia.

> Documentado y testeado por Facu

---

## 🏥 Endpoints de Pacientes

- **GET `/api/v1/pacientes`**
    - Retorna un listado de todos los pacientes.
    - **Respuesta esperada:** Array de objetos paciente, campo `total` y `status: "success"`.

    **Respuesta del Endpoint (TESTEADO):**
    - **Error 500: Internal Server Error**
    - **Body response:**
        ```json
        {
            "error": "Error interno del servidor"
        }
        ```

    **Console log:**
    ```log
    {"level":"error","ts":1755746694.286635,"caller":"handlers/pacientes.go:66","msg":"Error al consultar pacientes","error":"ERROR: column \"nro_documento\" does not exist (SQLSTATE 42703)","stacktrace":"github.com/FolkodeGroup/mediapp/internal/handlers.(*PacienteHandler).GetPacientes\n\tF:/Proyectos programacion/Folkode-Projects/mediapp/backend/internal/handlers/pacientes.go:66\ngithub.com/gin-gonic/gin.(*Context).Next\n\tC:/Users/Facuu/go/pkg/mod/github.com/gin-gonic/gin@v1.10.1/context.go:185\ngithub.com/gin-gonic/gin.CustomRecoveryWithWriter.func1\n\tC:/Users/Facuu/go/pkg/mod/github.com/gin-gonic/gin@v1.10.1/recovery.go:102\ngithub.com/gin-gonic/gin.(*Context).Next\n\tC:/Users/Facuu/go/pkg/mod/github.com/gin-gonic/gin@v1.10.1/context.go:185\nmain.main.LoggingMiddleware.func5\n\tF:/Proyectos programacion/Folkode-Projects/mediapp/backend/internal/middleware/logging.go:24\ngithub.com/gin-gonic/gin.(*Context).Next\n\tC:/Users/Facuu/go/pkg/mod/github.com/gin-gonic/gin@v1.10.1/context.go:185\nmain.main.RequestIDMiddleware.func4\n\tF:/Proyectos programacion/Folkode-Projects/mediapp/backend/internal/middleware/request_id.go:25\ngithub.com/gin-gonic/gin.(*Context).Next\n\tC:/Users/Facuu/go/pkg/mod/github.com/gin-gonic/gin@v1.10.1/context.go:185\ngithub.com/gin-gonic/gin.CustomRecoveryWithWriter.func1\n\tC:/Users/Facuu/go/pkg/mod/github.com/gin-gonic/gin@v1.10.1/recovery.go:102\ngithub.com/gin-gonic/gin.(*Context).Next\n\tC:/Users/Facuu/go/pkg/mod/github.com/gin-gonic/gin@v1.10.1/context.go:185\ngithub.com/gin-gonic/gin.LoggerWithConfig.func1\n\tC:/Users/Facuu/go/pkg/mod/github.com/gin-gonic/gin@v1.10.1/logger.go:249\ngithub.com/gin-gonic/gin.(*Context).Next\n\tC:/Users/Facuu/go/pkg/mod/github.com/gin-gonic/gin@v1.10.1/context.go:185\ngithub.com/gin-gonic/gin.(*Engine).handleHTTPRequest\n\tC:/Users/Facuu/go/pkg/mod/github.com/gin-gonic/gin@v1.10.1/gin.go:644\ngithub.com/gin-gonic/gin.(*Engine).ServeHTTP\n\tC:/Users/Facuu/go/pkg/mod/github.com/gin-gonic/gin@v1.10.1/gin.go:600\nnet/http.serverHandler.ServeHTTP\n\tC:/Program Files/Go/src/net/http/server.go:3340\nnet/http.(*conn).serve\n\tC:/Program Files/Go/src/net/http/server.go:2109"}
    ```

    - **Posible causa de error:** DB mal formulada.
    <br/>

- **GET `/api/v1/pacientes/{id}`**
    - Retorna el paciente que coincida con el `id` indicado.
    - **Respuesta esperada:** Objeto paciente.

    **Respuesta del Endpoint (TESTEADO, `id` = `$num`):**
    - **Error 404: Not Found**
    - **Body response:**
        ```json
        {
            "error": "Paciente no encontrado"
        }
        ```

---

## 🔍 Endpoints de Diagnóstico/Test

- **GET `/api/v1/test/supabase`**
    - Retorna estadísticas de la base de datos Supabase y el estado de conexión.

- **GET `/api/v1/connect/all-tables`**
    - Retorna estadísticas y sumario de la conexión con todas las tablas de la DB.

- **GET `/api/v1/inspect/tables?table={tableName}`**
    - Retorna información detallada de la tabla especificada en `tableName`.

- **GET `/health`**
    - Retorna información de estado y conexión de la base de datos.

- **GET `/swagger/index.html`**
    - Endpoint de la documentación interactiva Swagger.

    **Todos los endpoints anteriores (TESTEADOS):**
    - **Status 200: OK**
    - **Body responses correctos.**

---

## 🔐 Endpoints de Autenticación

- **POST `/login`**
    - Se debe enviar un body con los datos `username` y `password`.
    - **Ejemplo de body:**
        ```json
        {
            "username": "string",
            "password": "string"
        }
        ```

    **Console log:**
    ```log
    {"level":"info","ts":1755746621.2975824,"caller":"handlers/auth.go:45","msg":"Intento de login","requestID":"95414c5f-f729-4a15-9b00-8cd134824788","email":""}
    {"level":"error","ts":1755746621.7449331,"caller":"handlers/auth.go:69","msg":"Error al buscar usuario en la base de datos","error":"ERROR: syntax error at or near \"//\" (SQLSTATE 42601)","stacktrace":"github.com/FolkodeGroup/mediapp/internal/handlers.(*AuthHandler).Login\n\tF:/Proyectos programacion/Folkode-Projects/mediapp/backend/internal/handlers/auth.go:69\ngithub.com/gin-gonic/gin.(*Context).Next\n\tC:/Users/Facuu/go/pkg/mod/github.com/gin-gonic/gin@v1.10.1/context.go:185\ngithub.com/gin-gonic/gin.CustomRecoveryWithWriter.func1\n\tC:/Users/Facuu/go/pkg/mod/github.com/gin-gonic/gin@v1.10.1/recovery.go:102\ngithub.com/gin-gonic/gin.(*Context).Next\n\tC:/Users/Facuu/go/pkg/mod/github.com/gin-gonic/gin@v1.10.1/context.go:185\nmain.main.LoggingMiddleware.func5\n\tF:/Proyectos programacion/Folkode-Projects/mediapp/backend/internal/middleware/logging.go:24\ngithub.com/gin-gonic/gin.(*Context).Next\n\tC:/Users/Facuu/go/pkg/mod/github.com/gin-gonic/gin@v1.10.1/context.go:185\nmain.main.RequestIDMiddleware.func4\n\tF:/Proyectos programacion/Folkode-Projects/mediapp/backend/internal/middleware/request_id.go:25\ngithub.com/gin-gonic/gin.(*Context).Next\n\tC:/Users/Facuu/go/pkg/mod/github.com/gin-gonic/gin@v1.10.1/context.go:185\ngithub.com/gin-gonic/gin.CustomRecoveryWithWriter.func1\n\tC:/Users/Facuu/go/pkg/mod/github.com/gin-gonic/gin@v1.10.1/recovery.go:102\ngithub.com/gin-gonic/gin.(*Context).Next\n\tC:/Users/Facuu/go/pkg/mod/github.com/gin-gonic/gin@v1.10.1/context.go:185\ngithub.com/gin-gonic/gin.LoggerWithConfig.func1\n\tC:/Users/Facuu/go/pkg/mod/github.com/gin-gonic/gin@v1.10.1/logger.go:249\ngithub.com/gin-gonic/gin.(*Context).Next\n\tC:/Users/Facuu/go/pkg/mod/github.com/gin-gonic/gin@v1.10.1/context.go:185\ngithub.com/gin-gonic/gin.(*Engine).handleHTTPRequest\n\tC:/Users/Facuu/go/pkg/mod/github.com/gin-gonic/gin@v1.10.1/gin.go:644\ngithub.com/gin-gonic/gin.(*Engine).ServeHTTP\n\tC:/Users/Facuu/go/pkg/mod/github.com/gin-gonic/gin@v1.10.1/gin.go:600\nnet/http.serverHandler.ServeHTTP\n\tC:/Program Files/Go/src/net/http/server.go:3340\nnet/http.(*conn).serve\n\tC:/Program Files/Go/src/net/http/server.go:2109"}
    ```

- **GET `/protected`**
    - Endpoint protegido (requiere JWT en el header Authorization).

    **Respuesta del Endpoint (TESTEADO):**
    - **Error 401: Unauthorized**
    - **Body response:**
        ```json
        {
            "error": "Token requerido"
        }
        ```

---

## ✅ Notas de Testing

- Todos los endpoints retornan respuestas en formato JSON (excepto Swagger).
- Los endpoints protegidos requieren autenticación JWT.
- Se recomienda validar los campos y tipos de respuesta según la documentación oficial.
	