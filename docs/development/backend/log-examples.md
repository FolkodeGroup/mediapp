# Ejemplo de log generado por el sistema

```
{"level":"info","ts":1692450000,"msg":"Intento de login","request_id":"b1c2d3e4-...","email":"usuario@ejemplo.com"}
```

- `level`: Nivel del log (info, error, etc)
- `ts`: Timestamp
- `msg`: Mensaje del log
- `request_id`: ID único de la request
- Otros campos relevantes según el endpoint

## ¿Cómo probar?

1. Ejecuta el backend localmente.
2. Realiza peticiones a cualquier endpoint.
3. Observa en consola que cada log incluye un `request_id`.

## ¿Cómo ver logs en staging?

- Accede a los logs del servidor (según la infraestructura de staging).
- Busca por `request_id` para trazar una petición específica.
