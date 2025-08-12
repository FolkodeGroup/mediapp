# Instrucciones para aplicar migraciones y seed de base de datos

Hola @PaulaBigorra,

Para completar la configuración de la base de datos y aplicar las migraciones y el seed inicial, sigue estos pasos:

1. **Asegúrate de que el servidor PostgreSQL esté corriendo.**
2. **Crea la base de datos si aún no existe.** Por ejemplo:
   ```bash
   createdb -U postgres mediappdb
   ```
   (Reemplaza `mediappdb` por el nombre real que vayas a usar)
3. **Ejecuta las migraciones y el seed con goose:**
   ```bash
   goose -dir backend/migrations postgres "postgres://postgres:Administracion@localhost:5432/mediappdb?sslmode=disable" up
   ```
   (Ajusta usuario, contraseña y base de datos según corresponda)

Esto creará las tablas y agregará un médico de prueba con:
- **Email:** medico@prueba.com
- **Usuario:** medico_prueba
- **Contraseña:** Medico123 (hash bcrypt de ejemplo)

Si tienes dudas, revisa los archivos en `backend/migrations`.

¡Gracias!
