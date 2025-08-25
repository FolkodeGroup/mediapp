# Documentación Técnica: Flujo de Autenticación Mock

**Autor:** Facundo David Carrizo Lucero  
**Fecha:** 19-08-2025  
**Versión:** 1.0  

---

## 1. Resumen

Este documento detalla el flujo completo de autenticación de usuarios implementado en el frontend de la aplicación.  
Los cambios realizados representan un ciclo de autenticación completo (**login, gestión de sesión, cierre de sesión y protección de rutas**) utilizando un backend mock, lo que permite un desarrollo y pruebas desacoplados del servicio de autenticación real.

El objetivo principal es proporcionar un estado de autenticación global, persistente y robusto que controle el acceso a las diferentes secciones de la aplicación y mejore la experiencia del usuario con mensajes descriptivos de **error** y **success**.

---

## 2. Tecnologías Involucradas

- **Mock Service Worker (MSW):** Para interceptar peticiones de red y simular respuestas de la API a nivel de *service worker*.  
- **Context API:** Para la gestión del estado de autenticación global (en archivo `AuthContext.tsx`).  
- **Axios (Nueva!):** Cliente HTTP para realizar las peticiones de login.  
- **Zod (Nueva!):** Esquema de validación para el login.
- **React Router DOM:** Para la gestión de rutas, incluyendo la implementación de rutas protegidas y redirecciones.  
- **React Hook Form:** Para la gestión y validación de formularios.  

---

## 3. Diagrama del Flujo de Autenticación

```

Usuario en /login
       |
       v
[LoginForm.tsx] --(Ingresa credenciales)--> Llama a la función login() de AuthContext.tsx
       |
       v
[AuthContext.tsx] --(login())--> Realiza petición POST con Axios desde dicho archivo
       |
       v
[MSW Handler] --(/api/auth/login)--> Intercepta la petición
       |                                   |
 (Credenciales OK?)                        |
       | YES                               | NO
       v                                   v
Responde 200 OK + {token, user}     Responde 401 Unauthorized
       |                                   |
       v                                   v
[AuthContext.tsx] --(Actualiza estado)-->  [LoginForm.tsx] --(Maneja error)--> Muestra error con setMessage()
  - Guarda token y user en estado
  - Guarda en localStorage
  - isAuthenticated se vuelve 'true'
  - Muestra mensaje de success con setMessage()
       |
       v
[App] --(React re-renderiza)-->
       |
       v
[LoginForm.tsx] --(Redirige con useNavigate)--> /dashboard

```
---

## 4. Arquitectura y Componentes Clave

### 4.1. Mock Service Worker (`handlers.js`)

MSW intercepta las llamadas a la API salientes.  
Para la autenticación, se ha configurado un manejador que responde a `POST /api/auth/login`.

- **Lógica:** Comprueba si el username y password coinciden con credenciales predefinidas (`usuario/123`).  
- **Respuesta Exitosa (200):** Devuelve un objeto JSON con un token mock y la información del user.  
- **Respuesta de Error (401):** Devuelve un error 401 Unauthorized con un mensaje para el frontend.  

---

### 4.2. AuthContext (`AuthContext.tsx`)

**Estado que Gestiona:**
- `token (string | null)`: El JWT mock recibido de la API.  
- `user (User | null)`: Información del usuario autenticado.  
- `loading (boolean)`: Estado inicial durante la verificación en localStorage. Previene flashes innecesarios.  
- `isAuthenticated (boolean)`: Derivado de `!!token`. Flag principal para verificar autenticación.  

**Funciones Expuestas:**
- `login(credentials)`: Llama a la API, actualiza estado y guarda token/usuario en localStorage.  
- `logout()`: Limpia estado + localStorage y redirige al usuario a `/login`.  

**Persistencia de Sesión:**
- `useEffect` inicial que intenta restaurar sesión desde localStorage.  
- `loading` se pone en `false` en un `finally` para garantizar render correcto.  

---

### 4.3. Enrutamiento y Protección de Rutas (`App.tsx` y `PrivateRoute()`)

**PrivateRoute():**
- Consume `AuthContext` y revisa `isAuthenticated`.  
- Si `true`: renderiza dashboard.  
- Si `false`: redirige con `<Navigate />` a `/login`.  
- Requiere tanto `isAuthenticated` como `user` para acceder al dashboard.  

**Lógica adicional en `App.tsx`:**
- Si `isAuthenticated === true`, cualquier intento de acceder a `/login` redirige a `/dashboard`.  
- Previene que un usuario logueado vuelva a ver el login.  

---

### 4.4. Formulario de Login y Feedback (`LoginForm.tsx`)

**Interacción con Contexto:**
- El formulario no contiene lógica de API.  
- Usa `useAuth()` y llama `login()` en el `onSubmit`.  

**Manejo de Respuestas:**
- **Éxito (200 OK):**  
  - Muestra mensaje de éxito.  
  - Redirige a `/dashboard` con `useNavigate` tras 1 segundo.  
- **Error (401/500):**  
  - Usa `setMessage` para mostrar un error claro en el formulario.  

---

### 4.5. Logout Condicional (`MainLayout.tsx` o similar)

- Botón de **Cerrar Sesión** en el header, visible solo si `isAuthenticated === true`.  
- Al hacer clic, ejecuta `logout()`.  
- Esto borra sesión y redirige automáticamente a `/login`.  

---
## 5. Conclusión

Si hace falta algo más o hay dudas, **preguntar a Facu**, quien se encargó de esta implementación.  
¡Gracias!  

