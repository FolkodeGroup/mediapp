import { http } from 'msw';

const mockUser = {
  id: 1,
  username: 'usuario',
  email: 'usuario@example.com',
  token: 'fake-jwt-token-123',
};

export const handlers = [
  http.post('/api/auth/login', (req, res, ctx) => {
    const { username, password } = req.body;

    if (username === 'usuario' && password === '123') {
      return res(
        ctx.status(200),
        ctx.json({
          user: mockUser,
          token: mockUser.token,
          message: 'Login exitoso',
        })
      );
    }

    return res(
      ctx.status(401),
      ctx.json({
        message: 'Credenciales incorrectas',
      })
    );
  }),

  // Mock GET /patients
  http.get('/api/patients', (req, res, ctx) => {
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
    return res(
      ctx.status(200),
      ctx.json({ patients })
    );
  }),
];