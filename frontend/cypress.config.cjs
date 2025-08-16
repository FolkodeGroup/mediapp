// frontend/cypress.config.js

const { defineConfig } = require('cypress');

module.exports = defineConfig({
  e2e: {
    // La URL base de tu aplicación en desarrollo.
    // ¡Asegúrate de que el puerto sea el correcto! (ej. 5173 para Vite)
    baseUrl: 'http://localhost:3000',

    setupNodeEvents(on, config) {
      // implement node event listeners here
    },
  },
});