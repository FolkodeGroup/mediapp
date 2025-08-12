import { rest } from 'msw';

const mockUser = {
  id: 1,
  username: 'usuario',
  email: 'usuario@example.com',
  token: 'fake-jwt-token-123',
};

export const handlers = [
  rest.post('/api/auth/login', (req, res, ctx) => {
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

  // Puedes agregar más rutas aquí
  // rest.get('/api/profile', ...)
];