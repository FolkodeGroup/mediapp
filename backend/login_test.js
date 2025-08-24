// login_test.js
import http from 'k6/http';
import { check, sleep } from 'k6';

// --- Opciones de la Prueba ---
export const options = {
  stages: [
    { duration: '30s', target: 50 }, // Rampa de subida: de 0 a 50 usuarios en 30 segundos
    { duration: '1m', target: 50 },  // Fase de carga: mantener 50 usuarios durante 1 minuto
    { duration: '10s', target: 0 },  // Rampa de bajada: de 50 a 0 usuarios en 10 segundos
  ],
  thresholds: {
    // Criterios de aceptación:
    http_req_failed: ['rate<0.01'],   // La tasa de errores HTTP debe ser menor al 1%
    http_req_duration: ['p(95)<500'], // El 95% de las solicitudes deben completarse en menos de 500ms
  },
};

// --- Lógica de la Prueba (lo que hace cada usuario virtual) ---
export default function () {
  const baseURL = 'http://localhost:8080'; // Asegúrate de que sea la URL de tu API

  // 1. Iniciar Sesión
  const loginPayload = JSON.stringify({
    username: 'usuario', // Un usuario de prueba válido
    password: '123456',
  });

  const loginParams = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  const loginRes = http.post('http://localhost:8080/login', loginPayload, loginParams);

  // Verificar que el login fue exitoso
  check(loginRes, {
    'login: status is 200': (r) => r.status === 200,
    'login: response has refresh token': (r) => r.json('refresh_token') !== '',
  });

  // Extraer el refresh token para el siguiente paso
  const refreshToken = loginRes.json('refresh_token');

  // Simular tiempo de espera del usuario
  sleep(1);

  // 2. Refrescar el Token (solo si el login fue exitoso)
  if (refreshToken) {
    const refreshPayload = JSON.stringify({
      refresh_token: refreshToken,
    });

    const refreshRes = http.post(`${baseURL}/refresh`, refreshPayload, loginParams);

    // Verificar que el refresh fue exitoso
    check(refreshRes, {
      'refresh: status is 200': (r) => r.status === 200,
      'refresh: response has access token': (r) => r.json('access_token') !== '',
    });
  }
}