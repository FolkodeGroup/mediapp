import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import './index.css';
import './assets/style.css';
import App from './App.tsx';

// Importa el worker solo si estamos en el navegador y en dev
// if (import.meta.env.DEV) {
//   const { worker } = await import('./mocks/browser');
//   worker.start({
//     onUnhandledRequest: 'warn', // Ayuda a detectar rutas no mockeadas
//   });
// }

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <App />
  </StrictMode>
);