/// <reference types="vite/client" />


interface ImportMetaEnv {
  // Solo agrega variables personalizadas que empiecen con VITE_
  readonly VITE_API_BASE_URL: string;
  readonly VITE_APP_NAME: string;
  // No declares DEV, PROD, MODE aquí
}

// No necesitas redefinir ImportMeta si usas Vite

interface ImportMeta {
  readonly env: ImportMetaEnv;
}