import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import './index.css';
import App from './App.tsx';
import './assets/style.css';
import { Home } from './pages/Home';
import { Greeting } from './components/Greeting';

const root = document.querySelector('#app');
if (root) {
  root.appendChild(Home());
  root.appendChild(Greeting('Fede'));
}

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <App />
  </StrictMode>
);
