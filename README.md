# React + TypeScript + Vite

This template provides a minimal setup to get React working in Vite with HMR and some ESLint rules.

Currently, two official plugins are available:

- [@vitejs/plugin-react](https://github.com/vitejs/vite-plugin-react/blob/main/packages/plugin-react) uses [Babel](https://babeljs.io/) for Fast Refresh
- [@vitejs/plugin-react-swc](https://github.com/vitejs/vite-plugin-react/blob/main/packages/plugin-react-swc) uses [SWC](https://swc.rs/) for Fast Refresh

## Expanding the ESLint configuration

If you are developing a production application, we recommend updating the configuration to enable type-aware lint rules:

```js
export default tseslint.config([
  globalIgnores(['dist']),
  {
    files: ['**/*.{ts,tsx}'],
    extends: [
      // Other configs...

      // Remove tseslint.configs.recommended and replace with this
      ...tseslint.configs.recommendedTypeChecked,
      // Alternatively, use this for stricter rules
      ...tseslint.configs.strictTypeChecked,
      // Optionally, add this for stylistic rules
      ...tseslint.configs.stylisticTypeChecked,

      // Other configs...
    ],
    languageOptions: {
      parserOptions: {
        project: ['./tsconfig.node.json', './tsconfig.app.json'],
        tsconfigRootDir: import.meta.dirname,
      },
      // other options...
    },
  },
])
```

You can also install [eslint-plugin-react-x](https://github.com/Rel1cx/eslint-react/tree/main/packages/plugins/eslint-plugin-react-x) and [eslint-plugin-react-dom](https://github.com/Rel1cx/eslint-react/tree/main/packages/plugins/eslint-plugin-react-dom) for React-specific lint rules:

```js
// eslint.config.js
import reactX from 'eslint-plugin-react-x'
import reactDom from 'eslint-plugin-react-dom'

export default tseslint.config([
  globalIgnores(['dist']),
  {
    files: ['**/*.{ts,tsx}'],
    extends: [
      // Other configs...
      // Enable lint rules for React
      reactX.configs['recommended-typescript'],
      // Enable lint rules for React DOM
      reactDom.configs.recommended,
    ],
    languageOptions: {
      parserOptions: {
        project: ['./tsconfig.node.json', './tsconfig.app.json'],
        tsconfigRootDir: import.meta.dirname,
      },
      // other options...
    },
  },
])
```
# 📝 Documentación de Pruebas de Frontend

Este documento detalla la configuración y el flujo de trabajo para las pruebas unitarias y de integración de nuestro frontend. Hemos configurado herramientas esenciales como **Jest** (el *test runner*) y **React Testing Library** (para interactuar con nuestros componentes como lo haría un usuario), asegurando la compatibilidad con TypeScript.

---

## 🚀 Configuración del Entorno de Pruebas

Para que el entorno de pruebas funcione correctamente, es fundamental instalar todas las dependencias necesarias y ajustar algunos archivos de configuración clave.

### 1. Instalación de Dependencias

Abre tu terminal en la raíz del proyecto y ejecuta el siguiente comando para instalar todas las herramientas de prueba necesarias. Estas se añadirán como "dependencias de desarrollo" (`devDependencies`).

```bash
npm install --save-dev jest @testing-library/react @testing-library/jest-dom @testing-library/user-event @types/jest ts-jest jest-environment-jsdom @babel/preset-env @babel/preset-react @babel/preset-typescript babel-jest

2. Configuración de Archivos Esenciales
Una vez instaladas las dependencias, crea o modifica los siguientes archivos para que Jest y TypeScript trabajen en armonía.

package.json
Modifica el script test y agrega la sección jest al final de tu archivo package.json. Asegúrate de que no haya comas extra o faltantes en la estructura JSON.

{
  "name": "mediapp",
  "private": true,
  "version": "0.0.0",
  "type": "module",
  "scripts": {
    "dev": "vite",
    "build": "vite build",
    "preview": "vite preview",
    "lint": "eslint --ext .ts,.tsx src --fix",
    "format": "prettier --write \"src/**/*.{ts,tsx,js,jsx,json,css,md}\"",
    "typecheck": "tsc --noEmit",
    "test": "jest"
  },
  "dependencies": {
    "react": "^19.1.1",
    "react-dom": "^19.1.1"
  },
  "devDependencies": {
    "@eslint/js": "^9.32.0",
    "@testing-library/jest-dom": "^6.6.4",
    "@testing-library/react": "^16.3.0",
    "@testing-library/user-event": "^14.6.1",
    "@types/jest": "^30.0.0",
    "@types/mocha": "^10.0.10",
    "@types/react": "^19.1.9",
    "@types/react-dom": "^19.1.7",
    "@typescript-eslint/eslint-plugin": "^8.39.0",
    "@typescript-eslint/parser": "^8.39.0",
    "@vitejs/plugin-react-swc": "^3.11.0",
    "autoprefixer": "^10.4.21",
    "eslint": "^9.33.0",
    "eslint-config-prettier": "^10.1.8",
    "eslint-plugin-prettier": "^5.5.4",
    "eslint-plugin-react-hooks": "^5.2.0",
    "eslint-plugin-react-refresh": "^0.4.20",
    "globals": "^16.3.0",
    "jest": "^29.7.0",
    "postcss": "^8.5.6",
    "prettier": "^3.6.2",
    "tailwindcss": "^3.4.3",
    "typescript": "~5.8.3",
    "typescript-eslint": "^8.39.0",
    "vite": "^7.1.1"
  },
  "jest": {
    "testEnvironment": "jsdom",
    "setupFilesAfterEnv": ["<rootDir>/src/setupTests.ts"],
    "transform": {
      "^.+\\.(ts|tsx)$": ["ts-jest", { "tsconfig": "tsconfig.test.json" }]
    },
    "moduleNameMapper": {
      "\\.(css|less|sass|scss)$": "identity-obj-proxy"
    }
  }
}

.babelrc
Crea este archivo en la raíz de tu proyecto (al mismo nivel que package.json). Esto le dice a Babel cómo transformar tu código.

// .babelrc
{
  "presets": [
    "@babel/preset-env",
    ["@babel/preset-react", { "runtime": "automatic" }],
    "@babel/preset-typescript"
  ]
}

tsconfig.test.json
Crea este archivo en la raíz de tu proyecto (al mismo nivel que package.json). Este archivo extenderá tu configuración principal de TypeScript y agregará las opciones específicas para el entorno de pruebas.

// tsconfig.test.json
{
  "extends": "./tsconfig.app.json",
  "compilerOptions": {
    "jsx": "react-jsx",
    "esModuleInterop": true,
    "module": "esnext",
    "moduleResolution": "node",
    "verbatimModuleSyntax": false,
    "types": ["jest", "@testing-library/jest-dom"]
  },
  "include": ["src/**/*.ts", "src/**/*.tsx"]
}

src/setupTests.ts
Crea este archivo dentro de tu carpeta src/. Esto asegurará que los "matchers" de jest-dom estén disponibles en todas tus pruebas.

// src/setupTests.ts
import '@testing-library/jest-dom';

Markdown

# 📝 Documentación de Pruebas de Frontend

Este documento detalla la configuración y el flujo de trabajo para las pruebas unitarias y de integración de nuestro frontend. Hemos configurado herramientas esenciales como **Jest** (el *test runner*) y **React Testing Library** (para interactuar con nuestros componentes como lo haría un usuario), asegurando la compatibilidad con TypeScript.

---

## 🚀 Configuración del Entorno de Pruebas

Para que el entorno de pruebas funcione correctamente, es fundamental instalar todas las dependencias necesarias y ajustar algunos archivos de configuración clave.

### 1. Instalación de Dependencias

Abre tu terminal en la raíz del proyecto y ejecuta el siguiente comando para instalar todas las herramientas de prueba necesarias. Estas se añadirán como "dependencias de desarrollo" (`devDependencies`).

```bash
npm install --save-dev jest @testing-library/react @testing-library/jest-dom @testing-library/user-event @types/jest ts-jest jest-environment-jsdom @babel/preset-env @babel/preset-react @babel/preset-typescript babel-jest
Aquí un breve resumen de lo que instala cada paquete:

jest: El motor principal para ejecutar tus pruebas.

@testing-library/react: Utilidades para renderizar y consultar tus componentes de React en un entorno de prueba.

@testing-library/jest-dom: Extensiones de Jest que proporcionan "matchers" adicionales para hacer aserciones sobre el DOM (por ejemplo, toBeInTheDocument(), toHaveClass()).

@testing-library/user-event: Simula interacciones de usuario de forma más realista (clics, escritura, etc.).

@types/jest: Tipos de TypeScript para Jest, esenciales para el autocompletado y la verificación de tipos.

ts-jest: Un transformador que permite a Jest procesar archivos TypeScript (.ts, .tsx).

jest-environment-jsdom: Proporciona un entorno de Document Object Model (DOM) simulado, necesario para que React Testing Library pueda "montar" tus componentes sin un navegador real.

@babel/preset-env, @babel/preset-react, @babel/preset-typescript: Presets de Babel que enseñan a Babel a entender sintaxis moderna de JavaScript, JSX de React y TypeScript, respectivamente.

babel-jest: El adaptador para que Jest pueda usar Babel para transformar el código.

2. Configuración de Archivos Esenciales
Una vez instaladas las dependencias, crea o modifica los siguientes archivos para que Jest y TypeScript trabajen en armonía.

package.json
Modifica el script test y agrega la sección jest al final de tu archivo package.json. Asegúrate de que no haya comas extra o faltantes en la estructura JSON.

JSON

{
  "name": "mediapp",
  "private": true,
  "version": "0.0.0",
  "type": "module",
  "scripts": {
    "dev": "vite",
    "build": "vite build",
    "preview": "vite preview",
    "lint": "eslint --ext .ts,.tsx src --fix",
    "format": "prettier --write \"src/**/*.{ts,tsx,js,jsx,json,css,md}\"",
    "typecheck": "tsc --noEmit",
    "test": "jest"
  },
  "dependencies": {
    "react": "^19.1.1",
    "react-dom": "^19.1.1"
  },
  "devDependencies": {
    "@eslint/js": "^9.32.0",
    "@testing-library/jest-dom": "^6.6.4",
    "@testing-library/react": "^16.3.0",
    "@testing-library/user-event": "^14.6.1",
    "@types/jest": "^30.0.0",
    "@types/mocha": "^10.0.10",
    "@types/react": "^19.1.9",
    "@types/react-dom": "^19.1.7",
    "@typescript-eslint/eslint-plugin": "^8.39.0",
    "@typescript-eslint/parser": "^8.39.0",
    "@vitejs/plugin-react-swc": "^3.11.0",
    "autoprefixer": "^10.4.21",
    "eslint": "^9.33.0",
    "eslint-config-prettier": "^10.1.8",
    "eslint-plugin-prettier": "^5.5.4",
    "eslint-plugin-react-hooks": "^5.2.0",
    "eslint-plugin-react-refresh": "^0.4.20",
    "globals": "^16.3.0",
    "jest": "^29.7.0",
    "postcss": "^8.5.6",
    "prettier": "^3.6.2",
    "tailwindcss": "^3.4.3",
    "typescript": "~5.8.3",
    "typescript-eslint": "^8.39.0",
    "vite": "^7.1.1"
  },
  "jest": {
    "testEnvironment": "jsdom",
    "setupFilesAfterEnv": ["<rootDir>/src/setupTests.ts"],
    "transform": {
      "^.+\\.(ts|tsx)$": ["ts-jest", { "tsconfig": "tsconfig.test.json" }]
    },
    "moduleNameMapper": {
      "\\.(css|less|sass|scss)$": "identity-obj-proxy"
    }
  }
}
.babelrc
Crea este archivo en la raíz de tu proyecto (al mismo nivel que package.json). Esto le dice a Babel cómo transformar tu código.

JSON

// .babelrc
{
  "presets": [
    "@babel/preset-env",
    ["@babel/preset-react", { "runtime": "automatic" }],
    "@babel/preset-typescript"
  ]
}
tsconfig.test.json
Crea este archivo en la raíz de tu proyecto (al mismo nivel que package.json). Este archivo extenderá tu configuración principal de TypeScript y agregará las opciones específicas para el entorno de pruebas.

JSON

// tsconfig.test.json
{
  "extends": "./tsconfig.app.json",
  "compilerOptions": {
    "jsx": "react-jsx",
    "esModuleInterop": true,
    "module": "esnext",
    "moduleResolution": "node",
    "verbatimModuleSyntax": false,
    "types": ["jest", "@testing-library/jest-dom"]
  },
  "include": ["src/**/*.ts", "src/**/*.tsx"]
}
src/setupTests.ts
Crea este archivo dentro de tu carpeta src/. Esto asegurará que los "matchers" de jest-dom estén disponibles en todas tus pruebas.

TypeScript

// src/setupTests.ts
import '@testing-library/jest-dom';
💡 Flujo de Trabajo para Escribir Tests
Una vez que el entorno está configurado, puedes empezar a escribir tus pruebas siguiendo estas pautas.

1. Ubicación de los Archivos de Prueba
Es una buena práctica colocar los archivos de prueba justo al lado del componente que están probando. Nómbralos usando la convención [NombreDelComponente].test.tsx.

2. Escribir Pruebas Básicas
Para cada componente, un test básico debería verificar los siguientes puntos:

Renderizado de Contenido: Asegúrate de que el componente se renderice correctamente en la pantalla y muestre el texto o los elementos hijos esperados.

Manejo de Eventos: Confirma que el componente responde adecuadamente a las interacciones del usuario (por ejemplo, clics en un botón, cambios en un input).

Aplicación de props: Verifica que las propiedades que pasas al componente (como variant, size, disabled, etc.) se apliquen correctamente y afecten la apariencia o el comportamiento del componente.

3. Ejecutar las Pruebas
Para ejecutar todas las pruebas de tu proyecto, abre tu terminal en la raíz del mismo y ejecuta el siguiente comando:


npm run test

Jest buscará automáticamente todos los archivos con las extensiones de prueba (.test.ts, .test.tsx) y te mostrará los resultados en la terminal.

🔗 Recursos Adicionales
Aquí tienes enlaces a la documentación oficial y a un tutorial que pueden serte de gran ayuda:

Jest (Documentación Oficial):
https://jestjs.io/docs/es-ES/getting-started

React Testing Library (Documentación Oficial):
https://testing-library.com/docs/react-testing-library/intro/

Tutorial en Video (Jest y React Testing Library en español):
APRENDE React Testing Library Tutorial (RTL) y Jest en español Paso a Paso