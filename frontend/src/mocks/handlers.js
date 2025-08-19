import { http } from 'msw';
import { HttpResponse } from 'msw';

const mockUser = {
  id: 1,
  username: 'usuario',
  email: 'usuario@example.com',
  token: 'fake-jwt-token-123',
};

export const handlers = [
  // Corrige el manejador para que sea asíncrono y use await req.json()
  http.post('/api/auth/login', async ({ request }) => {
    // 1. Usa "await" para parsear el cuerpo de la solicitud como JSON
    const { username, password } = await request.json();

    if (username === 'usuario' && password === '123') {
      // 2. Usa HttpResponse para construir la respuesta (práctica recomendada en MSW v2)
      return HttpResponse.json({
        user: mockUser,
        token: mockUser.token,
        message: 'Login exitoso',
      }, { status: 200 });
    }
    return HttpResponse.json(
      { message: 'Credenciales incorrectas' },
      { status: 401 }
    );
  }),

  // Mock GET /patients
  http.get('/api/patients', () => {
    // Estos datos deben estar sincronizados con el tipo Patient de TypeScript
    const patients = [
      {
        id: 1,
        firstName: 'Juan',
        lastName: 'Pérez',
        dni: '12345678',
        medicalRecordId: 'MR001',
        birthDate: '1980-01-01',
        gender: 'M',
        email: 'juan.perez@example.com',
      },
      {
        id: 2,
        firstName: 'Ana',
        lastName: 'García',
        dni: '87654321',
        medicalRecordId: 'MR002',
        birthDate: '1990-05-15',
        gender: 'F',
        email: 'ana.garcia@example.com',
      },
    ];
    return HttpResponse.json({ patients }, { status: 200 });
  }),
];