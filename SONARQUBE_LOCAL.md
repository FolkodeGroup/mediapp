---

## Checklist de integración SonarQube

- [x] Servicio SonarQube agregado a docker-compose.yml
- [x] Guía y plantilla sonar-project.properties.example en el repo
- [x] Binario SonarScanner instalado localmente en tools/ y **ignorado por git**
- [x] Archivo sonar-project.properties creado localmente (no subir)
- [x] Proyecto y token generados en la web de SonarQube
- [x] Análisis ejecutado y resultados visibles en el dashboard

---

## Notas importantes para el equipo

- El archivo `sonar-project.properties.example` es solo plantilla y **sí debe subirse** al repo.
- El archivo real `sonar-project.properties` (con tu token) **no debe subirse**.
- La carpeta tools/sonar-scanner/ y cualquier binario descargado deben estar en el `.gitignore`.

---

## ¿Cómo interpretar el dashboard de SonarQube?

En la pestaña **Issues** verás los problemas detectados en el código. Los principales conceptos son:

- **Maintainability**: Problemas de mantenibilidad (ej: duplicación, código difícil de entender, convenciones).
- **Reliability**: Problemas que pueden causar bugs o fallos en ejecución.
- **Security**: Problemas de seguridad o vulnerabilidades.
- **Severity**: Nivel de severidad (Blocker, High, Medium, Low, Info).
- **Clean Code Attribute**: Atributos de código limpio (Consistencia, Intencionalidad, Adaptabilidad).

Ejemplo de issues comunes:
- Ordenar alfabéticamente los paquetes en Dockerfile (Maintainability, Low)
- Unir instrucciones RUN en Dockerfile (Maintainability, Low)
- Renombrar funciones para cumplir convenciones (Maintainability, Low)

Cada issue tiene sugerencias y un estimado de esfuerzo para corregirlo. Puedes navegar, filtrar y asignar issues desde el dashboard.

---
# SonarQube local

1. Levanta los servicios (incluye SonarQube):

```bash
docker compose up -d sonarqube
```

2. Accede a SonarQube en: http://localhost:9000
   - Usuario/contraseña por defecto: admin / admin

3. Crea un nuevo proyecto desde la interfaz web.
   - Ve a "Projects" > "Create Project".
   - Asigna un nombre y una clave.
   - Elige "Manually" para el método de análisis.

4. Genera un token de autenticación para el proyecto.
   - Sigue el asistente tras crear el proyecto o ve a "My Account" > "Security" > "Generate Tokens".


5. Instala SonarScanner CLI (no lo subas al repo):
    - Descarga el binario oficial desde:
       https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-5.0.1.3006-linux.zip
    - Descomprime el archivo donde prefieras (por ejemplo, en tu home o en una carpeta tools/ que esté en el .gitignore).
    - **No subas la carpeta ni el zip al repositorio.**
    - Si lo dejas dentro del proyecto, asegúrate de que la ruta esté en el `.gitignore` (ya está configurado).

6. Crea un archivo `sonar-project.properties` en la raíz del proyecto con el siguiente contenido (ajusta los valores):

```
sonar.projectKey=TU_CLAVE_PROYECTO
sonar.projectName=TU_NOMBRE_PROYECTO
sonar.host.url=http://localhost:9000
sonar.login=TU_TOKEN
sonar.sources=./backend,./frontend/src
sonar.language=ts
sonar.sourceEncoding=UTF-8
```


7. Ejecuta el análisis desde la raíz del proyecto usando la ruta al binario:

```bash
./sonar-scanner-5.0.1.3006-linux/bin/sonar-scanner
# o si lo pusiste en tools/
./tools/sonar-scanner/bin/sonar-scanner
```


8. Revisa el dashboard de SonarQube en http://localhost:9000 para ver los resultados.

---

> **Nota:** Cada desarrollador debe descargar SonarScanner localmente y nunca subir el binario ni el zip al repo. Sigue estas instrucciones para mantener el repositorio limpio.

---

> Si tienes dudas sobre algún paso, consulta la documentación oficial: https://docs.sonarsource.com/sonarqube/latest/
