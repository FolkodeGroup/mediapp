# Logging Mejorado y Request ID

## Cambios Realizados

- Se implementó un middleware que genera e inyecta un `request_id` único en cada request HTTP.
- Todos los logs relevantes ahora incluyen el `request_id` para trazabilidad.
- Cada endpoint loguea su entrada con detalles clave (usuario, endpoint, parámetros relevantes).

## Ejemplo de Log

```
{"level":"info","ts":1692450000,"msg":"Intento de login","request_id":"b1c2d3e4-...","email":"usuario@ejemplo.com"}
```

## Pruebas de Logs

### Entorno Local
1. Levanta el backend:
   ```bash
   cd backend
   go run cmd/server/main.go
   ```
2. Realiza una petición (por ejemplo, login):
   ```bash
   curl -X POST http://localhost:8080/login -d '{"username":"usuario","password":"123"}' -H 'Content-Type: application/json'
   ```
3. Observa en consola que el log incluye un campo `request_id` único por request.

### Entorno Staging
1. Despliega la rama en staging.
2. Realiza peticiones a los endpoints.
3. Verifica en los logs del servidor que cada request tenga su propio `request_id` y que los logs de entrada de endpoint incluyan detalles relevantes.

---

Si tienes dudas sobre cómo interpretar los logs o necesitas ejemplos adicionales, consulta al equipo de backend.
