# Sistema de Autenticación y Autorización

## 1. Resumen

Este documento describe el sistema de autenticación basado en JSON Web Tokens (JWT) con firma asimétrica RSA-4096. El objetivo es asegurar los endpoints de la API, garantizando que solo los clientes autenticados puedan acceder a los recursos protegidos.

## 2. Generación de Claves

Las claves criptográficas son la base de la seguridad de nuestros JWT.

- **Algoritmo:** RSA
- **Longitud de Clave:** 4096 bits

### Proceso de Generación (con OpenSSL):
> Requisitos: Instalar `OpenSSL` agregando el binario al PATH del sistema en las variables del sistema.

1.  **Generar la Clave Privada:**
    ```bash
    openssl genrsa -out private.pem 4096
    ```
    *Este archivo `private.pem` es **confidencial** y solo debe residir en entornos seguros. **NO SUBAS ESTE ARCHIVO AL REPOSITORIO** 
    <br/>

2.  **Generar la Clave Pública:**
    ```bash
    openssl rsa -in private.pem -pubout -out public.pem
    ```
    *Este archivo `public.pem` puede ser distribuido y utilizado por los servicios que necesiten verificar los tokens. **NO SUBAS ESTE ARCHIVO AL REPOSITORIO**, pero puedes compartirlo con los usuarios para que utilicen los servicios necesarios*

## 3. Almacenamiento de Claves


-   **Entorno de Desarrollo:** Las claves se gestionan a través de un archivo `.env` en la raíz del proyecto. Este archivo está explícitamente excluido del control de versiones a través de `.gitignore`.
    -   `RSA_PRIVATE_KEY`: Contiene la clave privada PEM.
    -   `RSA_PUBLIC_KEY`: Contiene la clave pública PEM.

-   **Entorno de Producción (Recomendado):** Todas las claves se almacenan en HashiCorp Vault bajo la ruta `secret/jwt-keys`. La aplicación se autentica en Vault en el arranque para obtener las claves de forma segura.

## 4. Flujo del Token esperado

1.  **Autenticación:** El usuario envía sus credenciales (usuario y contraseña) al endpoint `/login`.
2.  **Generación de Token:** Si las credenciales son válidas, el backend genera un JWT, establece las `claims` (`sub`, `exp`, `iat`), y lo firma con la **clave privada RSA**.
3.  **Transmisión:** El JWT firmado se devuelve al cliente.
4.  **Autorización:** Para acceder a rutas protegidas, el cliente debe incluir el JWT en el encabezado `Authorization` con el prefijo `Bearer`.
5.  **Verificación:** El middleware de la API intercepta la solicitud, extrae el token y verifica su firma usando la **clave pública RSA**. Si la firma es válida y el token no ha expirado, se permite el acceso al recurso.